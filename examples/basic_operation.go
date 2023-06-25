package examples

import (
	bitcask "bitcask-go"
	"fmt"
)

func main() {
	ops := bitcask.DefaultOptions
	ops.DirPath = "E:\\Code\\Go\\bitcask-go\\doc"
	db, err := bitcask.Open(ops)
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("name"), []byte("bitcask"))
	if err != nil {
		panic(err)
	}
	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("val = ", string(val))

	err = db.Delete([]byte("name"))
	if err != nil {
		panic(err)
	}
}
