package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

//func destoryFile(name string) {
//	if err := os.RemoveAll(name); err != nil {
//		panic(err)
//	}
//}

func destoryFile(io *FileIO) {
	if io != nil {
		io.fd.Close()
	}
	if err := os.RemoveAll(io.fd.Name()); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	path := filepath.Join("E:\\Code\\Go\\bitcask-go\\doc", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(fio)

	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("E:\\Code\\Go\\bitcask-go\\doc", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(fio)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Write([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)

	n, err = fio.Write([]byte("bitcask kv"))
	t.Log(n, err)
	assert.Equal(t, 10, n)
	assert.Nil(t, err)
	n, err = fio.Write([]byte("storage"))
	assert.Equal(t, 7, n)
	assert.Nil(t, err)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("E:\\Code\\Go\\bitcask-go\\doc", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(fio)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("key-a"))
	assert.Nil(t, err)

	_, err = fio.Write([]byte("key-b"))
	assert.Nil(t, err)

	b := make([]byte, 5)
	n, err := fio.Read(b, 0)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b)

	b2 := make([]byte, 5)
	n, err = fio.Read(b2, 5)
	t.Log(b2, err)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("E:\\Code\\Go\\bitcask-go\\doc", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(fio)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Sync()
	assert.Nil(t, err)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("E:\\Code\\Go\\bitcask-go\\doc", "a.data")
	fio, err := NewFileIOManager(path)
	defer destoryFile(fio)

	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Sync()
	assert.Nil(t, err)
}
