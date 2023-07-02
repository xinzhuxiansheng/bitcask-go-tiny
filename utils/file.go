package utils

import (
	_ "golang.org/x/sys/windows"
	"io/fs"
	"math"
	"path/filepath"
	_ "unsafe"
)

// DirSize 获取一个目录的大小
func DirSize(dirPath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// AvailableDiskSize 获取磁盘剩余可用空间大小
func AvailableDiskSize() (uint64, error) {
	//wd, err := syscall.Getwd()
	//if err != nil {
	//	return 0, err
	//}

	//kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	//getDiskFreeSpaceEx := kernel32.NewProc("GetDiskFreeSpaceExW")
	//
	//var freeBytesAvailableToCaller, totalNumberOfBytes, totalNumberOfFreeBytes int64
	//
	//r1, _, err := getDiskFreeSpaceEx.Call(
	//	uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("C:\\"))),
	//	uintptr(unsafe.Pointer(&freeBytesAvailableToCaller)),
	//	uintptr(unsafe.Pointer(&totalNumberOfBytes)),
	//	uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)),
	//)
	//
	//if r1 == 0 {
	//	fmt.Println("Call to GetDiskFreeSpaceExW failed:", err)
	//	return 0, err
	//}
	//
	//fmt.Printf("Free bytes available to caller: %d\n", freeBytesAvailableToCaller)
	//fmt.Printf("Total number of bytes: %d\n", totalNumberOfBytes)
	//fmt.Printf("Total number of free bytes: %d\n", totalNumberOfFreeBytes)

	return math.MaxInt64, nil
}
