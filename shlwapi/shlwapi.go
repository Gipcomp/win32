package shlwapi

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// Library
	libshlwapi *windows.LazyDLL

	// Functions
	shCreateMemStream *windows.LazyProc
)

func init() {
	// Initalize libary
	libshlwapi = windows.NewLazySystemDLL("shlwapi.dll")

	// Initialize Functions
	shCreateMemStream = libshlwapi.NewProc("SHCreateMemStream")
}

func SHCreateMemStream(data []byte) (uintptr, error) {
	ret, _, err := shCreateMemStream.Call(
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
	)
	if ret == 0 {
		return 0, err
	}

	return ret, nil
}
