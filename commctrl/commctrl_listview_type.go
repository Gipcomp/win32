// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/gfphoenix/win32/comctl32"
	"github.com/gfphoenix/win32/gdi32"
	"github.com/gfphoenix/win32/user32"
)

type LVCOLUMN struct {
	Mask       uint32
	Fmt        int32
	Cx         int32
	PszText    *uint16
	CchTextMax int32
	ISubItem   int32
	IImage     int32
	IOrder     int32
}

type LVITEM struct {
	Mask       uint32
	IItem      int32
	ISubItem   int32
	State      uint32
	StateMask  uint32
	PszText    *uint16
	CchTextMax int32
	IImage     int32
	LParam     uintptr
	IIndent    int32
	IGroupId   int32
	CColumns   uint32
	PuColumns  uint32
}

type LVHITTESTINFO struct {
	Pt       gdi32.POINT
	Flags    uint32
	IItem    int32
	ISubItem int32
	IGroup   int32
}

type NMITEMACTIVATE struct {
	Hdr       user32.NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  gdi32.POINT
	LParam    uintptr
	UKeyFlags uint32
}

type NMLISTVIEW struct {
	Hdr       user32.NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  gdi32.POINT
	LParam    uintptr
}

type NMLVCUSTOMDRAW struct {
	Nmcd        comctl32.NMCUSTOMDRAW
	ClrText     gdi32.COLORREF
	ClrTextBk   gdi32.COLORREF
	ISubItem    int32
	DwItemType  uint32
	ClrFace     gdi32.COLORREF
	IIconEffect int32
	IIconPhase  int32
	IPartId     int32
	IStateId    int32
	RcText      gdi32.RECT
	UAlign      uint32
}

type NMLVDISPINFO struct {
	Hdr  user32.NMHDR
	Item LVITEM
}

type NMLVSCROLL struct {
	Hdr user32.NMHDR
	Dx  int32
	Dy  int32
}
