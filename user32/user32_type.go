// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package user32

import (
	"github.com/gfphoenix/win32/gdi32"
	"github.com/gfphoenix/win32/handle"
	"github.com/gfphoenix/win32/kernel32"
	"github.com/gfphoenix/win32/win"
	"github.com/gfphoenix/win32/winuser"
)

type NMBCDROPDOWN struct {
	Hdr      NMHDR
	RcButton gdi32.RECT
}

type MONITORINFO struct {
	CbSize    uint32
	RcMonitor gdi32.RECT
	RcWork    gdi32.RECT
	DwFlags   uint32
}

// To solve import cycle:
// HWND moved to handle.HWND
// HMENU moved to winuser.HMENU

type (
	HACCEL    handle.HANDLE
	HCURSOR   handle.HANDLE
	HDWP      handle.HANDLE
	HICON     handle.HANDLE
	HMONITOR  handle.HANDLE
	HRAWINPUT handle.HANDLE
)

type MSG struct {
	HWnd    handle.HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      gdi32.POINT
}

type RAWINPUTDEVICE struct {
	UsUsagePage uint16
	UsUsage     uint16
	DwFlags     uint32
	HwndTarget  handle.HWND
}

type RAWINPUTHEADER struct {
	DwType  uint32
	DwSize  uint32
	HDevice handle.HANDLE
	WParam  uintptr
}

type RAWINPUTMOUSE struct {
	Header RAWINPUTHEADER
	Data   RAWMOUSE
}

type RAWINPUTKEYBOARD struct {
	Header RAWINPUTHEADER
	Data   RAWKEYBOARD
}

type RAWINPUTHID struct {
	Header RAWINPUTHEADER
	Data   RAWHID
}

type RAWMOUSE struct {
	UsFlags            uint16
	UsButtonFlags      uint16
	UsButtonData       uint16
	Pad_cgo_0          [2]byte
	UlRawButtons       uint32
	LLastX             int32
	LLastY             int32
	UlExtraInformation uint32
}

type RAWKEYBOARD struct {
	MakeCode         uint16
	Flags            uint16
	Reserved         int16
	VKey             uint16
	Message          uint32
	ExtraInformation uint32
}

type RAWHID struct {
	DwSizeHid uint32
	DwCount   uint32
	BRawData  [1]byte
}

type NMHDR struct {
	HwndFrom handle.HWND
	IdFrom   uintptr
	Code     uint32
}

type CREATESTRUCT struct {
	CreateParams    uintptr
	Instance        kernel32.HINSTANCE
	Menu            winuser.HMENU
	Parent          handle.HWND
	Cy              int32
	Cx              int32
	Y               int32
	X               int32
	Style           int32
	Name, ClassName uintptr
	ExStyle         uint32
}

type CHANGEFILTERSTRUCT struct {
	size      uint32
	extStatus uint32
}

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     kernel32.HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground gdi32.HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

type TPMPARAMS struct {
	CbSize    uint32
	RcExclude gdi32.RECT
}

type WINDOWPLACEMENT struct {
	Length           uint32
	Flags            uint32
	ShowCmd          uint32
	PtMinPosition    gdi32.POINT
	PtMaxPosition    gdi32.POINT
	RcNormalPosition gdi32.RECT
}

type DRAWTEXTPARAMS struct {
	CbSize        uint32
	ITabLength    int32
	ILeftMargin   int32
	IRightMargin  int32
	UiLengthDrawn uint32
}

type PAINTSTRUCT struct {
	Hdc         gdi32.HDC
	FErase      win.BOOL
	RcPaint     gdi32.RECT
	FRestore    win.BOOL
	FIncUpdate  win.BOOL
	RgbReserved [32]byte
}

type MINMAXINFO struct {
	PtReserved     gdi32.POINT
	PtMaxSize      gdi32.POINT
	PtMaxPosition  gdi32.POINT
	PtMinTrackSize gdi32.POINT
	PtMaxTrackSize gdi32.POINT
}

type NONCLIENTMETRICS struct {
	CbSize           uint32
	IBorderWidth     int32
	IScrollWidth     int32
	IScrollHeight    int32
	ICaptionWidth    int32
	ICaptionHeight   int32
	LfCaptionFont    gdi32.LOGFONT
	ISmCaptionWidth  int32
	ISmCaptionHeight int32
	LfSmCaptionFont  gdi32.LOGFONT
	IMenuWidth       int32
	IMenuHeight      int32
	LfMenuFont       gdi32.LOGFONT
	LfStatusFont     gdi32.LOGFONT
	LfMessageFont    gdi32.LOGFONT
}

type MEASUREITEMSTRUCT struct {
	CtlType    uint32
	CtlID      uint32
	ItemID     int32
	ItemWidth  uint32
	ItemHeight uint32
	ItemData   uintptr
}

type DRAWITEMSTRUCT struct {
	CtlType    uint32
	CtlID      uint32
	ItemID     int32
	ItemAction uint32
	ItemState  uint32
	HwndItem   handle.HWND
	HDC        gdi32.HDC
	RcItem     gdi32.RECT
	ItemData   uintptr
}

type ICONINFO struct {
	FIcon    win.BOOL
	XHotspot uint32
	YHotspot uint32
	HbmMask  gdi32.HBITMAP
	HbmColor gdi32.HBITMAP
}

type MOUSE_INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
}

type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

type KEYBD_INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
	Unused      [8]byte
}

type HARDWARE_INPUT struct {
	Type uint32
	Hi   HARDWAREINPUT
}

type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
	Unused  [16]byte
}

type SCROLLINFO struct {
	CbSize    uint32
	FMask     uint32
	NMin      int32
	NMax      int32
	NPage     uint32
	NPos      int32
	NTrackPos int32
}

type WINDOWPOS struct {
	Hwnd            handle.HWND
	HwndInsertAfter handle.HWND
	X               int32
	Y               int32
	Cx              int32
	Cy              int32
	Flags           uint32
}

type TRACKMOUSEEVENT struct {
	CbSize      uint32
	DwFlags     uint32
	HwndTrack   handle.HWND
	DwHoverTime uint32
}

type HIGHCONTRAST struct {
	CbSize            uint32
	DwFlags           uint32
	LpszDefaultScheme *uint16
}
