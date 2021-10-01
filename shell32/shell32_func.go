// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shell32

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/handle"
	"github.com/D4v1dW3bb/winapi/kernel32"
	"github.com/D4v1dW3bb/winapi/user32"
	"github.com/D4v1dW3bb/winapi/win"
)

func DragAcceptFiles(hWnd handle.HWND, fAccept bool) bool {
	ret, _, _ := syscall.Syscall(dragAcceptFiles.Addr(), 2,
		uintptr(hWnd),
		uintptr(win.BoolToBOOL(fAccept)),
		0)

	return ret != 0
}

func DragQueryFile(hDrop HDROP, iFile uint, lpszFile *uint16, cch uint) uint {
	ret, _, _ := syscall.Syscall6(dragQueryFile.Addr(), 4,
		uintptr(hDrop),
		uintptr(iFile),
		uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch),
		0,
		0)

	return uint(ret)
}

func DragFinish(hDrop HDROP) {
	syscall.Syscall(dragAcceptFiles.Addr(), 1,
		uintptr(hDrop),
		0,
		0)
}

func ExtractIcon(hInst kernel32.HINSTANCE, exeFileName *uint16, iconIndex int32) user32.HICON {
	ret, _, _ := syscall.Syscall(extractIcon.Addr(), 3,
		uintptr(hInst),
		uintptr(unsafe.Pointer(exeFileName)),
		uintptr(iconIndex))

	return user32.HICON(ret)
}

func SHBrowseForFolder(lpbi *BROWSEINFO) uintptr {
	ret, _, _ := syscall.Syscall(shBrowseForFolder.Addr(), 1,
		uintptr(unsafe.Pointer(lpbi)),
		0,
		0)

	return ret
}

func SHDefExtractIcon(pszIconFile *uint16, iIndex int32, uFlags uint32, phiconLarge, phiconSmall *user32.HICON, nIconSize uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall6(shDefExtractIcon.Addr(), 6,
		uintptr(unsafe.Pointer(pszIconFile)),
		uintptr(iIndex),
		uintptr(uFlags),
		uintptr(unsafe.Pointer(phiconLarge)),
		uintptr(unsafe.Pointer(phiconSmall)),
		uintptr(nIconSize))

	return win.HRESULT(ret)
}

func SHGetFileInfo(pszPath *uint16, dwFileAttributes uint32, psfi *SHFILEINFO, cbFileInfo, uFlags uint32) uintptr {
	ret, _, _ := syscall.Syscall6(shGetFileInfo.Addr(), 5,
		uintptr(unsafe.Pointer(pszPath)),
		uintptr(dwFileAttributes),
		uintptr(unsafe.Pointer(psfi)),
		uintptr(cbFileInfo),
		uintptr(uFlags),
		0)

	return ret
}

func SHGetPathFromIDList(pidl uintptr, pszPath *uint16) bool {
	ret, _, _ := syscall.Syscall(shGetPathFromIDList.Addr(), 2,
		pidl,
		uintptr(unsafe.Pointer(pszPath)),
		0)

	return ret != 0
}

func SHGetSpecialFolderPath(hwndOwner handle.HWND, lpszPath *uint16, csidl CSIDL, fCreate bool) bool {
	ret, _, _ := syscall.Syscall6(shGetSpecialFolderPath.Addr(), 4,
		uintptr(hwndOwner),
		uintptr(unsafe.Pointer(lpszPath)),
		uintptr(csidl),
		uintptr(win.BoolToBOOL(fCreate)),
		0,
		0)

	return ret != 0
}

func SHParseDisplayName(pszName *uint16, pbc uintptr, ppidl *uintptr, sfgaoIn uint32, psfgaoOut *uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall6(shParseDisplayName.Addr(), 5,
		uintptr(unsafe.Pointer(pszName)),
		pbc,
		uintptr(unsafe.Pointer(ppidl)),
		0,
		uintptr(unsafe.Pointer(psfgaoOut)),
		0)

	return win.HRESULT(ret)
}

func SHGetStockIconInfo(stockIconId int32, uFlags uint32, stockIcon *SHSTOCKICONINFO) win.HRESULT {
	if shGetStockIconInfo.Find() != nil {
		return win.HRESULT(0)
	}
	ret, _, _ := syscall.Syscall6(shGetStockIconInfo.Addr(), 3,
		uintptr(stockIconId),
		uintptr(uFlags),
		uintptr(unsafe.Pointer(stockIcon)),
		0,
		0,
		0,
	)
	return win.HRESULT(ret)
}

func ShellExecute(hWnd handle.HWND, verb *uint16, file *uint16, args *uint16, cwd *uint16, showCmd int) bool {
	ret, _, _ := syscall.Syscall6(shellExecute.Addr(), 6,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(verb)),
		uintptr(unsafe.Pointer(file)),
		uintptr(unsafe.Pointer(args)),
		uintptr(unsafe.Pointer(cwd)),
		uintptr(showCmd),
	)
	return ret != 0
}

func Shell_NotifyIcon(dwMessage uint32, lpdata *NOTIFYICONDATA) bool {
	ret, _, _ := syscall.Syscall(shell_NotifyIcon.Addr(), 2,
		uintptr(dwMessage),
		uintptr(unsafe.Pointer(lpdata)),
		0)

	return ret != 0
}
