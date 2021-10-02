// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import "github.com/Gipcomp/winapi/user32"

type UDACCEL struct {
	NSec uint32
	NInc uint32
}

type NMUPDOWN struct {
	Hdr    user32.NMHDR
	IPos   int32
	IDelta int32
}
