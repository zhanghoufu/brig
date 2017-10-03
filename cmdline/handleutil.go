package cmdline

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/disorganizer/brig/brigd/client"
	"github.com/disorganizer/brig/brigd/server"
	"github.com/disorganizer/brig/daemon"
)

// ExitCode is an error that maps the error interface to a specific error
// message and a unix exit code
type ExitCode struct {
	Code    int
	Message string
}

func (err ExitCode) Error() string {
	return err.Message
}

// guessRepoFolder tries to find the repository path
// by using a number of sources.
func guessRepoFolder() string {
	return ""
	// folder := repo.GuessFolder()
	// if folder == "" {
	// 	log.Errorf("This does not look like a brig repository (missing .brig)")
	// 	os.Exit(BadArgs)
	// }

	// return folder
}

func readPassword() (string, error) {
	// TODO: Implement again.
	return "klaus", nil
	//	repoFolder := guessRepoFolder()
	//	pwd, err := pwdutil.PromptPasswordMaxTries(4, func(pwd string) bool {
	//		err := repo.CheckPassword(repoFolder, pwd)
	//		return err == nil
	//	})
	//
	return pwd, err
}

func prefixSlash(s string) string {
	if !strings.HasPrefix(s, "/") {
		return "/" + s
	}

	return s
}

type cmdHandlerWithClient func(ctx *cli.Context, client *daemon.Client) error

func withDaemon(handler cmdHandlerWithClient, startNew bool) func(*cli.Context) {
	// If not, make sure we start a new one:
	return withExit(func(ctx *cli.Context) error {
		port := guessPort()

		// Check if the daemon is running:
		ctl, err := client.Dial(context.Backround(), port)
		if err == nil {
			return handler(ctx, ctl)
		}

		if !startNew {
			// Daemon was not running and we may not start a new one.
			return ExitCode{DaemonNotResponding, "Daemon not running"}
		}

		// // Check if the password was supplied via a commandline flag.
		// pwd := ctx.String("password")
		// if pwd == "" {
		// 	// Prompt the user:
		// 	var cmdPwd string
		// 	cmdPwd, err = readPassword()
		// 	if err != nil {
		// 		return ExitCode{
		// 			BadPassword,
		// 			fmt.Sprintf("Could not read password: %v", pwd),
		// 		}
		// 	}

		// 	pwd = cmdPwd
		// }

		// Start the server & pass the password:
		daemon, err := server.BootServer(
			guessRepoFolder(),
			server.NewDummyBackend(),
		)

		if err != nil {
			return ExitCode{
				DaemonNotResponding,
				fmt.Sprintf("Unable to start daemon: %v", err),
			}
		}

		ctl, err := client.Dial(context.Backround(), port)
		if err != nil {
			return ExitCode{
				DaemonNotResponding,
				fmt.Sprintf("Unable to reach newly started daemon: %v", err),
			}
		}

		// Run the actual handler:
		return handler(ctx, ctl)
	})
}

type checkFunc func(ctx *cli.Context) int

func withArgCheck(checker checkFunc, handler func(*cli.Context)) func(*cli.Context) {
	return func(ctx *cli.Context) {
		if checker(ctx) != Success {
			os.Exit(BadArgs)
		}

		handler(ctx)
	}
}

func withExit(handler func(*cli.Context) error) func(*cli.Context) {
	return func(ctx *cli.Context) {
		if err := handler(ctx); err != nil {
			log.Error(err.Error())
			cerr, ok := err.(ExitCode)
			if !ok {
				os.Exit(UnknownError)
			}

			os.Exit(cerr.Code)
		}

		os.Exit(Success)
	}
}

func needAtLeast(min int) checkFunc {
	return func(ctx *cli.Context) int {
		if ctx.NArg() < min {
			if min == 1 {
				log.Warningf("Need at least %d argument.", min)
			} else {
				log.Warningf("Need at least %d arguments.", min)
			}
			cli.ShowCommandHelp(ctx, ctx.Command.Name)
			return BadArgs
		}

		return Success
	}
}

func guessPort() int {
	envPort := os.Getenv("BRIG_PORT")
	if envPort != "" {
		// Somebody tried to set BRIG_PORT.
		// Try to parse and spit errors if wrong.
		port, err := strconv.Atoi(envPort)
		if err != nil {
			log.Fatalf("Could not parse $BRIG_PORT: %v", err)
		}

		return port
	}

	// Guess the default port.
	log.Warning("BRIG_PORT not given, using :6666")
	return 6666
}
