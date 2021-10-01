// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package kernel32

const MAX_PATH = 260

// Error codes
const (
	ERROR_SUCCESS             = 0
	ERROR_INVALID_FUNCTION    = 1
	ERROR_FILE_NOT_FOUND      = 2
	ERROR_INVALID_PARAMETER   = 87
	ERROR_INSUFFICIENT_BUFFER = 122
	ERROR_MORE_DATA           = 234
)

// GlobalAlloc flags
const (
	GHND          = 0x0042
	GMEM_FIXED    = 0x0000
	GMEM_MOVEABLE = 0x0002
	GMEM_ZEROINIT = 0x0040
	GPTR          = GMEM_FIXED | GMEM_ZEROINIT
)

// Predefined locale ids
const (
	LOCALE_CUSTOM_DEFAULT     LCID = 0x0c00
	LOCALE_CUSTOM_UI_DEFAULT  LCID = 0x1400
	LOCALE_CUSTOM_UNSPECIFIED LCID = 0x1000
	LOCALE_INVARIANT          LCID = 0x007f
	LOCALE_USER_DEFAULT       LCID = 0x0400
	LOCALE_SYSTEM_DEFAULT     LCID = 0x0800
)

// LCTYPE constants
const (
	LOCALE_SDECIMAL          LCTYPE = 14
	LOCALE_STHOUSAND         LCTYPE = 15
	LOCALE_SISO3166CTRYNAME  LCTYPE = 0x5a
	LOCALE_SISO3166CTRYNAME2 LCTYPE = 0x68
	LOCALE_SISO639LANGNAME   LCTYPE = 0x59
	LOCALE_SISO639LANGNAME2  LCTYPE = 0x67
)
