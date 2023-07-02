package utils

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestDirSize(t *testing.T) {
	dirSize, err := DirSize(filepath.Join("E:\\Code\\Go\\bitcask-go-tiny\\doc"))
	assert.Nil(t, err)
	t.Log(dirSize)
}

func TestAvailableDiskSize(t *testing.T) {
	size, err := AvailableDiskSize()
	assert.Nil(t, err)
	assert.True(t, size > 0)
}
