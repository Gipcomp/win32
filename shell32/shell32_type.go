// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shell32

import (
	"syscall"

	"github.com/Gipcomp/winapi/handle"
	"github.com/Gipcomp/winapi/kernel32"
	"github.com/Gipcomp/winapi/user32"
)

type CSIDL uint32
type HDROP handle.HANDLE

type NOTIFYICONDATA struct {
	CbSize           uint32
	HWnd             handle.HWND
	UID              uint32
	UFlags           uint32
	UCallbackMessage uint32
	HIcon            user32.HICON
	SzTip            [128]uint16
	DwState          uint32
	DwStateMask      uint32
	SzInfo           [256]uint16
	UVersion         uint32
	SzInfoTitle      [64]uint16
	DwInfoFlags      uint32
	GuidItem         syscall.GUID
	HBalloonIcon     user32.HICON
}

type SHFILEINFO struct {
	HIcon         user32.HICON
	IIcon         int32
	DwAttributes  uint32
	SzDisplayName [kernel32.MAX_PATH]uint16
	SzTypeName    [80]uint16
}

type BROWSEINFO struct {
	HwndOwner      handle.HWND
	PidlRoot       uintptr
	PszDisplayName *uint16
	LpszTitle      *uint16
	UlFlags        uint32
	Lpfn           uintptr
	LParam         uintptr
	IImage         int32
}

type SHSTOCKICONINFO struct {
	CbSize         uint32
	HIcon          user32.HICON
	ISysImageIndex int32
	IIcon          int32
	SzPath         [kernel32.MAX_PATH]uint16
}
