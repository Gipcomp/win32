// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import "github.com/Gipcomp/winapi/user32"

// ToolTip styles
const (
	TTS_ALWAYSTIP = 0x01
	TTS_NOPREFIX  = 0x02
	TTS_NOANIMATE = 0x10
	TTS_NOFADE    = 0x20
	TTS_BALLOON   = 0x40
	TTS_CLOSE     = 0x80
)

// ToolTip messages
const (
	TTM_ACTIVATE        = user32.WM_USER + 1
	TTM_SETDELAYTIME    = user32.WM_USER + 3
	TTM_ADDTOOL         = user32.WM_USER + 50
	TTM_DELTOOL         = user32.WM_USER + 51
	TTM_NEWTOOLRECT     = user32.WM_USER + 52
	TTM_RELAYEVENT      = user32.WM_USER + 7
	TTM_GETTOOLINFO     = user32.WM_USER + 53
	TTM_SETTOOLINFO     = user32.WM_USER + 54
	TTM_HITTEST         = user32.WM_USER + 55
	TTM_GETTEXT         = user32.WM_USER + 56
	TTM_UPDATETIPTEXT   = user32.WM_USER + 57
	TTM_GETTOOLCOUNT    = user32.WM_USER + 13
	TTM_ENUMTOOLS       = user32.WM_USER + 58
	TTM_GETCURRENTTOOL  = user32.WM_USER + 59
	TTM_WINDOWFROMPOINT = user32.WM_USER + 16
	TTM_TRACKACTIVATE   = user32.WM_USER + 17
	TTM_TRACKPOSITION   = user32.WM_USER + 18
	TTM_SETTIPBKCOLOR   = user32.WM_USER + 19
	TTM_SETTIPTEXTCOLOR = user32.WM_USER + 20
	TTM_GETDELAYTIME    = user32.WM_USER + 21
	TTM_GETTIPBKCOLOR   = user32.WM_USER + 22
	TTM_GETTIPTEXTCOLOR = user32.WM_USER + 23
	TTM_SETMAXTIPWIDTH  = user32.WM_USER + 24
	TTM_GETMAXTIPWIDTH  = user32.WM_USER + 25
	TTM_SETMARGIN       = user32.WM_USER + 26
	TTM_GETMARGIN       = user32.WM_USER + 27
	TTM_POP             = user32.WM_USER + 28
	TTM_UPDATE          = user32.WM_USER + 29
	TTM_GETBUBBLESIZE   = user32.WM_USER + 30
	TTM_ADJUSTRECT      = user32.WM_USER + 31
	TTM_SETTITLE        = user32.WM_USER + 33
	TTM_POPUP           = user32.WM_USER + 34
	TTM_GETTITLE        = user32.WM_USER + 35
)

// ToolTip flags
const (
	TTF_IDISHWND    = 0x0001
	TTF_CENTERTIP   = 0x0002
	TTF_RTLREADING  = 0x0004
	TTF_SUBCLASS    = 0x0010
	TTF_TRACK       = 0x0020
	TTF_ABSOLUTE    = 0x0080
	TTF_TRANSPARENT = 0x0100
	TTF_DI_SETITEM  = 0x8000
)

// ToolTip icons
const (
	TTI_NONE    = 0
	TTI_INFO    = 1
	TTI_WARNING = 2
	TTI_ERROR   = 3
)
