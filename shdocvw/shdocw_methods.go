// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shdocvw

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/handle"
	"github.com/D4v1dW3bb/winapi/ole32"
	"github.com/D4v1dW3bb/winapi/oleaut32"
	"github.com/D4v1dW3bb/winapi/user32"
	"github.com/D4v1dW3bb/winapi/win"
)

func (wb2 *IWebBrowser2) QueryInterface(riid ole32.REFIID, ppvObject *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.QueryInterface, 3,
		uintptr(unsafe.Pointer(wb2)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)))

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Release() win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(wb2)),
		0,
		0)

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Refresh() win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.Refresh, 1,
		uintptr(unsafe.Pointer(wb2)),
		0,
		0)

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Put_Left(Left int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.Put_Left, 2,
		uintptr(unsafe.Pointer(wb2)),
		uintptr(Left),
		0)

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Put_Top(Top int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.Put_Top, 2,
		uintptr(unsafe.Pointer(wb2)),
		uintptr(Top),
		0)

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Put_Width(Width int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.Put_Width, 2,
		uintptr(unsafe.Pointer(wb2)),
		uintptr(Width),
		0)

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Put_Height(Height int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.Put_Height, 2,
		uintptr(unsafe.Pointer(wb2)),
		uintptr(Height),
		0)

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Get_LocationURL(pbstrLocationURL **uint16 /*BSTR*/) win.HRESULT {
	ret, _, _ := syscall.Syscall(wb2.LpVtbl.Get_LocationURL, 2,
		uintptr(unsafe.Pointer(wb2)),
		uintptr(unsafe.Pointer(pbstrLocationURL)),
		0)

	return win.HRESULT(ret)
}

func (wb2 *IWebBrowser2) Navigate2(URL *oleaut32.VAR_BSTR, Flags *oleaut32.VAR_I4, TargetFrameName *oleaut32.VAR_BSTR, PostData unsafe.Pointer, Headers *oleaut32.VAR_BSTR) win.HRESULT {
	ret, _, _ := syscall.Syscall6(wb2.LpVtbl.Navigate2, 6,
		uintptr(unsafe.Pointer(wb2)),
		uintptr(unsafe.Pointer(URL)),
		uintptr(unsafe.Pointer(Flags)),
		uintptr(unsafe.Pointer(TargetFrameName)),
		uintptr(PostData),
		uintptr(unsafe.Pointer(Headers)))

	return win.HRESULT(ret)
}

func (activeObj *IOleInPlaceActiveObject) Release() win.HRESULT {
	ret, _, _ := syscall.Syscall(activeObj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(activeObj)),
		0,
		0)

	return win.HRESULT(ret)
}

func (activeObj *IOleInPlaceActiveObject) GetWindow(hWndPtr *handle.HWND) win.HRESULT {
	ret, _, _ := syscall.Syscall(activeObj.LpVtbl.GetWindow, 2,
		uintptr(unsafe.Pointer(activeObj)),
		uintptr(unsafe.Pointer(hWndPtr)),
		0)

	return win.HRESULT(ret)
}

func (activeObj *IOleInPlaceActiveObject) TranslateAccelerator(msg *user32.MSG) win.HRESULT {
	ret, _, _ := syscall.Syscall(activeObj.LpVtbl.TranslateAccelerator, 2,
		uintptr(unsafe.Pointer(activeObj)),
		uintptr(unsafe.Pointer(msg)),
		0)

	return win.HRESULT(ret)
}
