// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package commctrl

import (
	"unsafe"

	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/handle"
	"github.com/D4v1dW3bb/winapi/kernel32"
)

type TOOLINFO struct {
	CbSize     uint32
	UFlags     uint32
	Hwnd       handle.HWND
	UId        uintptr
	Rect       gdi32.RECT
	Hinst      kernel32.HINSTANCE
	LpszText   *uint16
	LParam     uintptr
	LpReserved unsafe.Pointer
}

type TTGETTITLE struct {
	DwSize       uint32
	UTitleBitmap uint32
	Cch          uint32
	PszTitle     *uint16
}
