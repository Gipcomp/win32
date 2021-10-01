// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shell32

import "golang.org/x/sys/windows"

var (
	// Library
	libshell32 *windows.LazyDLL

	// Functions
	dragAcceptFiles        *windows.LazyProc
	dragFinish             *windows.LazyProc
	dragQueryFile          *windows.LazyProc
	extractIcon            *windows.LazyProc
	shBrowseForFolder      *windows.LazyProc
	shDefExtractIcon       *windows.LazyProc
	shGetFileInfo          *windows.LazyProc
	shGetPathFromIDList    *windows.LazyProc
	shGetSpecialFolderPath *windows.LazyProc
	shParseDisplayName     *windows.LazyProc
	shGetStockIconInfo     *windows.LazyProc
	shellExecute           *windows.LazyProc
	shell_NotifyIcon       *windows.LazyProc
)

func init() {
	// Library
	libshell32 = windows.NewLazySystemDLL("shell32.dll")

	// Functions
	dragAcceptFiles = libshell32.NewProc("DragAcceptFiles")
	dragFinish = libshell32.NewProc("DragFinish")
	dragQueryFile = libshell32.NewProc("DragQueryFileW")
	extractIcon = libshell32.NewProc("ExtractIconW")
	shBrowseForFolder = libshell32.NewProc("SHBrowseForFolderW")
	shDefExtractIcon = libshell32.NewProc("SHDefExtractIconW")
	shGetFileInfo = libshell32.NewProc("SHGetFileInfoW")
	shGetPathFromIDList = libshell32.NewProc("SHGetPathFromIDListW")
	shGetSpecialFolderPath = libshell32.NewProc("SHGetSpecialFolderPathW")
	shGetStockIconInfo = libshell32.NewProc("SHGetStockIconInfo")
	shellExecute = libshell32.NewProc("ShellExecuteW")
	shell_NotifyIcon = libshell32.NewProc("Shell_NotifyIconW")
	shParseDisplayName = libshell32.NewProc("SHParseDisplayName")
}
