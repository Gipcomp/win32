// Copyright 2012 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shobj

import (
	"syscall"
	"unsafe"

	"github.com/gfphoenix/win32/handle"
	"github.com/gfphoenix/win32/user32"
	"github.com/gfphoenix/win32/win"
)

func (obj *ITaskbarList3) SetProgressState(hwnd handle.HWND, state int) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetProgressState, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(state))
	return win.HRESULT(ret)
}

func (obj *ITaskbarList3) SetOverlayIcon(hwnd handle.HWND, icon user32.HICON, description *uint16) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetOverlayIcon, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(icon),
		uintptr(unsafe.Pointer(description)),
		0,
		0)
	return win.HRESULT(ret)
}
