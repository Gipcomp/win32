// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package ole32

import (
	"syscall"
	"unsafe"

	"github.com/Gipcomp/win32/gdi32"
	"github.com/Gipcomp/win32/win"
)

func (cf *IClassFactory) Release() uint32 {
	ret, _, _ := syscall.Syscall(cf.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(cf)),
		0,
		0)

	return uint32(ret)
}

func (cf *IClassFactory) CreateInstance(pUnkOuter *IUnknown, riid REFIID, ppvObject *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall6(cf.LpVtbl.CreateInstance, 4,
		uintptr(unsafe.Pointer(cf)),
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)),
		0,
		0)

	return win.HRESULT(ret)
}

func (obj *IOleInPlaceObject) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)

	return uint32(ret)
}

func (obj *IOleInPlaceObject) SetObjectRects(lprcPosRect, lprcClipRect *gdi32.RECT) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetObjectRects, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lprcPosRect)),
		uintptr(unsafe.Pointer(lprcClipRect)))

	return win.HRESULT(ret)
}
