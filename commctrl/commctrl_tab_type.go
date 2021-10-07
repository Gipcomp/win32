// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/Gipcomp/win32/gdi32"
	"github.com/Gipcomp/win32/user32"
)

type TCITEMHEADER struct {
	Mask        uint32
	LpReserved1 uint32
	LpReserved2 uint32
	PszText     *uint16
	CchTextMax  int32
	IImage      int32
}

type TCITEM struct {
	Mask        uint32
	DwState     uint32
	DwStateMask uint32
	PszText     *uint16
	CchTextMax  int32
	IImage      int32
	LParam      uintptr
}

type TCHITTESTINFO struct {
	Pt    gdi32.POINT
	flags uint32
}

type NMTCKEYDOWN struct {
	Hdr   user32.NMHDR
	WVKey uint16
	Flags uint32
}
