// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oleaut32

import (
	"syscall"
	"unsafe"
)

func SysAllocString(s string) *uint16 /*BSTR*/ {
	ret, _, _ := syscall.Syscall(sysAllocString.Addr(), 1,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s))),
		0,
		0)

	return (*uint16) /*BSTR*/ (unsafe.Pointer(ret))
}

func SysFreeString(bstr *uint16 /*BSTR*/) {
	syscall.Syscall(sysFreeString.Addr(), 1,
		uintptr(unsafe.Pointer(bstr)),
		0,
		0)
}

func SysStringLen(bstr *uint16 /*BSTR*/) uint32 {
	ret, _, _ := syscall.Syscall(sysStringLen.Addr(), 1,
		uintptr(unsafe.Pointer(bstr)),
		0,
		0)

	return uint32(ret)
}
