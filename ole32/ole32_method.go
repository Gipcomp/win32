// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package ole32

import (
	"syscall"
	"unsafe"

	"github.com/gfphoenix/win32/gdi32"
	"github.com/gfphoenix/win32/handle"
	"github.com/gfphoenix/win32/user32"
	"github.com/gfphoenix/win32/win"
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

func (obj *IOleObject) QueryInterface(riid REFIID, ppvObject *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.QueryInterface, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)))

	return win.HRESULT(ret)
}

func (obj *IOleObject) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)

	return uint32(ret)
}

func (obj *IOleObject) SetClientSite(pClientSite *IOleClientSite) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetClientSite, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pClientSite)),
		0)

	return win.HRESULT(ret)
}

func (obj *IOleObject) SetHostNames(szContainerApp, szContainerObj *uint16) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetHostNames, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(szContainerApp)),
		uintptr(unsafe.Pointer(szContainerObj)))

	return win.HRESULT(ret)
}

func (obj *IOleObject) Close(dwSaveOption uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Close, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(dwSaveOption),
		0)

	return win.HRESULT(ret)
}

func (obj *IOleObject) DoVerb(iVerb int32, lpmsg *user32.MSG, pActiveSite *IOleClientSite, lindex int32, hwndParent handle.HWND, lprcPosRect *gdi32.RECT) win.HRESULT {
	ret, _, _ := syscall.Syscall9(obj.LpVtbl.DoVerb, 7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iVerb),
		uintptr(unsafe.Pointer(lpmsg)),
		uintptr(unsafe.Pointer(pActiveSite)),
		uintptr(lindex),
		uintptr(hwndParent),
		uintptr(unsafe.Pointer(lprcPosRect)),
		0,
		0)

	return win.HRESULT(ret)
}

func (cpc *IConnectionPointContainer) Release() uint32 {
	ret, _, _ := syscall.Syscall(cpc.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(cpc)),
		0,
		0)

	return uint32(ret)
}

func (cpc *IConnectionPointContainer) FindConnectionPoint(riid REFIID, ppCP **IConnectionPoint) win.HRESULT {
	ret, _, _ := syscall.Syscall(cpc.LpVtbl.FindConnectionPoint, 3,
		uintptr(unsafe.Pointer(cpc)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppCP)))

	return win.HRESULT(ret)
}

func (cp *IConnectionPoint) Release() uint32 {
	ret, _, _ := syscall.Syscall(cp.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(cp)),
		0,
		0)

	return uint32(ret)
}

func (cp *IConnectionPoint) Advise(pUnkSink unsafe.Pointer, pdwCookie *uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall(cp.LpVtbl.Advise, 3,
		uintptr(unsafe.Pointer(cp)),
		uintptr(pUnkSink),
		uintptr(unsafe.Pointer(pdwCookie)))

	return win.HRESULT(ret)
}
