// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows,amd64

package oleacc

import (
	"syscall"
	"unsafe"

	"github.com/Gipcomp/win32/handle"
	"github.com/Gipcomp/win32/oleaut32"
	"github.com/Gipcomp/win32/win"
	"github.com/Gipcomp/win32/winuser"
)

// SetPropValue identifies the accessible element to be annotated, specify the property to be annotated, and provide a new value for that property.
// If server developers know the HWND of the accessible element they want to annotate, they can use one of the following methods: SetHwndPropStr, SetHwndProp, or SetHwndPropServer
func (obj *IAccPropServices) SetPropValue(idString []byte, idProp *MSAAPROPID, v *oleaut32.VARIANT) win.HRESULT {
	var idStringPtr unsafe.Pointer
	idStringLen := len(idString)
	if idStringLen != 0 {
		idStringPtr = unsafe.Pointer(&idString[0])
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetPropValue, 5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(idStringPtr),
		uintptr(idStringLen),
		uintptr(unsafe.Pointer(idProp)),
		uintptr(unsafe.Pointer(v)),
		0)
	return win.HRESULT(ret)
}

// SetHwndProp wraps SetPropValue, providing a convenient entry point for callers who are annotating HWND-based accessible elements. If the new value is a string, you can use SetHwndPropStr instead.
func (obj *IAccPropServices) SetHwndProp(hwnd handle.HWND, idObject int32, idChild uint32, idProp *MSAAPROPID, v *oleaut32.VARIANT) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetHwndProp, 6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idChild),
		uintptr(unsafe.Pointer(idProp)),
		uintptr(unsafe.Pointer(v)))
	return win.HRESULT(ret)
}

// SetHwndPropStr wraps SetPropValue, providing a more convenient entry point for callers who are annotating HWND-based accessible elements.
func (obj *IAccPropServices) SetHwndPropStr(hwnd handle.HWND, idObject int32, idChild uint32, idProp *MSAAPROPID, str string) win.HRESULT {
	str16, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return -((win.E_INVALIDARG ^ 0xFFFFFFFF) + 1)
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetHwndPropStr, 6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idChild),
		uintptr(unsafe.Pointer(idProp)),
		uintptr(unsafe.Pointer(str16)))
	return win.HRESULT(ret)
}

// SetHmenuProp wraps SetPropValue, providing a convenient entry point for callers who are annotating winuser.HMENU-based accessible elements. If the new value is a string, you can use IAccPropServices::SetHmenuPropStr instead.
func (obj *IAccPropServices) SetHmenuProp(hmenu winuser.HMENU, idChild uint32, idProp *MSAAPROPID, v *oleaut32.VARIANT) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetHmenuProp, 5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hmenu),
		uintptr(idChild),
		uintptr(unsafe.Pointer(idProp)),
		uintptr(unsafe.Pointer(v)),
		0)
	return win.HRESULT(ret)
}

// SetHmenuPropStr wraps SetPropValue, providing a more convenient entry point for callers who are annotating winuser.HMENU-based accessible elements.
func (obj *IAccPropServices) SetHmenuPropStr(hmenu winuser.HMENU, idChild uint32, idProp *MSAAPROPID, str string) win.HRESULT {
	str16, err := syscall.UTF16PtrFromString(str)
	if err != nil {
		return -((win.E_INVALIDARG ^ 0xFFFFFFFF) + 1)
	}
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.SetHmenuPropStr, 5,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hmenu),
		uintptr(idChild),
		uintptr(unsafe.Pointer(idProp)),
		uintptr(unsafe.Pointer(str16)),
		0)
	return win.HRESULT(ret)
}
