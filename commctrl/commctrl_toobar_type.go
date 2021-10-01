// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/user32"
)

type NMMOUSE struct {
	Hdr        user32.NMHDR
	DwItemSpec uintptr
	DwItemData uintptr
	Pt         gdi32.POINT
	DwHitInfo  uintptr
}

type NMTOOLBAR struct {
	Hdr      user32.NMHDR
	IItem    int32
	TbButton TBBUTTON
	CchText  int32
	PszText  *uint16
	RcButton gdi32.RECT
}

type TBBUTTON struct {
	IBitmap   int32
	IdCommand int32
	FsState   byte
	FsStyle   byte
	//#ifdef _WIN64
	//    BYTE bReserved[6]          // padding for alignment
	//#elif defined(_WIN32)
	BReserved [2]byte // padding for alignment
	//#endif
	DwData  uintptr
	IString uintptr
}

type TBBUTTONINFO struct {
	CbSize    uint32
	DwMask    uint32
	IdCommand int32
	IImage    int32
	FsState   byte
	FsStyle   byte
	Cx        uint16
	LParam    uintptr
	PszText   uintptr
	CchText   int32
}

type TBMETRICS struct {
	CbSize          uint32
	DwMask          uint32
	CxPad           int32
	CyPad           int32
	CxBarPad        int32
	CyBarPad        int32
	CxButtonSpacing int32
	CyButtonSpacing int32
}
