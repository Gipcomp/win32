// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package ole32

import (
	"syscall"
	"unsafe"

	"github.com/Gipcomp/winapi/win"
)

func EqualREFIID(a, b REFIID) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if a.Data1 != b.Data1 || a.Data2 != b.Data2 || a.Data3 != b.Data3 {
		return false
	}

	for i := 0; i < 8; i++ {
		if a.Data4[i] != b.Data4[i] {
			return false
		}
	}

	return true
}

func CoCreateInstance(rclsid REFCLSID, pUnkOuter *IUnknown, dwClsContext uint32, riid REFIID, ppv *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall6(coCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppv)),
		0)

	return win.HRESULT(ret)
}

func CoGetClassObject(rclsid REFCLSID, dwClsContext uint32, pServerInfo *COSERVERINFO, riid REFIID, ppv *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall6(coGetClassObject.Addr(), 5,
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(pServerInfo)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppv)),
		0)

	return win.HRESULT(ret)
}

func CoInitializeEx(reserved unsafe.Pointer, coInit uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall(coInitializeEx.Addr(), 2,
		uintptr(reserved),
		uintptr(coInit),
		0)

	return win.HRESULT(ret)
}

func CoUninitialize() {
	syscall.Syscall(coUninitialize.Addr(), 0,
		0,
		0,
		0)
}

func CoTaskMemFree(pv uintptr) {
	syscall.Syscall(coTaskMemFree.Addr(), 1,
		pv,
		0,
		0)
}

func OleInitialize() win.HRESULT {
	ret, _, _ := syscall.Syscall(oleInitialize.Addr(), 1, // WTF, why does 0 not work here?
		0,
		0,
		0)

	return win.HRESULT(ret)
}

func OleSetContainedObject(pUnknown *IUnknown, fContained bool) win.HRESULT {
	ret, _, _ := syscall.Syscall(oleSetContainedObject.Addr(), 2,
		uintptr(unsafe.Pointer(pUnknown)),
		uintptr(win.BoolToBOOL(fContained)),
		0)

	return win.HRESULT(ret)
}

func OleUninitialize() {
	syscall.Syscall(oleUninitialize.Addr(), 0,
		0,
		0,
		0)
}
