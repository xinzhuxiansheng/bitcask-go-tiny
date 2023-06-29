package bitcask_go

import "os"

type Options struct {
	DirPath string // 数据库数据目录

	DataFileSize int64 // 数据文件大小

	SyncWrites bool // 每次写数据是否持久化

	BytesPerSync uint // 累计写到多少字节后进行持久化

	IndexType IndexerType // 索引类型

	// 启动时是否使用 MMap 加载数据
	MMapAtStartup bool
}

// WriteBatchOptions 批量写配置项
type WriteBatchOptions struct {
	// 一个批次当中最大的数据量
	MaxBatchNum uint
	// 提交事务时是否 sync 持久化
	SyncWrites bool
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

	// BPlusTree B+ 树索引
	BPlusTree
)

var DefaultOptions = Options{
	DirPath:       os.TempDir(),
	DataFileSize:  256 * 1024 * 1024,
	SyncWrites:    false,
	BytesPerSync:  0,
	IndexType:     BTree,
	MMapAtStartup: true,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}

var DefaultWriteBatchOptions = WriteBatchOptions{
	MaxBatchNum: 10000,
	SyncWrites:  true,
}
