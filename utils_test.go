package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coronon/go-storage"
	"github.com/coronon/go-storage/internal/testutils"
)

func TestExists(t *testing.T) {
	ctx := t.Context()

	withMem(func(fs storage.FS) {
		assert.False(t, storage.Exists(ctx, fs, "foo"))
		testutils.Create(t, fs, "foo", "bar")
		assert.True(t, storage.Exists(ctx, fs, "foo"))
	})
}

func TestRead(t *testing.T) {
	ctx := t.Context()
	withMem(func(fs storage.FS) {
		testutils.Create(t, fs, "foo", "bar")

		data, err := storage.Read(ctx, fs, "foo", nil)
		assert.NoError(t, err)
		assert.Equal(t, []byte("bar"), data)
	})
}

func TestWrite(t *testing.T) {
	ctx := t.Context()
	withMem(func(fs storage.FS) {
		assert.NoError(t, storage.Write(ctx, fs, "foo", []byte("bar"), nil))
		testutils.OpenExists(t, fs, "foo", "bar")
	})
}
