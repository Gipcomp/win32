// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oleaut32

import (
	"github.com/gfphoenix/win32/ole32"
	"golang.org/x/sys/windows"
)

var (
	IID_IDispatch = ole32.IID{0x00020400, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
)

var (
	// Library
	liboleaut32 *windows.LazyDLL

	// Functions
	sysAllocString *windows.LazyProc
	sysFreeString  *windows.LazyProc
	sysStringLen   *windows.LazyProc
)

func init() {
	// Library
	liboleaut32 = windows.NewLazySystemDLL("oleaut32.dll")

	// Functions
	sysAllocString = liboleaut32.NewProc("SysAllocString")
	sysFreeString = liboleaut32.NewProc("SysFreeString")
	sysStringLen = liboleaut32.NewProc("SysStringLen")
}
