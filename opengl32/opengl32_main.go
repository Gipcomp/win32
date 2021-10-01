// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package opengl32

import "golang.org/x/sys/windows"

var (
	// Library
	lib *windows.LazyDLL

	// Functions
	wglCopyContext            *windows.LazyProc
	wglCreateContext          *windows.LazyProc
	wglCreateLayerContext     *windows.LazyProc
	wglDeleteContext          *windows.LazyProc
	wglDescribeLayerPlane     *windows.LazyProc
	wglGetCurrentContext      *windows.LazyProc
	wglGetCurrentDC           *windows.LazyProc
	wglGetLayerPaletteEntries *windows.LazyProc
	wglGetProcAddress         *windows.LazyProc
	wglMakeCurrent            *windows.LazyProc
	wglRealizeLayerPalette    *windows.LazyProc
	wglSetLayerPaletteEntries *windows.LazyProc
	wglShareLists             *windows.LazyProc
	wglSwapLayerBuffers       *windows.LazyProc
	wglUseFontBitmaps         *windows.LazyProc
	wglUseFontOutlines        *windows.LazyProc
)

func init() {
	// Library
	lib = windows.NewLazySystemDLL("opengl32.dll")

	// Functions
	wglCopyContext = lib.NewProc("wglCopyContext")
	wglCreateContext = lib.NewProc("wglCreateContext")
	wglCreateLayerContext = lib.NewProc("wglCreateLayerContext")
	wglDeleteContext = lib.NewProc("wglDeleteContext")
	wglDescribeLayerPlane = lib.NewProc("wglDescribeLayerPlane")
	wglGetCurrentContext = lib.NewProc("wglGetCurrentContext")
	wglGetCurrentDC = lib.NewProc("wglGetCurrentDC")
	wglGetLayerPaletteEntries = lib.NewProc("wglGetLayerPaletteEntries")
	wglGetProcAddress = lib.NewProc("wglGetProcAddress")
	wglMakeCurrent = lib.NewProc("wglMakeCurrent")
	wglRealizeLayerPalette = lib.NewProc("wglRealizeLayerPalette")
	wglSetLayerPaletteEntries = lib.NewProc("wglSetLayerPaletteEntries")
	wglShareLists = lib.NewProc("wglShareLists")
	wglSwapLayerBuffers = lib.NewProc("wglSwapLayerBuffers")
	wglUseFontBitmaps = lib.NewProc("wglUseFontBitmapsW")
	wglUseFontOutlines = lib.NewProc("wglUseFontOutlinesW")
}
