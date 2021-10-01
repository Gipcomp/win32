// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package richole

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/kernel32"
	"github.com/D4v1dW3bb/winapi/objidl"
	"github.com/D4v1dW3bb/winapi/ole32"
	"github.com/D4v1dW3bb/winapi/richedit"
	"github.com/D4v1dW3bb/winapi/win"
)

func (obj *IRichEditOle) QueryInterface(riid ole32.REFIID, ppvObject *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.QueryInterface, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)))
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.AddRef, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *IRichEditOle) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *IRichEditOle) GetClientSite(lplpolesite **ole32.IOleClientSite) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetClientSite, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lplpolesite)),
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) GetObjectCount() int32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetObjectCount, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return int32(ret)
}

func (obj *IRichEditOle) GetLinkCount() int32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetLinkCount, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return int32(ret)
}

func (obj *IRichEditOle) GetObject(iob int32, lpreobject *REOBJECT, dwFlags uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.GetObject, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(unsafe.Pointer(lpreobject)),
		uintptr(dwFlags),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) InsertObject(lpreobject *REOBJECT) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.InsertObject, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpreobject)),
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) ConvertObject(iob int32, rclsidNew ole32.REFCLSID, lpstrUserTypeNew *byte) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.ConvertObject, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(unsafe.Pointer(rclsidNew)),
		uintptr(unsafe.Pointer(lpstrUserTypeNew)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) ActivateAs(rclsid ole32.REFCLSID, rclsidAs ole32.REFCLSID) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.ActivateAs, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(rclsidAs)))
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) SetHostNames(lpstrContainerApp *byte, lpstrContainerObj *byte) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetHostNames, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpstrContainerApp)),
		uintptr(unsafe.Pointer(lpstrContainerObj)))
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) SetLinkAvailable(iob int32, fAvailable win.BOOL) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetLinkAvailable, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(fAvailable))
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) SetDvaspect(iob int32, dvaspect uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetDvaspect, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(dvaspect))
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) HandsOffStorage(iob int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.HandsOffStorage, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) SaveCompleted(iob int32, lpstg *objidl.IStorage) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SaveCompleted, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(unsafe.Pointer(lpstg)))
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) InPlaceDeactivate() win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.InPlaceDeactivate, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) ContextSensitiveHelp(fEnterMode win.BOOL) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.ContextSensitiveHelp, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(fEnterMode),
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) GetClipboardData(lpchrg *richedit.CHARRANGE, reco uint32, lplpdataobj **objidl.IDataObject) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.GetClipboardData, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpchrg)),
		uintptr(reco),
		uintptr(unsafe.Pointer(lplpdataobj)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *IRichEditOle) ImportDataObject(lpdataobj *objidl.IDataObject, cf gdi32.CLIPFORMAT, hMetaPict kernel32.HGLOBAL) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.ImportDataObject, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpdataobj)),
		uintptr(cf),
		uintptr(hMetaPict),
		0,
		0)
	return win.HRESULT(ret)
}
