// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/user32"
)

type LITEM struct {
	Mask      uint32
	ILink     int32
	State     uint32
	StateMask uint32
	SzID      [MAX_LINKID_TEXT]uint16
	SzUrl     [L_MAX_URL_LENGTH]uint16
}

type LHITTESTINFO struct {
	Pt   gdi32.POINT
	Item LITEM
}

type NMLINK struct {
	Hdr  user32.NMHDR
	Item LITEM
}
