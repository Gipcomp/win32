// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package comctl32

import (
	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/handle"
	"github.com/D4v1dW3bb/winapi/user32"
)

type HIMAGELIST handle.HANDLE

type INITCOMMONCONTROLSEX struct {
	DwSize, DwICC uint32
}

type NMCUSTOMDRAW struct {
	Hdr         user32.NMHDR
	DwDrawStage uint32
	Hdc         gdi32.HDC
	Rc          gdi32.RECT
	DwItemSpec  uintptr
	UItemState  uint32
	LItemlParam uintptr
}
