package bitcask_go

import "os"

type Options struct {
	DirPath string // 数据库数据目录

	DataFileSize int64 // 数据文件大小

	SyncWrites bool // 每次写数据是否持久化

	IndexType IndexerType // 索引类型
}

// IteratorOptions 索引迭代器配置项
type IteratorOptions struct {
	// 遍历前缀为指定的 Key，默认为空
	Prefix []byte
	// 是否反向遍历，默认 false 是正向
	Reverse bool
}

type IndexerType = int8

const (
	// BTree 索引
	BTree IndexerType = iota + 1

	// ART Adpative Radix Tree 自适应基数树索引
	ART
)

var DefaultOptions = Options{
	DirPath:      os.TempDir(),
	DataFileSize: 256 * 1024 * 1024,
	SyncWrites:   false,
	IndexType:    BTree,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}