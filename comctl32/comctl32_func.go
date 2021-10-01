// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package comctl32

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/kernel32"
	"github.com/D4v1dW3bb/winapi/user32"
	"github.com/D4v1dW3bb/winapi/win"
)

func ImageList_Add(himl HIMAGELIST, hbmImage, hbmMask gdi32.HBITMAP) int32 {
	ret, _, _ := syscall.Syscall(imageList_Add.Addr(), 3,
		uintptr(himl),
		uintptr(hbmImage),
		uintptr(hbmMask))

	return int32(ret)
}

func ImageList_AddMasked(himl HIMAGELIST, hbmImage gdi32.HBITMAP, crMask gdi32.COLORREF) int32 {
	ret, _, _ := syscall.Syscall(imageList_AddMasked.Addr(), 3,
		uintptr(himl),
		uintptr(hbmImage),
		uintptr(crMask))

	return int32(ret)
}

func ImageList_Create(cx, cy int32, flags uint32, cInitial, cGrow int32) HIMAGELIST {
	ret, _, _ := syscall.Syscall6(imageList_Create.Addr(), 5,
		uintptr(cx),
		uintptr(cy),
		uintptr(flags),
		uintptr(cInitial),
		uintptr(cGrow),
		0)

	return HIMAGELIST(ret)
}

func ImageList_Destroy(hIml HIMAGELIST) bool {
	ret, _, _ := syscall.Syscall(imageList_Destroy.Addr(), 1,
		uintptr(hIml),
		0,
		0)

	return ret != 0
}

func ImageList_DrawEx(himl HIMAGELIST, i int32, hdcDst gdi32.HDC, x, y, dx, dy int32, rgbBk gdi32.COLORREF, rgbFg gdi32.COLORREF, fStyle uint32) bool {
	ret, _, _ := syscall.Syscall12(imageList_DrawEx.Addr(), 10,
		uintptr(himl),
		uintptr(i),
		uintptr(hdcDst),
		uintptr(x),
		uintptr(y),
		uintptr(dx),
		uintptr(dy),
		uintptr(rgbBk),
		uintptr(rgbFg),
		uintptr(fStyle),
		0,
		0)

	return ret != 0
}

func ImageList_ReplaceIcon(himl HIMAGELIST, i int32, hicon user32.HICON) int32 {
	ret, _, _ := syscall.Syscall(imageList_ReplaceIcon.Addr(), 3,
		uintptr(himl),
		uintptr(i),
		uintptr(hicon))

	return int32(ret)
}

func InitCommonControlsEx(lpInitCtrls *INITCOMMONCONTROLSEX) bool {
	ret, _, _ := syscall.Syscall(initCommonControlsEx.Addr(), 1,
		uintptr(unsafe.Pointer(lpInitCtrls)),
		0,
		0)

	return ret != 0
}

func LoadIconMetric(hInstance kernel32.HINSTANCE, lpIconName *uint16, lims int32, hicon *user32.HICON) win.HRESULT {
	if loadIconMetric.Find() != nil {
		return win.HRESULT(0)
	}
	ret, _, _ := syscall.Syscall6(loadIconMetric.Addr(), 4,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
		uintptr(lims),
		uintptr(unsafe.Pointer(hicon)),
		0,
		0)

	return win.HRESULT(ret)
}

func LoadIconWithScaleDown(hInstance kernel32.HINSTANCE, lpIconName *uint16, w int32, h int32, hicon *user32.HICON) win.HRESULT {
	if loadIconWithScaleDown.Find() != nil {
		return win.HRESULT(0)
	}
	ret, _, _ := syscall.Syscall6(loadIconWithScaleDown.Addr(), 5,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
		uintptr(w),
		uintptr(h),
		uintptr(unsafe.Pointer(hicon)),
		0)

	return win.HRESULT(ret)
}
