// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package comdlg32

import (
	"unsafe"

	"github.com/Gipcomp/win32/gdi32"
	"github.com/Gipcomp/win32/handle"
	"github.com/Gipcomp/win32/kernel32"
)

type (
	LPOFNHOOKPROC  uintptr
	HPROPSHEETPAGE handle.HANDLE
	LPUNKNOWN      uintptr
)

type OPENFILENAME struct {
	LStructSize       uint32
	HwndOwner         handle.HWND
	HInstance         kernel32.HINSTANCE
	LpstrFilter       *uint16
	LpstrCustomFilter *uint16
	NMaxCustFilter    uint32
	NFilterIndex      uint32
	LpstrFile         *uint16
	NMaxFile          uint32
	LpstrFileTitle    *uint16
	NMaxFileTitle     uint32
	LpstrInitialDir   *uint16
	LpstrTitle        *uint16
	Flags             uint32
	NFileOffset       uint16
	NFileExtension    uint16
	LpstrDefExt       *uint16
	LCustData         uintptr
	LpfnHook          LPOFNHOOKPROC
	LpTemplateName    *uint16
	PvReserved        unsafe.Pointer
	DwReserved        uint32
	FlagsEx           uint32
}

type PRINTPAGERANGE struct {
	NFromPage uint32
	NToPage   uint32
}

type DEVNAMES struct {
	WDriverOffset uint16
	WDeviceOffset uint16
	WOutputOffset uint16
	WDefault      uint16
}

type PRINTDLGEX struct {
	LStructSize         uint32
	HwndOwner           handle.HWND
	HDevMode            kernel32.HGLOBAL
	HDevNames           kernel32.HGLOBAL
	HDC                 gdi32.HDC
	Flags               uint32
	Flags2              uint32
	ExclusionFlags      uint32
	NPageRanges         uint32
	NMaxPageRanges      uint32
	LpPageRanges        *PRINTPAGERANGE
	NMinPage            uint32
	NMaxPage            uint32
	NCopies             uint32
	HInstance           kernel32.HINSTANCE
	LpPrintTemplateName *uint16
	LpCallback          LPUNKNOWN
	NPropertyPages      uint32
	LphPropertyPages    *HPROPSHEETPAGE
	NStartPage          uint32
	DwResultAction      uint32
}

type CHOOSECOLOR struct {
	LStructSize    uint32
	HwndOwner      handle.HWND
	HInstance      handle.HWND
	RgbResult      gdi32.COLORREF
	LpCustColors   *[16]gdi32.COLORREF
	Flags          uint32
	LCustData      uintptr
	LpfnHook       uintptr
	LpTemplateName *uint16
}
