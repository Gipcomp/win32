// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package uxtheme

import "golang.org/x/sys/windows"

var (
	// Library
	libuxtheme *windows.LazyDLL

	// Functions
	closeThemeData      *windows.LazyProc
	drawThemeBackground *windows.LazyProc
	drawThemeTextEx     *windows.LazyProc
	getThemeColor       *windows.LazyProc
	getThemePartSize    *windows.LazyProc
	getThemeTextExtent  *windows.LazyProc
	isAppThemed         *windows.LazyProc
	openThemeData       *windows.LazyProc
	setWindowTheme      *windows.LazyProc
)

func init() {
	// Library
	libuxtheme = windows.NewLazySystemDLL("uxtheme.dll")

	// Functions
	closeThemeData = libuxtheme.NewProc("CloseThemeData")
	drawThemeBackground = libuxtheme.NewProc("DrawThemeBackground")
	drawThemeTextEx = libuxtheme.NewProc("DrawThemeTextEx")
	getThemeColor = libuxtheme.NewProc("GetThemeColor")
	getThemePartSize = libuxtheme.NewProc("GetThemePartSize")
	getThemeTextExtent = libuxtheme.NewProc("GetThemeTextExtent")
	isAppThemed = libuxtheme.NewProc("IsAppThemed")
	openThemeData = libuxtheme.NewProc("OpenThemeData")
	setWindowTheme = libuxtheme.NewProc("SetWindowTheme")
}
