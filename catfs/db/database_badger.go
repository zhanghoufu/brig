package db

import (

	// Because ipfs' package manager sucks a lot (sorry, but it does)
	// it imports badger with the import url below. This calls a few init()s,
	// which will panic when being called twice due to expvar defines e.g.
	// (i.e. when using the "correct" import github.com/dgraph-io/badger)
	//
	// So gx forces us to use their badger version for no good reason at all.

	"gx/ipfs/QmZ7bFqkoHU2ARF68y9fSQVKcmhjYrTQgtCQ4i3chwZCgQ/badger"
	"io"
	"strings"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// BadgerDatabase is a database implementation based on BadgerDB
type BadgerDatabase struct {
	mu         sync.Mutex
	db         *badger.DB
	txn        *badger.Txn
	refCount   int
	haveWrites bool
}

// NewBadgerDatabase creates a new badger database.
func NewBadgerDatabase(path string) (*BadgerDatabase, error) {
	opts := badger.DefaultOptions

	opts.Dir = path
	opts.ValueDir = path

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &BadgerDatabase{
		db: db,
	}, nil
}

func (db *BadgerDatabase) view(fn func(txn *badger.Txn) error) error {
	// If we have an open transaction, retrieve the values from there.
	// Otherwise we would not be able to retrieve in-memory values.
	if db.txn != nil {
		return fn(db.txn)
	}

	// If no transaction is running (no Batch()-call), use a fresh view txn.
	return db.db.View(fn)
}

// Get is the badger implementation of Database.Get.
func (db *BadgerDatabase) Get(key ...string) ([]byte, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	data := []byte{}
	err := db.view(func(txn *badger.Txn) error {
		if db.txn != nil {
			txn = db.txn
		}

		keyPath := strings.Join(key, ".")
		item, err := txn.Get([]byte(keyPath))
		if err == badger.ErrKeyNotFound {
			return ErrNoSuchKey
		}

		if err != nil {
			return err
		}

		data, err = item.ValueCopy(nil)
		return err
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Keys is the badger implementation of Database.Keys.
func (db *BadgerDatabase) Keys(fn func(key []string) error, prefix ...string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.view(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.IteratorOptions{})
		defer iter.Close()

		for iter.Rewind(); iter.Valid(); iter.Next() {
			item := iter.Item()

			fullKey := string(item.Key())
			splitKey := strings.Split(fullKey, ".")

			hasPrefix := len(prefix) <= len(splitKey)
			for i := 0; hasPrefix && i < len(prefix) && i < len(splitKey); i++ {
				if prefix[i] != splitKey[i] {
					hasPrefix = false
				}
			}

			if hasPrefix {
				db.mu.Unlock()
				if err := fn(strings.Split(fullKey, ".")); err != nil {
					db.mu.Lock()
					return err
				}
				db.mu.Lock()
			}
		}

		return nil
	})
}

// Export is the badger implementation of Database.Export.
func (db *BadgerDatabase) Export(w io.Writer) error {
	_, err := db.db.Backup(w, 0)
	return err
}

// Import is the badger implementation of Database.Import.
func (db *BadgerDatabase) Import(r io.Reader) error {
	return db.db.Load(r)
}

// Glob is the badger implementation of Database.Glob
func (db *BadgerDatabase) Glob(prefix []string) ([][]string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	fullPrefix := strings.Join(prefix, ".")

	results := [][]string{}
	err := db.view(func(txn *badger.Txn) error {
		iter := txn.NewIterator(badger.IteratorOptions{})
		defer iter.Close()

		for iter.Seek([]byte(fullPrefix)); iter.Valid(); iter.Next() {
			fullKey := string(iter.Item().Key())
			if !strings.HasPrefix(fullKey, fullPrefix) {
				break
			}

			// Don't do recursive globbing:
			leftOver := fullKey[len(fullPrefix):]
			if !strings.Contains(leftOver, ".") {
				results = append(results, strings.Split(fullKey, "."))
			}
		}

		return nil
	})

	return results, err
}

// Batch is the badger implementation of Database.Batch
func (db *BadgerDatabase) Batch() Batch {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.txn == nil {
		db.txn = db.db.NewTransaction(true)
	}

	db.refCount++
	return db
}

// Put is the badger implementation of Database.Put
func (db *BadgerDatabase) Put(val []byte, key ...string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.haveWrites = true

	fullKey := []byte(strings.Join(key, "."))
	db.txn.Set(fullKey, val)
}

// Clear is the badger implementation of Database.Clear
func (db *BadgerDatabase) Clear(key ...string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.haveWrites = true

	iter := db.txn.NewIterator(badger.IteratorOptions{})
	defer iter.Close()

	prefix := strings.Join(key, ".")

	keys := [][]byte{}
	for iter.Rewind(); iter.Valid(); iter.Next() {
		item := iter.Item()

		key := []byte{}
		keys = append(keys, item.KeyCopy(key))
	}

	for _, key := range keys {
		if !strings.HasPrefix(string(key), prefix) {
			continue
		}

		if err := db.txn.Delete(key); err != nil {
			return err
		}
	}

	return nil
}

// Erase is the badger implementation of Database.Erase
func (db *BadgerDatabase) Erase(key ...string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.haveWrites = true

	fullKey := []byte(strings.Join(key, "."))
	db.txn.Delete(fullKey)
}

// Flush is the badger implementation of Database.Flush
func (db *BadgerDatabase) Flush() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.refCount--
	if db.refCount > 0 {
		return nil
	}

	if db.refCount < 0 {
		log.Errorf("negative batch ref count: %d", db.refCount)
		return nil
	}

	defer db.txn.Discard()
	if err := db.txn.Commit(nil); err != nil {
		return err
	}

	db.txn = nil
	db.haveWrites = false
	return nil
}

// Rollback is the badger implementation of Database.Rollback
func (db *BadgerDatabase) Rollback() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.refCount--
	if db.refCount > 0 {
		return
	}

	if db.refCount < 0 {
		log.Errorf("negative batch ref count: %d", db.refCount)
		return
	}

	db.txn.Discard()
	db.txn = nil
	db.haveWrites = false
	db.refCount = 0
}

// HaveWrites is the badger implementation of Database.HaveWrites
func (db *BadgerDatabase) HaveWrites() bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.haveWrites
}

// Close is the badger implementation of Database.Close
func (db *BadgerDatabase) Close() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// With an open transaction it would deadlock:
	if db.txn != nil {
		db.txn.Discard()
		db.txn = nil
		db.haveWrites = false
	}

	if db.db != nil {
		oldDb := db.db
		db.db = nil
		if err := oldDb.Close(); err != nil {
			return err
		}
	}

	return nil
}
