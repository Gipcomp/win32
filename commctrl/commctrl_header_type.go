// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/gfphoenix/win32/gdi32"
	"github.com/gfphoenix/win32/user32"
)

type HDITEM struct {
	Mask       uint32
	Cxy        int32
	PszText    *uint16
	Hbm        gdi32.HBITMAP
	CchTextMax int32
	Fmt        int32
	LParam     uintptr
	IImage     int32
	IOrder     int32
	Type       uint32
	PvFilter   uintptr
}

type HDLAYOUT struct {
	Prc   *gdi32.RECT
	Pwpos *user32.WINDOWPOS
}

type HDHITTESTINFO struct {
	Pt    gdi32.POINT
	Flags uint32
	IItem int32
}
