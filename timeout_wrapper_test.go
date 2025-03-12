package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coronon/go-storage"
)

func TestNewTimeoutWrapper(t *testing.T) {
	withSlowWrapper(slowDelay*2, slowDelay*2, func(fs storage.FS) {
		fs = storage.NewTimeoutWrapper(fs, slowDelay, slowDelay)
		file, err := fs.Open(t.Context(), "foo", nil)
		assert.Nil(t, file)
		assert.EqualError(t, err, "context deadline exceeded")
	})
}
