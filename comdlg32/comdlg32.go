// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package comdlg32

import "golang.org/x/sys/windows"

var (
	// Library
	libcomdlg32 *windows.LazyDLL

	// Functions
	chooseColor          *windows.LazyProc
	commDlgExtendedError *windows.LazyProc
	getOpenFileName      *windows.LazyProc
	getSaveFileName      *windows.LazyProc
	printDlgEx           *windows.LazyProc
)

func init() {
	// Library
	libcomdlg32 = windows.NewLazySystemDLL("comdlg32.dll")

	// Functions
	chooseColor = libcomdlg32.NewProc("ChooseColorW")
	commDlgExtendedError = libcomdlg32.NewProc("CommDlgExtendedError")
	getOpenFileName = libcomdlg32.NewProc("GetOpenFileNameW")
	getSaveFileName = libcomdlg32.NewProc("GetSaveFileNameW")
	printDlgEx = libcomdlg32.NewProc("PrintDlgExW")
}
