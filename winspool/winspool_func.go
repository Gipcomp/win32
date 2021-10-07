// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winspool

import (
	"syscall"
	"unsafe"

	"github.com/Gipcomp/win32/gdi32"
	"github.com/Gipcomp/win32/handle"
)

func DeviceCapabilities(pDevice, pPort *uint16, fwCapability uint16, pOutput *uint16, pDevMode *gdi32.DEVMODE) uint32 {
	ret, _, _ := syscall.Syscall6(deviceCapabilities.Addr(), 5,
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(unsafe.Pointer(pPort)),
		uintptr(fwCapability),
		uintptr(unsafe.Pointer(pOutput)),
		uintptr(unsafe.Pointer(pDevMode)),
		0)

	return uint32(ret)
}

func DocumentProperties(hWnd handle.HWND, hPrinter handle.HANDLE, pDeviceName *uint16, pDevModeOutput, pDevModeInput *gdi32.DEVMODE, fMode uint32) int32 {
	ret, _, _ := syscall.Syscall6(documentProperties.Addr(), 6,
		uintptr(hWnd),
		uintptr(hPrinter),
		uintptr(unsafe.Pointer(pDeviceName)),
		uintptr(unsafe.Pointer(pDevModeOutput)),
		uintptr(unsafe.Pointer(pDevModeInput)),
		uintptr(fMode))

	return int32(ret)
}

func EnumPrinters(Flags uint32, Name *uint16, Level uint32, pPrinterEnum *byte, cbBuf uint32, pcbNeeded, pcReturned *uint32) bool {
	ret, _, _ := syscall.Syscall9(enumPrinters.Addr(), 7,
		uintptr(Flags),
		uintptr(unsafe.Pointer(Name)),
		uintptr(Level),
		uintptr(unsafe.Pointer(pPrinterEnum)),
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(pcbNeeded)),
		uintptr(unsafe.Pointer(pcReturned)),
		0,
		0)

	return ret != 0
}

func GetDefaultPrinter(pszBuffer *uint16, pcchBuffer *uint32) bool {
	ret, _, _ := syscall.Syscall(getDefaultPrinter.Addr(), 2,
		uintptr(unsafe.Pointer(pszBuffer)),
		uintptr(unsafe.Pointer(pcchBuffer)),
		0)

	return ret != 0
}
