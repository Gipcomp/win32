// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/gfphoenix/win32/comctl32"
	"github.com/gfphoenix/win32/user32"
)

const UDN_FIRST = ^uint32(720)

const (
	UD_MAXVAL = 0x7fff
	UD_MINVAL = ^uintptr(UD_MAXVAL - 1)
)

const (
	UDS_WRAP        = 0x0001
	UDS_SETBUDDYINT = 0x0002
	UDS_ALIGNRIGHT  = 0x0004
	UDS_ALIGNLEFT   = 0x0008
	UDS_AUTOBUDDY   = 0x0010
	UDS_ARROWKEYS   = 0x0020
	UDS_HORZ        = 0x0040
	UDS_NOTHOUSANDS = 0x0080
	UDS_HOTTRACK    = 0x0100
)

const (
	UDM_SETRANGE         = user32.WM_USER + 101
	UDM_GETRANGE         = user32.WM_USER + 102
	UDM_SETPOS           = user32.WM_USER + 103
	UDM_GETPOS           = user32.WM_USER + 104
	UDM_SETBUDDY         = user32.WM_USER + 105
	UDM_GETBUDDY         = user32.WM_USER + 106
	UDM_SETACCEL         = user32.WM_USER + 107
	UDM_GETACCEL         = user32.WM_USER + 108
	UDM_SETBASE          = user32.WM_USER + 109
	UDM_GETBASE          = user32.WM_USER + 110
	UDM_SETRANGE32       = user32.WM_USER + 111
	UDM_GETRANGE32       = user32.WM_USER + 112
	UDM_SETUNICODEFORMAT = comctl32.CCM_SETUNICODEFORMAT
	UDM_GETUNICODEFORMAT = comctl32.CCM_GETUNICODEFORMAT
	UDM_SETPOS32         = user32.WM_USER + 113
	UDM_GETPOS32         = user32.WM_USER + 114
)

const UDN_DELTAPOS = UDN_FIRST - 1
