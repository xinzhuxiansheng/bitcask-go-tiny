package bitcask_go

import "errors"

var (
	ErrKeyIsEmpty             = errors.New("the key is empty")
	ErrIndexUpdateFailed      = errors.New("failed to update index")
	ErrKeyNotFound            = errors.New("key not found in database")
	ErrDataFileNotFound       = errors.New("data file is not found")
	ErrDataDircetoryCorrupted = errors.New("the database directory maybe corrupted")
	ErrMergeIsProgress        = errors.New("merge is in progress, try again later")
	ErrExceedMaxBatchNum      = errors.New("exceed the max batch num")
	ErrDatabaseIsUsing        = errors.New("the database directory is used by another process")
	ErrMergeRatioUnreached    = errors.New("the merge ratio do not reach the option")
	ErrNoEnoughSpaceForMerge  = errors.New("no enouth disk space for merge")
)
