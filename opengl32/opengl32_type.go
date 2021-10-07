// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package opengl32

import (
	"github.com/Gipcomp/win32/gdi32"
	"github.com/Gipcomp/win32/handle"
)

type (
	HGLRC handle.HANDLE
)

type LAYERPLANEDESCRIPTOR struct {
	NSize           uint16
	NVersion        uint16
	DwFlags         uint32
	IPixelType      uint8
	CColorBits      uint8
	CRedBits        uint8
	CRedShift       uint8
	CGreenBits      uint8
	CGreenShift     uint8
	CBlueBits       uint8
	CBlueShift      uint8
	CAlphaBits      uint8
	CAlphaShift     uint8
	CAccumBits      uint8
	CAccumRedBits   uint8
	CAccumGreenBits uint8
	CAccumBlueBits  uint8
	CAccumAlphaBits uint8
	CDepthBits      uint8
	CStencilBits    uint8
	CAuxBuffers     uint8
	ILayerType      uint8
	BReserved       uint8
	CrTransparent   gdi32.COLORREF
}

type POINTFLOAT struct {
	X, Y float32
}

type GLYPHMETRICSFLOAT struct {
	GmfBlackBoxX     float32
	GmfBlackBoxY     float32
	GmfptGlyphOrigin POINTFLOAT
	GmfCellIncX      float32
	GmfCellIncY      float32
}
