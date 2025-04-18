package storage_test

import (
	"os"
	"testing"

	"github.com/coronon/go-storage"
	"github.com/coronon/go-storage/internal/testutils"
)

func withLocal(cb func(storage.FS)) {
	dir, err := os.MkdirTemp("", "go-storage-local-test")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	fs := storage.NewLocalFS(dir)
	cb(fs)
}

func TestLocalOpen(t *testing.T) {
	withLocal(func(fs storage.FS) {
		testutils.OpenNotExists(t, fs, "foo")
	})
}

func TestLocalCreate(t *testing.T) {
	withLocal(func(fs storage.FS) {
		testutils.Create(t, fs, "foo", "")
		testutils.Create(t, fs, "foo", "bar")
	})
}

func TestLocalDelete(t *testing.T) {
	withLocal(func(fs storage.FS) {
		testutils.Delete(t, fs, "foo")
	})
}
