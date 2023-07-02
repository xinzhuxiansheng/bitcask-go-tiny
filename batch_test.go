package bitcask_go

import (
	"bitcask-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDB_NewWriteBatch1(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("E:\\Code\\Go\\bitcask-go-tiny\\doc", "bitcask-go-batch-1")
	t.Log(dir)
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	wb := db.NewWriteBatch(DefaultWriteBatchOptions)
	err = wb.Put(utils.GetTestKey(1), utils.RandomValue(10))
	assert.Nil(t, err)
	err = wb.Delete(utils.GetTestKey(2))
	assert.Nil(t, err)

	_, err = db.Get(utils.GetTestKey(1))
	assert.Equal(t, ErrKeyNotFound, err)

	// 正常提交数据
	err = wb.Commit()
	assert.Nil(t, err)

	val1, err := db.Get(utils.GetTestKey(1))
	assert.NotNil(t, val1)
	assert.Nil(t, err)

	// 删除有效的数据
	wb2 := db.NewWriteBatch(DefaultWriteBatchOptions)
	err = wb2.Delete(utils.GetTestKey(1))
	assert.Nil(t, err)
	err = wb2.Commit()
	assert.Nil(t, err)

	val2, err := db.Get(utils.GetTestKey(1))
	t.Log(val2)
	t.Log(err)

}

func TestDB_NewWriteBatch2(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("E:\\Code\\Go\\bitcask-go-tiny\\doc", "bitcask-go-batch-2")
	t.Log(dir)
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	wb := db.NewWriteBatch(DefaultWriteBatchOptions)
	err = wb.Put(utils.GetTestKey(1), utils.RandomValue(10))
	assert.Nil(t, err)
	err = wb.Delete(utils.GetTestKey(1))
	assert.Nil(t, err)

	err = wb.Commit()
	assert.Nil(t, err)

	err = wb.Put(utils.GetTestKey(11), utils.RandomValue(10))
	assert.Nil(t, err)
	err = wb.Commit()
	assert.Nil(t, err)

	// 重启
	err = db.Close()
	assert.Nil(t, err)

	db2, err := Open(opts)
	assert.Nil(t, err)

	_, err = db2.Get(utils.GetTestKey(1))
	assert.Equal(t, ErrKeyNotFound, err)

	// 校验序列号
	assert.Equal(t, uint64(2), db.seqNo)
}

func TestDB_NewWriteBatch3(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("E:\\Code\\Go\\bitcask-go-tiny\\doc", "bitcask-go-batch-3")
	t.Log(dir)
	opts.DirPath = dir
	db, err := Open(opts)
	//defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	keys := db.ListKeys()
	t.Log(len(keys))

	wbOpts := DefaultWriteBatchOptions
	wbOpts.MaxBatchNum = 10000000
	wb := db.NewWriteBatch(DefaultWriteBatchOptions)
	for i := 0; i < 500000; i++ {
		err := wb.Put(utils.GetTestKey(i), utils.RandomValue(1024))
		assert.Nil(t, err)
	}
	err = wb.Commit()
	assert.Nil(t, err)

	//err = db.Close()
	//assert.Nil(t, err)

}

func TestDB_Start(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("E:\\Code\\Go\\bitcask-go-tiny\\doc", "bitcask-go-batch-3")
	t.Log(dir)
	opts.DirPath = dir
	db, err := Open(opts)
	//defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	//
}
