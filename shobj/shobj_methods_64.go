// Copyright 2012 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows,amd64 windows,arm64

package shobj

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/handle"
	"github.com/D4v1dW3bb/winapi/win"
)

func (obj *ITaskbarList3) SetProgressValue(hwnd handle.HWND, current uint32, length uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetProgressValue, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(current),
		uintptr(length),
		0,
		0)

	return win.HRESULT(ret)
}
