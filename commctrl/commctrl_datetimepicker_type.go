// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"github.com/gfphoenix/win32/gdi32"
	"github.com/gfphoenix/win32/kernel32"
	"github.com/gfphoenix/win32/user32"
)

type (
	NMDATETIMECHANGE struct {
		Nmhdr   user32.NMHDR
		DwFlags uint32
		St      kernel32.SYSTEMTIME
	}

	NMDATETIMESTRING struct {
		Nmhdr         user32.NMHDR
		PszUserString *uint16
		St            kernel32.SYSTEMTIME
		DwFlags       uint32
	}

	NMDATETIMEWMKEYDOWN struct {
		Nmhdr     user32.NMHDR
		NVirtKey  int
		PszFormat *uint16
		St        kernel32.SYSTEMTIME
	}

	NMDATETIMEFORMAT struct {
		Nmhdr      user32.NMHDR
		PszFormat  *uint16
		St         kernel32.SYSTEMTIME
		PszDisplay *uint16
		SzDisplay  [64]uint16
	}

	NMDATETIMEFORMATQUERY struct {
		Nmhdr     user32.NMHDR
		PszFormat *uint16
		SzMax     gdi32.SIZE
	}
)
