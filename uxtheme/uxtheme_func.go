// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package uxtheme

import (
	"syscall"
	"unsafe"

	"github.com/Gipcomp/winapi/gdi32"
	"github.com/Gipcomp/winapi/handle"
	"github.com/Gipcomp/winapi/win"
)

func CloseThemeData(hTheme HTHEME) win.HRESULT {
	ret, _, _ := syscall.Syscall(closeThemeData.Addr(), 1,
		uintptr(hTheme),
		0,
		0)

	return win.HRESULT(ret)
}

func DrawThemeBackground(hTheme HTHEME, hdc gdi32.HDC, iPartId, iStateId int32, pRect, pClipRect *gdi32.RECT) win.HRESULT {
	ret, _, _ := syscall.Syscall6(drawThemeBackground.Addr(), 6,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(unsafe.Pointer(pClipRect)))

	return win.HRESULT(ret)
}

func DrawThemeTextEx(hTheme HTHEME, hdc gdi32.HDC, iPartId, iStateId int32, pszText *uint16, iCharCount int32, dwFlags uint32, pRect *gdi32.RECT, pOptions *DTTOPTS) win.HRESULT {
	if drawThemeTextEx.Find() != nil {
		return win.HRESULT(0)
	}
	ret, _, _ := syscall.Syscall9(drawThemeTextEx.Addr(), 9,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pszText)),
		uintptr(iCharCount),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(pRect)),
		uintptr(unsafe.Pointer(pOptions)))

	return win.HRESULT(ret)
}

func GetThemeColor(hTheme HTHEME, iPartId, iStateId, iPropId int32, pColor *gdi32.COLORREF) win.HRESULT {
	ret, _, _ := syscall.Syscall6(getThemeColor.Addr(), 5,
		uintptr(hTheme),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(iPropId),
		uintptr(unsafe.Pointer(pColor)),
		0)

	return win.HRESULT(ret)
}

func GetThemePartSize(hTheme HTHEME, hdc gdi32.HDC, iPartId, iStateId int32, prc *gdi32.RECT, eSize THEMESIZE, psz *gdi32.SIZE) win.HRESULT {
	ret, _, _ := syscall.Syscall9(getThemePartSize.Addr(), 7,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(prc)),
		uintptr(eSize),
		uintptr(unsafe.Pointer(psz)),
		0,
		0)

	return win.HRESULT(ret)
}

func GetThemeTextExtent(hTheme HTHEME, hdc gdi32.HDC, iPartId, iStateId int32, pszText *uint16, iCharCount int32, dwTextFlags uint32, pBoundingRect, pExtentRect *gdi32.RECT) win.HRESULT {
	ret, _, _ := syscall.Syscall9(getThemeTextExtent.Addr(), 9,
		uintptr(hTheme),
		uintptr(hdc),
		uintptr(iPartId),
		uintptr(iStateId),
		uintptr(unsafe.Pointer(pszText)),
		uintptr(iCharCount),
		uintptr(dwTextFlags),
		uintptr(unsafe.Pointer(pBoundingRect)),
		uintptr(unsafe.Pointer(pExtentRect)))

	return win.HRESULT(ret)
}

func IsAppThemed() bool {
	ret, _, _ := syscall.Syscall(isAppThemed.Addr(), 0,
		0,
		0,
		0)

	return ret != 0
}

func OpenThemeData(hwnd handle.HWND, pszClassList *uint16) HTHEME {
	ret, _, _ := syscall.Syscall(openThemeData.Addr(), 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pszClassList)),
		0)

	return HTHEME(ret)
}

func SetWindowTheme(hwnd handle.HWND, pszSubAppName, pszSubIdList *uint16) win.HRESULT {
	ret, _, _ := syscall.Syscall(setWindowTheme.Addr(), 3,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pszSubAppName)),
		uintptr(unsafe.Pointer(pszSubIdList)))

	return win.HRESULT(ret)
}
