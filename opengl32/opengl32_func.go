// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package opengl32

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/win"
)

func WglCopyContext(hglrcSrc, hglrcDst HGLRC, mask uint) bool {
	ret, _, _ := syscall.Syscall(wglCopyContext.Addr(), 3,
		uintptr(hglrcSrc),
		uintptr(hglrcDst),
		uintptr(mask))

	return ret != 0
}

func WglCreateContext(hdc gdi32.HDC) HGLRC {
	ret, _, _ := syscall.Syscall(wglCreateContext.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return HGLRC(ret)
}

func WglCreateLayerContext(hdc gdi32.HDC, iLayerPlane int) HGLRC {
	ret, _, _ := syscall.Syscall(wglCreateLayerContext.Addr(), 2,
		uintptr(hdc),
		uintptr(iLayerPlane),
		0)

	return HGLRC(ret)
}

func WglDeleteContext(hglrc HGLRC) bool {
	ret, _, _ := syscall.Syscall(wglDeleteContext.Addr(), 1,
		uintptr(hglrc),
		0,
		0)

	return ret != 0
}

func WglDescribeLayerPlane(hdc gdi32.HDC, iPixelFormat, iLayerPlane int, nBytes uint8, plpd *LAYERPLANEDESCRIPTOR) bool {
	ret, _, _ := syscall.Syscall6(wglDescribeLayerPlane.Addr(), 5,
		uintptr(hdc),
		uintptr(iPixelFormat),
		uintptr(iLayerPlane),
		uintptr(nBytes),
		uintptr(unsafe.Pointer(plpd)),
		0)

	return ret != 0
}

func WglGetCurrentContext() HGLRC {
	ret, _, _ := syscall.Syscall(wglGetCurrentContext.Addr(), 0,
		0,
		0,
		0)

	return HGLRC(ret)
}

func WglGetCurrentDC() gdi32.HDC {
	ret, _, _ := syscall.Syscall(wglGetCurrentDC.Addr(), 0,
		0,
		0,
		0)

	return gdi32.HDC(ret)
}

func WglGetLayerPaletteEntries(hdc gdi32.HDC, iLayerPlane, iStart, cEntries int, pcr *gdi32.COLORREF) int {
	ret, _, _ := syscall.Syscall6(wglGetLayerPaletteEntries.Addr(), 5,
		uintptr(hdc),
		uintptr(iLayerPlane),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(pcr)),
		0)

	return int(ret)
}

func WglGetProcAddress(lpszProc *byte) uintptr {
	ret, _, _ := syscall.Syscall(wglGetProcAddress.Addr(), 1,
		uintptr(unsafe.Pointer(lpszProc)),
		0,
		0)

	return uintptr(ret)
}

func WglMakeCurrent(hdc gdi32.HDC, hglrc HGLRC) bool {
	ret, _, _ := syscall.Syscall(wglMakeCurrent.Addr(), 2,
		uintptr(hdc),
		uintptr(hglrc),
		0)

	return ret != 0
}

func WglRealizeLayerPalette(hdc gdi32.HDC, iLayerPlane int, bRealize bool) bool {
	ret, _, _ := syscall.Syscall(wglRealizeLayerPalette.Addr(), 3,
		uintptr(hdc),
		uintptr(iLayerPlane),
		uintptr(win.BoolToBOOL(bRealize)))

	return ret != 0
}

func WglSetLayerPaletteEntries(hdc gdi32.HDC, iLayerPlane, iStart, cEntries int, pcr *gdi32.COLORREF) int {
	ret, _, _ := syscall.Syscall6(wglSetLayerPaletteEntries.Addr(), 5,
		uintptr(hdc),
		uintptr(iLayerPlane),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(pcr)),
		0)

	return int(ret)
}

func WglShareLists(hglrc1, hglrc2 HGLRC) bool {
	ret, _, _ := syscall.Syscall(wglShareLists.Addr(), 2,
		uintptr(hglrc1),
		uintptr(hglrc2),
		0)

	return ret != 0
}

func WglSwapLayerBuffers(hdc gdi32.HDC, fuPlanes uint) bool {
	ret, _, _ := syscall.Syscall(wglSwapLayerBuffers.Addr(), 2,
		uintptr(hdc),
		uintptr(fuPlanes),
		0)

	return ret != 0
}

func WglUseFontBitmaps(hdc gdi32.HDC, first, count, listbase uint32) bool {
	ret, _, _ := syscall.Syscall6(wglUseFontBitmaps.Addr(), 4,
		uintptr(hdc),
		uintptr(first),
		uintptr(count),
		uintptr(listbase),
		0,
		0)

	return ret != 0
}

func WglUseFontOutlines(hdc gdi32.HDC, first, count, listbase uint32, deviation, extrusion float32, format int, pgmf *GLYPHMETRICSFLOAT) bool {
	ret, _, _ := syscall.Syscall12(wglUseFontBitmaps.Addr(), 8,
		uintptr(hdc),
		uintptr(first),
		uintptr(count),
		uintptr(listbase),
		uintptr(deviation),
		uintptr(extrusion),
		uintptr(format),
		uintptr(unsafe.Pointer(pgmf)),
		0,
		0,
		0,
		0)

	return ret != 0
}
