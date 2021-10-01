// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oleacc

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/handle"
	"github.com/D4v1dW3bb/winapi/ole32"
	"github.com/D4v1dW3bb/winapi/win"
	"github.com/D4v1dW3bb/winapi/winuser"
)

func (obj *IAccPropServices) QueryInterface(riid ole32.REFIID, ppvObject *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.QueryInterface, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)))
	return win.HRESULT(ret)
}

func (obj *IAccPropServices) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.AddRef, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *IAccPropServices) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

// SetPropServer specifies a callback object to be used to annotate an array of properties for the accessible element. You can also specify whether the annotation is to be applied to this accessible element or to the element and its children. This method is used for server annotation.
// If server developers know the handle.HWND of the accessible element they want to annotate, they can use SetHwndPropServer.
func (obj *IAccPropServices) SetPropServer(idString []byte, idProps []MSAAPROPID, server *IAccPropServer, annoScope AnnoScope) win.HRESULT {
	var idStringPtr unsafe.Pointer
	idStringLen := len(idString)
	if idStringLen != 0 {
		idStringPtr = unsafe.Pointer(&idString[0])
	}
	var idPropsPtr unsafe.Pointer
	idPropsLen := len(idProps)
	if idPropsLen != 0 {
		idPropsPtr = unsafe.Pointer(&idProps[0])
	}
	ret, _, _ := syscall.Syscall9(obj.LpVtbl.SetPropServer, 7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(idStringPtr),
		uintptr(idStringLen),
		uintptr(idPropsPtr),
		uintptr(idPropsLen),
		uintptr(unsafe.Pointer(server)),
		uintptr(annoScope),
		0,
		0)
	return win.HRESULT(ret)
}

// ClearProps restores default values to properties of accessible elements that they had previously annotated.
// If servers know the handle.HWND of the object they want to clear, they can use ClearHwndProps.
func (obj *IAccPropServices) ClearProps(idString []byte, idProps []MSAAPROPID) win.HRESULT {
	var idStringPtr unsafe.Pointer
	idStringLen := len(idString)
	if idStringLen != 0 {
		idStringPtr = unsafe.Pointer(&idString[0])
	}
	var idPropsPtr unsafe.Pointer
	idPropsLen := len(idProps)
	if idPropsLen != 0 {
		idPropsPtr = unsafe.Pointer(&idProps[0])
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.ClearProps, 5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(idStringPtr),
		uintptr(idStringLen),
		uintptr(idPropsPtr),
		uintptr(idPropsLen),
		0)
	return win.HRESULT(ret)
}

// SetHwndPropServer wraps SetPropServer, providing a convenient entry point for callers who are annotating handle.HWND-based accessible elements.
func (obj *IAccPropServices) SetHwndPropServer(hwnd handle.HWND, idObject int32, idChild uint32, idProps []MSAAPROPID, server *IAccPropServer, annoScope AnnoScope) win.HRESULT {
	var idPropsPtr unsafe.Pointer
	idPropsLen := len(idProps)
	if idPropsLen != 0 {
		idPropsPtr = unsafe.Pointer(&idProps[0])
	}
	ret, _, _ := syscall.Syscall9(obj.LpVtbl.SetHwndPropServer, 8,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idChild),
		uintptr(idPropsPtr),
		uintptr(idPropsLen),
		uintptr(unsafe.Pointer(server)),
		uintptr(annoScope),
		0)
	return win.HRESULT(ret)
}

// ClearHwndProps wraps SetPropValue, SetPropServer, and ClearProps, and provides a convenient entry point for callers who are annotating handle.HWND-based accessible elements.
func (obj *IAccPropServices) ClearHwndProps(hwnd handle.HWND, idObject int32, idChild uint32, idProps []MSAAPROPID) win.HRESULT {
	var idPropsPtr unsafe.Pointer
	idPropsLen := len(idProps)
	if idPropsLen != 0 {
		idPropsPtr = unsafe.Pointer(&idProps[0])
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.ClearHwndProps, 6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idChild),
		uintptr(idPropsPtr),
		uintptr(idPropsLen))
	return win.HRESULT(ret)
}

// ComposeHwndIdentityString retrievs an identity string.
func (obj *IAccPropServices) ComposeHwndIdentityString(hwnd handle.HWND, idObject int32, idChild uint32) (hr win.HRESULT, idString []byte) {
	var data *[1<<31 - 1]byte
	var len uint32
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.ComposeHwndIdentityString, 6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idChild),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(&len)))
	hr = win.HRESULT(ret)
	if win.FAILED(hr) {
		return
	}
	defer ole32.CoTaskMemFree(uintptr(unsafe.Pointer(data)))
	idString = make([]byte, len)
	copy(idString, data[:len])
	return
}

// DecomposeHwndIdentityString determines the handle.HWND, object ID, and child ID for the accessible element identified by the identity string.
func (obj *IAccPropServices) DecomposeHwndIdentityString(idString []byte) (hr win.HRESULT, hwnd handle.HWND, idObject int32, idChild uint32) {
	var idStringPtr unsafe.Pointer
	idStringLen := len(idString)
	if idStringLen != 0 {
		idStringPtr = unsafe.Pointer(&idString[0])
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.DecomposeHwndIdentityString, 6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(idStringPtr),
		uintptr(idStringLen),
		uintptr(unsafe.Pointer(&hwnd)),
		uintptr(unsafe.Pointer(&idObject)),
		uintptr(unsafe.Pointer(&idChild)))
	hr = win.HRESULT(ret)
	return
}

// SetHmenuPropServer wraps SetPropServer, providing a convenient entry point for callers who are annotating winuser.HMENU-based accessible elements.
func (obj *IAccPropServices) SetHmenuPropServer(hmenu winuser.HMENU, idChild uint32, idProps []MSAAPROPID, server *IAccPropServer, annoScope AnnoScope) win.HRESULT {
	var idPropsPtr unsafe.Pointer
	idPropsLen := len(idProps)
	if idPropsLen != 0 {
		idPropsPtr = unsafe.Pointer(&idProps[0])
	}
	ret, _, _ := syscall.Syscall9(obj.LpVtbl.SetHmenuPropServer, 7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hmenu),
		uintptr(idChild),
		uintptr(idPropsPtr),
		uintptr(idPropsLen),
		uintptr(unsafe.Pointer(server)),
		uintptr(annoScope),
		0,
		0)
	return win.HRESULT(ret)
}

// ClearHmenuProps wraps ClearProps, and provides a convenient entry point for callers who are annotating winuser.HMENU-based accessible elements.
func (obj *IAccPropServices) ClearHmenuProps(hmenu winuser.HMENU, idChild uint32, idProps []MSAAPROPID) win.HRESULT {
	var idPropsPtr unsafe.Pointer
	idPropsLen := len(idProps)
	if idPropsLen != 0 {
		idPropsPtr = unsafe.Pointer(&idProps[0])
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.ClearHmenuProps, 5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hmenu),
		uintptr(idChild),
		uintptr(idPropsPtr),
		uintptr(idPropsLen),
		0)
	return win.HRESULT(ret)
}

// ComposeHmenuIdentityString retrieves an identity string for an winuser.HMENU-based accessible element.
func (obj *IAccPropServices) ComposeHmenuIdentityString(hmenu winuser.HMENU, idChild uint32) (hr win.HRESULT, idString []byte) {
	var data *[1<<31 - 1]byte
	var len uint32
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.ComposeHmenuIdentityString, 5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hmenu),
		uintptr(idChild),
		uintptr(unsafe.Pointer(&data)),
		uintptr(unsafe.Pointer(&len)),
		0)
	hr = win.HRESULT(ret)
	if win.FAILED(hr) {
		return
	}
	defer ole32.CoTaskMemFree(uintptr(unsafe.Pointer(data)))
	idString = make([]byte, len)
	copy(idString, data[:len])
	return
}

// DecomposeHmenuIdentityString determines the winuser.HMENU, object ID, and child ID for the accessible element identified by the identity string.
func (obj *IAccPropServices) DecomposeHmenuIdentityString(idString []byte) (hr win.HRESULT, hmenu winuser.HMENU, idChild uint32) {
	var idStringPtr unsafe.Pointer
	idStringLen := len(idString)
	if idStringLen != 0 {
		idStringPtr = unsafe.Pointer(&idString[0])
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.DecomposeHmenuIdentityString, 5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(idStringPtr),
		uintptr(idStringLen),
		uintptr(unsafe.Pointer(&hmenu)),
		uintptr(unsafe.Pointer(&idChild)),
		0)
	hr = win.HRESULT(ret)
	return
}
