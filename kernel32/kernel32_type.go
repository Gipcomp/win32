// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package kernel32

import "github.com/Gipcomp/winapi/handle"

// To solve import cycle:
// HANDLE moved to handle.HANDLE

type (
	ATOM          uint16
	HGLOBAL       handle.HANDLE
	HINSTANCE     handle.HANDLE
	LCID          uint32
	LCTYPE        uint32
	LANGID        uint16
	HMODULE       uintptr
	HWINEVENTHOOK handle.HANDLE
	HRSRC         uintptr
)

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type NUMBERFMT struct {
	NumDigits     uint32
	LeadingZero   uint32
	Grouping      uint32
	LpDecimalSep  *uint16
	LpThousandSep *uint16
	NegativeOrder uint32
}

type SYSTEMTIME struct {
	WYear         uint16
	WMonth        uint16
	WDayOfWeek    uint16
	WDay          uint16
	WHour         uint16
	WMinute       uint16
	WSecond       uint16
	WMilliseconds uint16
}

type ACTCTX struct {
	size                  uint32
	Flags                 uint32
	Source                *uint16 // UTF-16 string
	ProcessorArchitecture uint16
	LangID                uint16
	AssemblyDirectory     *uint16 // UTF-16 string
	ResourceName          *uint16 // UTF-16 string
	ApplicationName       *uint16 // UTF-16 string
	Module                HMODULE
}
