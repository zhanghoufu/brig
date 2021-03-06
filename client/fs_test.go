package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
	colorlog "github.com/sahib/brig/util/log"
	"github.com/stretchr/testify/require"
)

var CurrBackendPort = 10000

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&colorlog.FancyLogFormatter{
		UseColors: true,
	})
}

func stringify(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}

func hardKillDaemonForPort(t *testing.T, port int) {
	pidPath := filepath.Join(os.TempDir(), fmt.Sprintf("brig.%d.pid", port))
	data, err := ioutil.ReadFile(pidPath)
	if os.IsNotExist(err) {
		// pid file does not exist yet.
		return
	}

	defer os.Remove(pidPath)

	// Handle other errors.
	require.Nil(t, err, stringify(err))

	pid := string(data)
	cmd := exec.Command("/bin/kill", "-9", pid)
	require.Nil(t, cmd.Start())
}

func withDaemon(t *testing.T, name string, port, backendPort int, fn func(ctl *Client)) {
	basePath, err := ioutil.TempDir("", "brig-ctl-test")
	require.Nil(t, err, stringify(err))

	// Path to the global registry file for this daemon:
	regPath := filepath.Join(os.TempDir(), "test-reg.yml")
	// Path to use for the net mock backend (stores dns names)
	netPath := filepath.Join(os.TempDir(), "test-net-dir")

	defer func() {
		os.RemoveAll(basePath)
		os.RemoveAll(regPath)
		os.RemoveAll(netPath)
	}()

	hardKillDaemonForPort(t, port)
	require.Nil(t, os.MkdirAll(basePath, 0700))

	cmd := exec.Command(
		"brig",
		"--repo", basePath,
		"--port", fmt.Sprintf("%d", port),
		"daemon", "launch",
		"--log-to-stdout",
	)

	cmd.Env = append(cmd.Env, fmt.Sprintf("BRIG_REGISTRY_PATH=/tmp/%s", regPath))
	cmd.Env = append(cmd.Env, fmt.Sprintf("BRIG_MOCK_NET_DB_PATH=%s", netPath))
	cmd.Env = append(cmd.Env, fmt.Sprintf("BRIG_MOCK_USER=%s", name))
	cmd.Env = append(cmd.Env, fmt.Sprintf("BRIG_MOCK_PORT=%d", backendPort))
	cmd.Env = append(cmd.Env, "BRIG_LOG_SHOW_PID=true")

	// Pipe the daemon output to the test output:
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.Nil(t, cmd.Start())

	pidPath := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("brig.%d.pid", port),
	)

	require.Nil(t,
		ioutil.WriteFile(
			pidPath,
			[]byte(fmt.Sprintf("%d", cmd.Process.Pid)),
			0644,
		),
	)

	// Timeout to make sure that the dameon started.
	// Only after this we try to connect normally.
	time.Sleep(200 * time.Millisecond)

	// Loop until we give up or have a valid connection.
	var ctl *Client
	for idx := 0; idx < 100; idx++ {
		ctl, err = Dial(context.Background(), port)
		if err == nil {
			break
		}

		time.Sleep(50 * time.Millisecond)
	}

	require.NotNil(t, ctl, "could not connect")

	// Make sure that the daemon shut down correctly.
	defer func() {
		// Send the death signal:
		require.Nil(t, ctl.Quit())

		// Wait until we cannot ping it anymore.
		daemonErroredOut := false
		for idx := 0; idx < 200; idx++ {
			if err := ctl.Ping(); err != nil {
				daemonErroredOut = true
				break
			}

			time.Sleep(50 * time.Millisecond)
		}

		if !daemonErroredOut {
			t.Fatalf("daemon is still up and running after quit")
		}
	}()

	// Init the repo. This should block until done.
	err = ctl.Init(basePath, name, "password", "mock")
	require.Nil(t, err, stringify(err))

	// Run the actual test function:
	fn(ctl)
}

func TestStageAndCat(t *testing.T) {
	withDaemon(t, "ali", 6667, 9999, func(ctl *Client) {
		fd, err := ioutil.TempFile("", "brig-dummy-data")
		path := fd.Name()

		require.Nil(t, err, stringify(err))
		_, err = fd.Write([]byte("hello"))
		require.Nil(t, err, stringify(err))
		require.Nil(t, fd.Close())

		require.Nil(t, ctl.Stage(path, "/hello"))
		rw, err := ctl.Cat("hello")
		require.Nil(t, err, stringify(err))

		data, err := ioutil.ReadAll(rw)
		require.Nil(t, err, stringify(err))

		require.Equal(t, []byte("hello"), data)
		require.Nil(t, rw.Close())
	})
}

func TestMkdir(t *testing.T) {
	withDaemon(t, "ali", 6667, 9999, func(ctl *Client) {
		// Create something nested with -p...
		require.Nil(t, ctl.Mkdir("/a/b/c", true))

		// Create it twice...
		require.Nil(t, ctl.Mkdir("/a/b/c", true))

		// Create something nested without -p
		err := ctl.Mkdir("/x/y/z", false)
		require.Contains(t, err.Error(), "No such file")

		require.Nil(t, ctl.Mkdir("/x", false))
		require.Nil(t, ctl.Mkdir("/x/y", false))
		require.Nil(t, ctl.Mkdir("/x/y/z", false))

		lst, err := ctl.List("/", -1)
		require.Nil(t, err, stringify(err))

		paths := []string{}
		for _, info := range lst {
			paths = append(paths, info.Path)
		}

		sort.Strings(paths)
		require.Equal(t, paths, []string{
			"/",
			"/a",
			"/a/b",
			"/a/b/c",
			"/x",
			"/x/y",
			"/x/y/z",
		})
	})
}

func withConnectedDaemonPair(t *testing.T, fn func(aliCtl, bobCtl *Client)) {
	// Use a shared directory for our shared data:
	basePath, err := ioutil.TempDir("", "brig-test-sync-pair-test")
	require.Nil(t, err, stringify(err))

	defer func() {
		CurrBackendPort += 2
		os.RemoveAll(basePath)
	}()

	withDaemon(t, "ali", 6668, CurrBackendPort, func(aliCtl *Client) {
		withDaemon(t, "bob", 6669, CurrBackendPort+1, func(bobCtl *Client) {
			aliWhoami, err := aliCtl.Whoami()
			require.Nil(t, err, stringify(err))

			bobWhoami, err := bobCtl.Whoami()
			require.Nil(t, err, stringify(err))

			// add bob to ali as remote
			err = aliCtl.RemoteAdd(Remote{
				Name:        "bob",
				Fingerprint: bobWhoami.Fingerprint,
			})
			require.Nil(t, err, stringify(err))

			// add ali to bob as remote
			err = bobCtl.RemoteAdd(Remote{
				Name:        "ali",
				Fingerprint: aliWhoami.Fingerprint,
			})
			require.Nil(t, err, stringify(err))

			fn(aliCtl, bobCtl)
		})
	})
}

func TestSyncBasic(t *testing.T) {
	withConnectedDaemonPair(t, func(aliCtl, bobCtl *Client) {
		err := aliCtl.StageFromReader("/ali_file", bytes.NewReader([]byte{42}))
		require.Nil(t, err, stringify(err))

		err = bobCtl.StageFromReader("/bob_file", bytes.NewReader([]byte{23}))
		require.Nil(t, err, stringify(err))

		_, err = aliCtl.Sync("bob", true)
		require.Nil(t, err, stringify(err))

		_, err = bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		// We cannot query the file contents, since the mock backend
		// does not yet store the file content anywhere.
		bobFileStat, err := aliCtl.Stat("/bob_file")
		require.Nil(t, err, stringify(err))
		require.Equal(t, "/bob_file", bobFileStat.Path)

		aliFileStat, err := bobCtl.Stat("/ali_file")
		require.Nil(t, err, stringify(err))
		require.Equal(t, "/ali_file", aliFileStat.Path)
	})
}

func pathsFromListing(l []StatInfo) []string {
	result := []string{}
	for _, entry := range l {
		result = append(result, entry.Path)
	}

	return result
}

func TestSyncConflict(t *testing.T) {
	withConnectedDaemonPair(t, func(aliCtl, bobCtl *Client) {
		// Create two files with the same content on both sides:
		err := aliCtl.StageFromReader("/README", bytes.NewReader([]byte{42}))
		require.Nil(t, err, stringify(err))

		err = bobCtl.StageFromReader("/README", bytes.NewReader([]byte{42}))
		require.Nil(t, err, stringify(err))

		// Sync and check if the files are still equal:
		_, err = bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		aliFileStat, err := aliCtl.Stat("/README")
		require.Nil(t, err, stringify(err))
		bobFileStat, err := bobCtl.Stat("/README")
		require.Nil(t, err, stringify(err))
		require.Equal(t, aliFileStat.ContentHash, bobFileStat.ContentHash)

		// Modify bob's side only. A sync should have no effect.
		err = bobCtl.StageFromReader("/README", bytes.NewReader([]byte{43}))
		require.Nil(t, err, stringify(err))

		_, err = bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		bobFileStat, err = bobCtl.Stat("/README")
		require.Nil(t, err, stringify(err))

		require.NotEqual(t, aliFileStat.ContentHash, bobFileStat.ContentHash)

		// Modify ali's side additionally. Now we should get a conflicting file.
		err = aliCtl.StageFromReader("/README", bytes.NewReader([]byte{41}))
		require.Nil(t, err, stringify(err))

		dirs, err := bobCtl.List("/", -1)
		require.Nil(t, err, stringify(err))
		require.Equal(t, []string{"/", "/README"}, pathsFromListing(dirs))

		_, err = bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		dirs, err = bobCtl.List("/", -1)
		require.Nil(t, err, stringify(err))
		require.Equal(
			t,
			[]string{"/", "/README", "/README.conflict.0"},
			pathsFromListing(dirs),
		)
	})
}

func TestSyncSeveralTimes(t *testing.T) {
	withConnectedDaemonPair(t, func(aliCtl, bobCtl *Client) {
		err := aliCtl.StageFromReader("/ali_file_1", bytes.NewReader([]byte{1}))
		require.Nil(t, err, stringify(err))

		_, err = bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		dirs, err := bobCtl.List("/", -1)
		require.Nil(t, err, stringify(err))
		require.Equal(
			t,
			[]string{"/", "/ali_file_1"},
			pathsFromListing(dirs),
		)

		err = aliCtl.StageFromReader("/ali_file_2", bytes.NewReader([]byte{2}))
		require.Nil(t, err, stringify(err))

		_, err = bobCtl.Sync("ali", true)

		require.Nil(t, err, stringify(err))

		dirs, err = bobCtl.List("/", -1)
		require.Nil(t, err, stringify(err))
		require.Equal(
			t,
			[]string{"/", "/ali_file_1", "/ali_file_2"},
			pathsFromListing(dirs),
		)

		err = aliCtl.StageFromReader("/ali_file_3", bytes.NewReader([]byte{3}))
		require.Nil(t, err, stringify(err))

		_, err = bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		dirs, err = bobCtl.List("/", -1)
		require.Nil(t, err, stringify(err))
		require.Equal(
			t,
			[]string{"/", "/ali_file_1", "/ali_file_2", "/ali_file_3"},
			pathsFromListing(dirs),
		)
	})
}

func TestSyncPartial(t *testing.T) {
	withConnectedDaemonPair(t, func(aliCtl, bobCtl *Client) {
		aliWhoami, err := aliCtl.Whoami()
		require.Nil(t, err, stringify(err))

		bobWhoami, err := bobCtl.Whoami()
		require.Nil(t, err, stringify(err))

		require.Nil(t, aliCtl.RemoteSave([]Remote{
			{
				Name:        "bob",
				Fingerprint: bobWhoami.Fingerprint,
				Folders:     []string{"/photos"},
			},
		}))

		require.Nil(t, bobCtl.RemoteSave([]Remote{
			{
				Name:        "ali",
				Fingerprint: aliWhoami.Fingerprint,
				Folders:     []string{"/photos"},
			},
		}))

		err = aliCtl.StageFromReader("/docs/ali_secret.txt", bytes.NewReader([]byte{0}))
		require.Nil(t, err, stringify(err))
		err = aliCtl.StageFromReader("/photos/ali.png", bytes.NewReader([]byte{42}))
		require.Nil(t, err, stringify(err))

		err = bobCtl.StageFromReader("/docs/bob_secret.txt", bytes.NewReader([]byte{0}))
		require.Nil(t, err, stringify(err))
		err = bobCtl.StageFromReader("/photos/bob.png", bytes.NewReader([]byte{23}))
		require.Nil(t, err, stringify(err))

		_, err = aliCtl.Sync("bob", true)
		require.Nil(t, err, stringify(err))

		_, err = bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		// We cannot query the file contents, since the mock backend
		// does not yet store the file content anywhere.
		aliLs, err := aliCtl.List("/", -1)
		require.Nil(t, err, stringify(err))

		aliPaths := []string{}
		for _, entry := range aliLs {
			aliPaths = append(aliPaths, entry.Path)
		}

		bobLs, err := bobCtl.List("/", -1)
		require.Nil(t, err, stringify(err))

		bobPaths := []string{}
		for _, entry := range bobLs {
			bobPaths = append(bobPaths, entry.Path)
		}

		require.Equal(
			t,
			[]string{
				"/",
				"/docs",
				"/photos",
				"/docs/ali_secret.txt",
				"/photos/ali.png",
				"/photos/bob.png",
			},
			aliPaths,
		)

		require.Equal(
			t,
			[]string{
				"/",
				"/docs",
				"/photos",
				"/docs/bob_secret.txt",
				"/photos/ali.png",
				"/photos/bob.png",
			},
			bobPaths,
		)
	})
}

func TestSyncMovedFile(t *testing.T) {
	withConnectedDaemonPair(t, func(aliCtl, bobCtl *Client) {
		require.Nil(t, aliCtl.StageFromReader("/ali-file", bytes.NewReader([]byte{1, 2, 3})))
		require.Nil(t, bobCtl.StageFromReader("/bob-file", bytes.NewReader([]byte{4, 5, 6})))

		aliDiff, err := aliCtl.Sync("bob", true)
		require.Nil(t, err, stringify(err))

		bobDiff, err := bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		require.Equal(t, aliDiff.Added[0].Path, "/bob-file")
		require.Equal(t, bobDiff.Added[0].Path, "/ali-file")

		require.Nil(t, aliCtl.Move("/ali-file", "/bali-file"))
		bobDiffAfter, err := bobCtl.Sync("ali", true)
		require.Nil(t, err, stringify(err))

		require.Len(t, bobDiffAfter.Added, 0)
		require.Len(t, bobDiffAfter.Removed, 0)
		require.Len(t, bobDiffAfter.Moved, 1)
	})
}
