// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/gfphoenix/win32/comctl32"
	"github.com/gfphoenix/win32/user32"
)

// Styles
const (
	SBARS_SIZEGRIP = 0x100
	SBARS_TOOLTIPS = 0x800
)

// Messages
const (
	SB_SETPARTS         = user32.WM_USER + 4
	SB_GETPARTS         = user32.WM_USER + 6
	SB_GETBORDERS       = user32.WM_USER + 7
	SB_SETMINHEIGHT     = user32.WM_USER + 8
	SB_SIMPLE           = user32.WM_USER + 9
	SB_GETRECT          = user32.WM_USER + 10
	SB_SETTEXT          = user32.WM_USER + 11
	SB_GETTEXTLENGTH    = user32.WM_USER + 12
	SB_GETTEXT          = user32.WM_USER + 13
	SB_ISSIMPLE         = user32.WM_USER + 14
	SB_SETICON          = user32.WM_USER + 15
	SB_SETTIPTEXT       = user32.WM_USER + 17
	SB_GETTIPTEXT       = user32.WM_USER + 19
	SB_GETICON          = user32.WM_USER + 20
	SB_SETUNICODEFORMAT = comctl32.CCM_SETUNICODEFORMAT
	SB_GETUNICODEFORMAT = comctl32.CCM_GETUNICODEFORMAT
	SB_SETBKCOLOR       = comctl32.CCM_SETBKCOLOR
)

// SB_SETTEXT options
const (
	SBT_NOBORDERS    = 0x100
	SBT_POPOUT       = 0x200
	SBT_RTLREADING   = 0x400
	SBT_NOTABPARSING = 0x800
	SBT_OWNERDRAW    = 0x1000
)

const (
	SBN_FIRST            = -880
	SBN_SIMPLEMODECHANGE = SBN_FIRST - 0
)

const SB_SIMPLEID = 0xff
