// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package uxtheme

import (
	"github.com/Gipcomp/win32/gdi32"
	"github.com/Gipcomp/win32/handle"
	"github.com/Gipcomp/win32/win"
)

type HTHEME handle.HANDLE

type THEMESIZE int

type DTTOPTS struct {
	DwSize              uint32
	DwFlags             uint32
	CrText              gdi32.COLORREF
	CrBorder            gdi32.COLORREF
	CrShadow            gdi32.COLORREF
	ITextShadowType     int32
	PtShadowOffset      gdi32.POINT
	IBorderSize         int32
	IFontPropId         int32
	IColorPropId        int32
	IStateId            int32
	FApplyOverlay       win.BOOL
	IGlowSize           int32
	PfnDrawTextCallback uintptr
	LParam              uintptr
}
