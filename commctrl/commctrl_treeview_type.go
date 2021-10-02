// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/Gipcomp/winapi/gdi32"
	"github.com/Gipcomp/winapi/handle"
	"github.com/Gipcomp/winapi/user32"
)

type HTREEITEM handle.HANDLE

type TVITEM struct {
	Mask           uint32
	HItem          HTREEITEM
	State          uint32
	StateMask      uint32
	PszText        uintptr
	CchTextMax     int32
	IImage         int32
	ISelectedImage int32
	CChildren      int32
	LParam         uintptr
}

/*type TVITEMEX struct {
	mask           UINT
	hItem          HTREEITEM
	state          UINT
	stateMask      UINT
	pszText        LPWSTR
	cchTextMax     int
	iImage         int
	iSelectedImage int
	cChildren      int
	lParam         LPARAM
	iIntegral      int
	uStateEx       UINT
	hwnd           HWND
	iExpandedImage int
}*/

type TVINSERTSTRUCT struct {
	HParent      HTREEITEM
	HInsertAfter HTREEITEM
	Item         TVITEM
	//	itemex       TVITEMEX
}

type NMTREEVIEW struct {
	Hdr     user32.NMHDR
	Action  uint32
	ItemOld TVITEM
	ItemNew TVITEM
	PtDrag  gdi32.POINT
}

type NMTVDISPINFO struct {
	Hdr  user32.NMHDR
	Item TVITEM
}

type NMTVKEYDOWN struct {
	Hdr   user32.NMHDR
	WVKey uint16
	Flags uint32
}

type TVHITTESTINFO struct {
	Pt    gdi32.POINT
	Flags uint32
	HItem HTREEITEM
}
