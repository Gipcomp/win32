// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package comdlg32

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/win"
)

func ChooseColor(lpcc *CHOOSECOLOR) bool {
	ret, _, _ := syscall.Syscall(chooseColor.Addr(), 1,
		uintptr(unsafe.Pointer(lpcc)),
		0,
		0)

	return ret != 0
}

func CommDlgExtendedError() uint32 {
	ret, _, _ := syscall.Syscall(commDlgExtendedError.Addr(), 0,
		0,
		0,
		0)

	return uint32(ret)
}

func GetOpenFileName(lpofn *OPENFILENAME) bool {
	ret, _, _ := syscall.Syscall(getOpenFileName.Addr(), 1,
		uintptr(unsafe.Pointer(lpofn)),
		0,
		0)

	return ret != 0
}

func GetSaveFileName(lpofn *OPENFILENAME) bool {
	ret, _, _ := syscall.Syscall(getSaveFileName.Addr(), 1,
		uintptr(unsafe.Pointer(lpofn)),
		0,
		0)

	return ret != 0
}

func PrintDlgEx(lppd *PRINTDLGEX) win.HRESULT {
	ret, _, _ := syscall.Syscall(printDlgEx.Addr(), 1,
		uintptr(unsafe.Pointer(lppd)),
		0,
		0)

	return win.HRESULT(ret)
}
