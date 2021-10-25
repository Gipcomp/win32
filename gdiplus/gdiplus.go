// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package gdiplus

import "golang.org/x/sys/windows"

var (
	// Library
	libgdiplus *windows.LazyDLL

	// Functions
	gdipCreateBitmapFromFile    *windows.LazyProc
	gdipCreateBitmapFromHBITMAP *windows.LazyProc
	gdipCreateHBITMAPFromBitmap *windows.LazyProc
	gdipDisposeImage            *windows.LazyProc
	gdiplusShutdown             *windows.LazyProc
	gdiplusStartup              *windows.LazyProc
)

var (
	token uintptr
)

func init() {
	// Library
	libgdiplus = windows.NewLazySystemDLL("gdiplus.dll")

	// Functions
	gdipCreateBitmapFromFile = libgdiplus.NewProc("GdipCreateBitmapFromFile")
	gdipCreateBitmapFromHBITMAP = libgdiplus.NewProc("GdipCreateBitmapFromHBITMAP")
	gdipCreateHBITMAPFromBitmap = libgdiplus.NewProc("GdipCreateHBITMAPFromBitmap")
	gdipDisposeImage = libgdiplus.NewProc("GdipDisposeImage")
	gdiplusShutdown = libgdiplus.NewProc("GdiplusShutdown")
	gdiplusStartup = libgdiplus.NewProc("GdiplusStartup")
}
