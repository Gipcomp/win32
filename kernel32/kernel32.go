// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package kernel32

import "golang.org/x/sys/windows"

var (
	// Library
	libkernel32 *windows.LazyDLL

	// Functions
	activateActCtx                     *windows.LazyProc
	closeHandle                        *windows.LazyProc
	createActCtx                       *windows.LazyProc
	fileTimeToSystemTime               *windows.LazyProc
	findResource                       *windows.LazyProc
	getConsoleTitle                    *windows.LazyProc
	getConsoleWindow                   *windows.LazyProc
	getCurrentThreadId                 *windows.LazyProc
	getLastError                       *windows.LazyProc
	getLocaleInfo                      *windows.LazyProc
	getLogicalDriveStrings             *windows.LazyProc
	getModuleHandle                    *windows.LazyProc
	getModuleHandleExW                 *windows.LazyProc
	getNumberFormat                    *windows.LazyProc
	getPhysicallyInstalledSystemMemory *windows.LazyProc
	getProfileString                   *windows.LazyProc
	getThreadLocale                    *windows.LazyProc
	getThreadUILanguage                *windows.LazyProc
	getVersion                         *windows.LazyProc
	globalAlloc                        *windows.LazyProc
	globalFree                         *windows.LazyProc
	globalLock                         *windows.LazyProc
	globalUnlock                       *windows.LazyProc
	moveMemory                         *windows.LazyProc
	mulDiv                             *windows.LazyProc
	loadResource                       *windows.LazyProc
	lockResource                       *windows.LazyProc
	setLastError                       *windows.LazyProc
	sizeofResource                     *windows.LazyProc
	systemTimeToFileTime               *windows.LazyProc
)

func init() {
	// Library
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	// Functions
	activateActCtx = libkernel32.NewProc("ActivateActCtx")
	closeHandle = libkernel32.NewProc("CloseHandle")
	createActCtx = libkernel32.NewProc("CreateActCtxW")
	fileTimeToSystemTime = libkernel32.NewProc("FileTimeToSystemTime")
	findResource = libkernel32.NewProc("FindResourceW")
	getConsoleTitle = libkernel32.NewProc("GetConsoleTitleW")
	getConsoleWindow = libkernel32.NewProc("GetConsoleWindow")
	getCurrentThreadId = libkernel32.NewProc("GetCurrentThreadId")
	getLastError = libkernel32.NewProc("GetLastError")
	getLocaleInfo = libkernel32.NewProc("GetLocaleInfoW")
	getLogicalDriveStrings = libkernel32.NewProc("GetLogicalDriveStringsW")
	getModuleHandle = libkernel32.NewProc("GetModuleHandleW")
	getModuleHandleExW = libkernel32.NewProc("GetModuleHandleExW")
	getNumberFormat = libkernel32.NewProc("GetNumberFormatW")
	getPhysicallyInstalledSystemMemory = libkernel32.NewProc("GetPhysicallyInstalledSystemMemory")
	getProfileString = libkernel32.NewProc("GetProfileStringW")
	getThreadLocale = libkernel32.NewProc("GetThreadLocale")
	getThreadUILanguage = libkernel32.NewProc("GetThreadUILanguage")
	getVersion = libkernel32.NewProc("GetVersion")
	globalAlloc = libkernel32.NewProc("GlobalAlloc")
	globalFree = libkernel32.NewProc("GlobalFree")
	globalLock = libkernel32.NewProc("GlobalLock")
	globalUnlock = libkernel32.NewProc("GlobalUnlock")
	moveMemory = libkernel32.NewProc("RtlMoveMemory")
	mulDiv = libkernel32.NewProc("MulDiv")
	loadResource = libkernel32.NewProc("LoadResource")
	lockResource = libkernel32.NewProc("LockResource")
	setLastError = libkernel32.NewProc("SetLastError")
	sizeofResource = libkernel32.NewProc("SizeofResource")
	systemTimeToFileTime = libkernel32.NewProc("SystemTimeToFileTime")
}
