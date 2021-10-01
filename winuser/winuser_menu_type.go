// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winuser

import (
	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/handle"
)

// winuser.HMENU removed from user32
type HMENU handle.HANDLE

type MENUITEMINFO struct {
	CbSize        uint32
	FMask         uint32
	FType         uint32
	FState        uint32
	WID           uint32
	HSubMenu      HMENU
	HbmpChecked   gdi32.HBITMAP
	HbmpUnchecked gdi32.HBITMAP
	DwItemData    uintptr
	DwTypeData    *uint16
	Cch           uint32
	HbmpItem      gdi32.HBITMAP
}

type MENUINFO struct {
	CbSize          uint32
	FMask           uint32
	DwStyle         uint32
	CyMax           uint32
	HbrBack         gdi32.HBRUSH
	DwContextHelpID uint32
	DwMenuData      uintptr
}
