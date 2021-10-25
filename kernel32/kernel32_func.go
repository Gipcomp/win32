// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package kernel32

import (
	"syscall"
	"unsafe"

	"github.com/Gipcomp/win32/handle"
)

func ActivateActCtx(ctx handle.HANDLE) (uintptr, bool) {
	var cookie uintptr
	ret, _, _ := syscall.Syscall(activateActCtx.Addr(), 2,
		uintptr(ctx),
		uintptr(unsafe.Pointer(&cookie)),
		0)
	return cookie, ret != 0
}

func CloseHandle(hObject handle.HANDLE) bool {
	ret, _, _ := syscall.Syscall(closeHandle.Addr(), 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}

func CreateActCtx(ctx *ACTCTX) handle.HANDLE {
	if ctx != nil {
		ctx.size = uint32(unsafe.Sizeof(*ctx))
	}
	ret, _, _ := syscall.Syscall(
		createActCtx.Addr(),
		1,
		uintptr(unsafe.Pointer(ctx)),
		0,
		0)
	return handle.HANDLE(ret)
}

func FileTimeToSystemTime(lpFileTime *FILETIME, lpSystemTime *SYSTEMTIME) bool {
	ret, _, _ := syscall.Syscall(fileTimeToSystemTime.Addr(), 2,
		uintptr(unsafe.Pointer(lpFileTime)),
		uintptr(unsafe.Pointer(lpSystemTime)),
		0)

	return ret != 0
}

func FindResource(hModule HMODULE, lpName, lpType *uint16) HRSRC {
	ret, _, _ := syscall.Syscall(findResource.Addr(), 3,
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(lpType)))

	return HRSRC(ret)
}

func GetConsoleTitle(lpConsoleTitle *uint16, nSize uint32) uint32 {
	ret, _, _ := syscall.Syscall(getConsoleTitle.Addr(), 2,
		uintptr(unsafe.Pointer(lpConsoleTitle)),
		uintptr(nSize),
		0)

	return uint32(ret)
}

func GetConsoleWindow() handle.HWND {
	ret, _, _ := syscall.Syscall(getConsoleWindow.Addr(), 0,
		0,
		0,
		0)

	return handle.HWND(ret)
}

func GetCurrentThreadId() uint32 {
	ret, _, _ := syscall.Syscall(getCurrentThreadId.Addr(), 0,
		0,
		0,
		0)

	return uint32(ret)
}

func GetLastError() uint32 {
	ret, _, _ := syscall.Syscall(getLastError.Addr(), 0,
		0,
		0,
		0)

	return uint32(ret)
}

func GetLocaleInfo(Locale LCID, LCType LCTYPE, lpLCData *uint16, cchData int32) int32 {
	ret, _, _ := syscall.Syscall6(getLocaleInfo.Addr(), 4,
		uintptr(Locale),
		uintptr(LCType),
		uintptr(unsafe.Pointer(lpLCData)),
		uintptr(cchData),
		0,
		0)

	return int32(ret)
}

func GetLogicalDriveStrings(nBufferLength uint32, lpBuffer *uint16) uint32 {
	ret, _, _ := syscall.Syscall(getLogicalDriveStrings.Addr(), 2,
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)

	return uint32(ret)
}

func GetModuleHandle(lpModuleName *uint16) HINSTANCE {
	ret, _, _ := syscall.Syscall(getModuleHandle.Addr(), 1,
		uintptr(unsafe.Pointer(lpModuleName)),
		0,
		0)

	return HINSTANCE(ret)
}

func GetModuleHandleEx(flags uint32, moduleName *uint16, module *handle.HWND) (err error) {
	r1, _, e1 := syscall.Syscall(getModuleHandleExW.Addr(), 3, uintptr(flags), uintptr(unsafe.Pointer(moduleName)), uintptr(unsafe.Pointer(module)))
	if r1 == 0 {
		err = e1
	}
	return
}

func GetNumberFormat(Locale LCID, dwFlags uint32, lpValue *uint16, lpFormat *NUMBERFMT, lpNumberStr *uint16, cchNumber int32) int32 {
	ret, _, _ := syscall.Syscall6(getNumberFormat.Addr(), 6,
		uintptr(Locale),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpValue)),
		uintptr(unsafe.Pointer(lpFormat)),
		uintptr(unsafe.Pointer(lpNumberStr)),
		uintptr(cchNumber))

	return int32(ret)
}

func GetPhysicallyInstalledSystemMemory(totalMemoryInKilobytes *uint64) bool {
	if getPhysicallyInstalledSystemMemory.Find() != nil {
		return false
	}
	ret, _, _ := syscall.Syscall(getPhysicallyInstalledSystemMemory.Addr(), 1,
		uintptr(unsafe.Pointer(totalMemoryInKilobytes)),
		0,
		0)

	return ret != 0
}

func GetProfileString(lpAppName, lpKeyName, lpDefault *uint16, lpReturnedString uintptr, nSize uint32) bool {
	ret, _, _ := syscall.Syscall6(getProfileString.Addr(), 5,
		uintptr(unsafe.Pointer(lpAppName)),
		uintptr(unsafe.Pointer(lpKeyName)),
		uintptr(unsafe.Pointer(lpDefault)),
		lpReturnedString,
		uintptr(nSize),
		0)
	return ret != 0
}

func GetThreadLocale() LCID {
	ret, _, _ := syscall.Syscall(getThreadLocale.Addr(), 0,
		0,
		0,
		0)

	return LCID(ret)
}

func GetThreadUILanguage() LANGID {
	if getThreadUILanguage.Find() != nil {
		return 0
	}

	ret, _, _ := syscall.Syscall(getThreadUILanguage.Addr(), 0,
		0,
		0,
		0)

	return LANGID(ret)
}

func GetVersion() uint32 {
	ret, _, _ := syscall.Syscall(getVersion.Addr(), 0,
		0,
		0,
		0)
	return uint32(ret)
}

func GlobalAlloc(uFlags uint32, dwBytes uintptr) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalAlloc.Addr(), 2,
		uintptr(uFlags),
		dwBytes,
		0)

	return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) HGLOBAL {
	ret, _, _ := syscall.Syscall(globalFree.Addr(), 1,
		uintptr(hMem),
		0,
		0)

	return HGLOBAL(ret)
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(globalLock.Addr(), 1,
		uintptr(hMem),
		0,
		0)

	return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) bool {
	ret, _, _ := syscall.Syscall(globalUnlock.Addr(), 1,
		uintptr(hMem),
		0,
		0)

	return ret != 0
}

func MoveMemory(destination, source unsafe.Pointer, length uintptr) {
	syscall.Syscall(moveMemory.Addr(), 3,
		uintptr(unsafe.Pointer(destination)),
		uintptr(source),
		uintptr(length))
}

func MulDiv(nNumber, nNumerator, nDenominator int32) int32 {
	ret, _, _ := syscall.Syscall(mulDiv.Addr(), 3,
		uintptr(nNumber),
		uintptr(nNumerator),
		uintptr(nDenominator))

	return int32(ret)
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) HGLOBAL {
	ret, _, _ := syscall.Syscall(loadResource.Addr(), 2,
		uintptr(hModule),
		uintptr(hResInfo),
		0)

	return HGLOBAL(ret)
}

func LockResource(hResData HGLOBAL) uintptr {
	ret, _, _ := syscall.Syscall(lockResource.Addr(), 1,
		uintptr(hResData),
		0,
		0)

	return ret
}

func SetLastError(dwErrorCode uint32) {
	syscall.Syscall(setLastError.Addr(), 1,
		uintptr(dwErrorCode),
		0,
		0)
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC) uint32 {
	ret, _, _ := syscall.Syscall(sizeofResource.Addr(), 2,
		uintptr(hModule),
		uintptr(hResInfo),
		0)

	return uint32(ret)
}

func SystemTimeToFileTime(lpSystemTime *SYSTEMTIME, lpFileTime *FILETIME) bool {
	ret, _, _ := syscall.Syscall(systemTimeToFileTime.Addr(), 2,
		uintptr(unsafe.Pointer(lpSystemTime)),
		uintptr(unsafe.Pointer(lpFileTime)),
		0)

	return ret != 0
}
