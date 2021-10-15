// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package user32

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/Gipcomp/win32/gdi32"
	"github.com/Gipcomp/win32/handle"
	"github.com/Gipcomp/win32/kernel32"
	"github.com/Gipcomp/win32/win"
	"github.com/Gipcomp/win32/winuser"
	"golang.org/x/sys/windows"
)

// Functions from the win.go file of lxn/win

func GET_X_LPARAM(lp uintptr) int32 {
	return int32(int16(win.LOWORD(uint32(lp))))
}

func GET_Y_LPARAM(lp uintptr) int32 {
	return int32(int16(win.HIWORD(uint32(lp))))
}

func AddClipboardFormatListener(hwnd handle.HWND) bool {
	if addClipboardFormatListener.Find() != nil {
		return false
	}

	ret, _, _ := syscall.Syscall(addClipboardFormatListener.Addr(), 1,
		uintptr(hwnd),
		0,
		0)

	return ret != 0
}

func AdjustWindowRect(lpRect *gdi32.RECT, dwStyle uint32, bMenu bool) bool {
	ret, _, _ := syscall.Syscall(adjustWindowRect.Addr(), 3,
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(dwStyle),
		uintptr(win.BoolToBOOL(bMenu)))

	return ret != 0
}

func AttachThreadInput(idAttach int32, idAttachTo int32, fAttach bool) bool {
	ret, _, _ := syscall.Syscall(attachThreadInput.Addr(), 3,
		uintptr(idAttach),
		uintptr(idAttachTo),
		uintptr(win.BoolToBOOL(fAttach)))

	return ret != 0
}

func AnimateWindow(hwnd handle.HWND, dwTime, dwFlags uint32) bool {
	ret, _, _ := syscall.Syscall(animateWindow.Addr(), 3,
		uintptr(hwnd),
		uintptr(dwTime),
		uintptr(dwFlags))

	return ret != 0
}

func BeginDeferWindowPos(nNumWindows int32) HDWP {
	ret, _, _ := syscall.Syscall(beginDeferWindowPos.Addr(), 1,
		uintptr(nNumWindows),
		0,
		0)

	return HDWP(ret)
}

func GetWindowThreadProcessId(hwnd handle.HWND, processId *uint32) uint32 {
	ret, _, _ := syscall.Syscall(getWindowThreadProcessId.Addr(), 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(processId)),
		0)

	return uint32(ret)
}

func BeginPaint(hwnd handle.HWND, lpPaint *PAINTSTRUCT) gdi32.HDC {
	ret, _, _ := syscall.Syscall(beginPaint.Addr(), 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpPaint)),
		0)

	return gdi32.HDC(ret)
}

func BringWindowToTop(hwnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(bringWindowToTop.Addr(), 1,
		uintptr(hwnd),
		0,
		0)
	return ret != 0
}

func CallWindowProc(lpPrevWndFunc uintptr, hWnd handle.HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(callWindowProc.Addr(), 5,
		lpPrevWndFunc,
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam,
		0)

	return ret
}

func ChangeWindowMessageFilterEx(hwnd handle.HWND, msg uint32, action uint32, changeFilterStruct *CHANGEFILTERSTRUCT) bool {
	ret, _, _ := syscall.Syscall6(changeWindowMessageFilterEx.Addr(), 4,
		uintptr(hwnd),
		uintptr(msg),
		uintptr(action),
		uintptr(unsafe.Pointer(changeFilterStruct)),
		0,
		0)
	return ret != 0
}

func CheckMenuRadioItem(hmenu winuser.HMENU, first, last, check, flags uint32) bool {
	ret, _, _ := syscall.Syscall6(checkMenuRadioItem.Addr(), 5,
		uintptr(hmenu),
		uintptr(first),
		uintptr(last),
		uintptr(check),
		uintptr(flags),
		0)

	return ret != 0
}

func ClientToScreen(hwnd handle.HWND, lpPoint *gdi32.POINT) bool {
	ret, _, _ := syscall.Syscall(clientToScreen.Addr(), 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpPoint)),
		0)

	return ret != 0
}

func CloseClipboard() bool {
	ret, _, _ := syscall.Syscall(closeClipboard.Addr(), 0,
		0,
		0,
		0)

	return ret != 0
}

func CreateDialogParam(instRes kernel32.HINSTANCE, name *uint16, parent handle.HWND,
	proc, param uintptr) handle.HWND {
	ret, _, _ := syscall.Syscall6(createDialogParam.Addr(), 5,
		uintptr(instRes),
		uintptr(unsafe.Pointer(name)),
		uintptr(parent),
		proc,
		param,
		0)

	return handle.HWND(ret)
}

func CreateIconIndirect(lpiconinfo *ICONINFO) HICON {
	ret, _, _ := syscall.Syscall(createIconIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lpiconinfo)),
		0,
		0)

	return HICON(ret)
}

func CreateMenu() winuser.HMENU {
	ret, _, _ := syscall.Syscall(createMenu.Addr(), 0,
		0,
		0,
		0)

	return winuser.HMENU(ret)
}

func CreatePopupMenu() winuser.HMENU {
	ret, _, _ := syscall.Syscall(createPopupMenu.Addr(), 0,
		0,
		0,
		0)

	return winuser.HMENU(ret)
}

func CreateWindowEx(dwExStyle uint32, lpClassName, lpWindowName *uint16, dwStyle uint32, x, y, nWidth, nHeight int32, hWndParent handle.HWND, hMenu winuser.HMENU, hInstance kernel32.HINSTANCE, lpParam unsafe.Pointer) handle.HWND {
	ret, _, _ := syscall.Syscall12(createWindowEx.Addr(), 12,
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam))

	return handle.HWND(ret)
}

func DeferWindowPos(hWinPosInfo HDWP, hWnd, hWndInsertAfter handle.HWND, x, y, cx, cy int32, uFlags uint32) HDWP {
	ret, _, _ := syscall.Syscall9(deferWindowPos.Addr(), 8,
		uintptr(hWinPosInfo),
		uintptr(hWnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags),
		0)

	return HDWP(ret)
}

func DefWindowProc(hWnd handle.HWND, Msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(defWindowProc.Addr(), 4,
		uintptr(hWnd),
		uintptr(Msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func DeleteMenu(hMenu winuser.HMENU, uPosition uint32, uFlags uint32) bool {
	ret, _, _ := syscall.Syscall(deleteMenu.Addr(), 3,
		uintptr(hMenu),
		uintptr(uPosition),
		uintptr(uFlags))

	return ret != 0
}

func DestroyIcon(hIcon HICON) bool {
	ret, _, _ := syscall.Syscall(destroyIcon.Addr(), 1,
		uintptr(hIcon),
		0,
		0)

	return ret != 0
}

func DestroyMenu(hMenu winuser.HMENU) bool {
	ret, _, _ := syscall.Syscall(destroyMenu.Addr(), 1,
		uintptr(hMenu),
		0,
		0)

	return ret != 0
}

func DestroyWindow(hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(destroyWindow.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func DialogBoxParam(instRes kernel32.HINSTANCE, name *uint16, parent handle.HWND, proc, param uintptr) int {
	ret, _, _ := syscall.Syscall6(dialogBoxParam.Addr(), 5,
		uintptr(instRes),
		uintptr(unsafe.Pointer(name)),
		uintptr(parent),
		proc,
		param,
		0)

	return int(ret)
}

func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := syscall.Syscall(dispatchMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret
}

func DrawFocusRect(hDC gdi32.HDC, lprc *gdi32.RECT) bool {
	ret, _, _ := syscall.Syscall(drawFocusRect.Addr(), 2,
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		0)

	return ret != 0
}

func DrawIconEx(hdc gdi32.HDC, xLeft, yTop int32, hIcon HICON, cxWidth, cyWidth int32, istepIfAniCur uint32, hbrFlickerFreeDraw gdi32.HBRUSH, diFlags uint32) bool {
	ret, _, _ := syscall.Syscall9(drawIconEx.Addr(), 9,
		uintptr(hdc),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(hIcon),
		uintptr(cxWidth),
		uintptr(cyWidth),
		uintptr(istepIfAniCur),
		uintptr(hbrFlickerFreeDraw),
		uintptr(diFlags))

	return ret != 0
}

func DrawMenuBar(hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(drawMenuBar.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func DrawTextEx(hdc gdi32.HDC, lpchText *uint16, cchText int32, lprc *gdi32.RECT, dwDTFormat uint32, lpDTParams *DRAWTEXTPARAMS) int32 {
	ret, _, _ := syscall.Syscall6(drawTextEx.Addr(), 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpchText)),
		uintptr(cchText),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(dwDTFormat),
		uintptr(unsafe.Pointer(lpDTParams)))

	return int32(ret)
}

func EmptyClipboard() bool {
	ret, _, _ := syscall.Syscall(emptyClipboard.Addr(), 0,
		0,
		0,
		0)

	return ret != 0
}

func EnableMenuItem(hMenu winuser.HMENU, uIDEnableItem uint32, uEnable uint32) bool {
	ret, _, _ := syscall.Syscall(enableMenuItem.Addr(), 3,
		uintptr(hMenu),
		uintptr(uIDEnableItem),
		uintptr(uEnable))

	return ret != 0
}

func EnableWindow(hWnd handle.HWND, bEnable bool) bool {
	ret, _, _ := syscall.Syscall(enableWindow.Addr(), 2,
		uintptr(hWnd),
		uintptr(win.BoolToBOOL(bEnable)),
		0)

	return ret != 0
}

func EndDeferWindowPos(hWinPosInfo HDWP) bool {
	ret, _, _ := syscall.Syscall(endDeferWindowPos.Addr(), 1,
		uintptr(hWinPosInfo),
		0,
		0)

	return ret != 0
}

func EndDialog(hwnd handle.HWND, result int) bool {
	ret, _, _ := syscall.Syscall(endDialog.Addr(), 2,
		uintptr(hwnd),
		uintptr(result),
		0)

	return ret != 0
}

func EndPaint(hwnd handle.HWND, lpPaint *PAINTSTRUCT) bool {
	ret, _, _ := syscall.Syscall(endPaint.Addr(), 2,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpPaint)),
		0)

	return ret != 0
}

func EnumChildWindows(hWndParent handle.HWND, lpEnumFunc, lParam uintptr) bool {
	ret, _, _ := syscall.Syscall(enumChildWindows.Addr(), 3,
		uintptr(hWndParent),
		lpEnumFunc,
		lParam)

	return ret != 0
}

func FillRect(hDc gdi32.HDC, lprc *gdi32.RECT, hbr gdi32.HBRUSH) int {
	ret, _, _ := syscall.Syscall(fillRect.Addr(), 3,
		uintptr(hDc),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr))

	return int(ret)
}

func FindWindow(lpClassName, lpWindowName *uint16) handle.HWND {
	ret, _, _ := syscall.Syscall(findWindow.Addr(), 2,
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		0)

	return handle.HWND(ret)
}

func GetActiveWindow() handle.HWND {
	ret, _, _ := syscall.Syscall(getActiveWindow.Addr(), 0,
		0,
		0,
		0)

	return handle.HWND(ret)
}

func GetAncestor(hWnd handle.HWND, gaFlags uint32) handle.HWND {
	ret, _, _ := syscall.Syscall(getAncestor.Addr(), 2,
		uintptr(hWnd),
		uintptr(gaFlags),
		0)

	return handle.HWND(ret)
}

func GetCaretPos(lpPoint *gdi32.POINT) bool {
	ret, _, _ := syscall.Syscall(getCaretPos.Addr(), 1,
		uintptr(unsafe.Pointer(lpPoint)),
		0,
		0)

	return ret != 0
}

func GetClassName(hWnd handle.HWND, className *uint16, maxCount int) (int, error) {
	ret, _, e := syscall.Syscall(getClassName.Addr(), 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(className)),
		uintptr(maxCount))
	if ret == 0 {
		return 0, e
	}
	return int(ret), nil
}

func GetClientRect(hWnd handle.HWND, rect *gdi32.RECT) bool {
	ret, _, _ := syscall.Syscall(getClientRect.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)),
		0)

	return ret != 0
}

func GetClientRectWV2(hwnd handle.HWND) (*gdi32.RECT, error) {
	var rect gdi32.RECT
	_, _, err := getClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect)))
	if err != nil && !errors.Is(err, syscall.Errno(0)) {
		return nil, err
	}

	return &rect, nil
}

func GetClipboardData(uFormat uint32) handle.HANDLE {
	ret, _, _ := syscall.Syscall(getClipboardData.Addr(), 1,
		uintptr(uFormat),
		0,
		0)

	return handle.HANDLE(ret)
}

func GetCursorPos(lpPoint *gdi32.POINT) bool {
	ret, _, _ := syscall.Syscall(getCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(lpPoint)),
		0,
		0)

	return ret != 0
}

func GetDesktopWindow() handle.HWND {
	ret, _, _ := syscall.Syscall(getDesktopWindow.Addr(), 0,
		0,
		0,
		0)

	return handle.HWND(ret)
}

func GetDC(hWnd handle.HWND) gdi32.HDC {
	ret, _, _ := syscall.Syscall(getDC.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return gdi32.HDC(ret)
}

func GetDlgItem(hDlg handle.HWND, nIDDlgItem int32) handle.HWND {
	ret, _, _ := syscall.Syscall(getDlgItem.Addr(), 2,
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		0)

	return handle.HWND(ret)
}

func GetDpiForWindow(hwnd handle.HWND) uint32 {
	if getDpiForWindow.Find() != nil {
		hdc := GetDC(hwnd)
		defer ReleaseDC(hwnd, hdc)

		return uint32(gdi32.GetDeviceCaps(hdc, gdi32.LOGPIXELSY))
	}

	ret, _, _ := syscall.Syscall(getDpiForWindow.Addr(), 1,
		uintptr(hwnd),
		0,
		0)

	return uint32(ret)
}

func GetFocus() handle.HWND {
	ret, _, _ := syscall.Syscall(getFocus.Addr(), 0,
		0,
		0,
		0)

	return handle.HWND(ret)
}

func GetForegroundWindow() handle.HWND {
	ret, _, _ := syscall.Syscall(getForegroundWindow.Addr(), 0,
		0,
		0,
		0)

	return handle.HWND(ret)
}

func GetIconInfo(hicon HICON, piconinfo *ICONINFO) bool {
	ret, _, _ := syscall.Syscall(getIconInfo.Addr(), 2,
		uintptr(hicon),
		uintptr(unsafe.Pointer(piconinfo)),
		0)

	return ret != 0
}

func GetKeyState(nVirtKey int32) int16 {
	ret, _, _ := syscall.Syscall(getKeyState.Addr(), 1,
		uintptr(nVirtKey),
		0,
		0)

	return int16(ret)
}

func GetMenuCheckMarkDimensions() int32 {
	ret, _, _ := syscall.Syscall(getMenuCheckMarkDimensions.Addr(), 0,
		0,
		0,
		0)

	return int32(ret)
}

func GetMenuInfo(hmenu winuser.HMENU, lpcmi *winuser.MENUINFO) bool {
	ret, _, _ := syscall.Syscall(getMenuInfo.Addr(), 2,
		uintptr(hmenu),
		uintptr(unsafe.Pointer(lpcmi)),
		0)

	return ret != 0
}

func GetMenuItemCount(hMenu winuser.HMENU) int32 {
	ret, _, _ := syscall.Syscall(getMenuItemCount.Addr(), 1,
		uintptr(hMenu),
		0,
		0)

	return int32(ret)
}

func GetMenuItemID(hMenu winuser.HMENU, nPos int32) uint32 {
	ret, _, _ := syscall.Syscall(getMenuItemID.Addr(), 2,
		uintptr(hMenu),
		uintptr(nPos),
		0)

	return uint32(ret)
}

func GetMenuItemInfo(hmenu winuser.HMENU, item uint32, fByPosition win.BOOL, lpmii *winuser.MENUITEMINFO) bool {
	ret, _, _ := syscall.Syscall6(getMenuItemInfo.Addr(), 4,
		uintptr(hmenu),
		uintptr(item),
		uintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmii)),
		0,
		0)

	return ret != 0
}

func GetMessage(msg *MSG, hWnd handle.HWND, msgFilterMin, msgFilterMax uint32) win.BOOL {
	ret, _, _ := syscall.Syscall6(getMessage.Addr(), 4,
		uintptr(unsafe.Pointer(msg)),
		uintptr(hWnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
		0,
		0)

	return win.BOOL(ret)
}

func GetMessageWV2() (*MSG, error) {
	var msg MSG

	r, _, _ := getMessage.Call(
		uintptr(unsafe.Pointer(&msg)),
		0,
		0,
		0,
	)

	if int32(r) == -1 {
		return nil, windows.GetLastError()
	}

	return &msg, nil
}

func GetMonitorInfo(hMonitor HMONITOR, lpmi *MONITORINFO) bool {
	ret, _, _ := syscall.Syscall(getMonitorInfo.Addr(), 2,
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lpmi)),
		0)

	return ret != 0
}

func GetParent(hWnd handle.HWND) handle.HWND {
	ret, _, _ := syscall.Syscall(getParent.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return handle.HWND(ret)
}

func GetRawInputData(hRawInput HRAWINPUT, uiCommand uint32, pData unsafe.Pointer, pcbSize *uint32, cBSizeHeader uint32) uint32 {
	ret, _, _ := syscall.Syscall6(getRawInputData.Addr(), 5,
		uintptr(hRawInput),
		uintptr(uiCommand),
		uintptr(pData),
		uintptr(unsafe.Pointer(pcbSize)),
		uintptr(cBSizeHeader),
		0)

	return uint32(ret)
}

func GetScrollInfo(hwnd handle.HWND, fnBar int32, lpsi *SCROLLINFO) bool {
	ret, _, _ := syscall.Syscall(getScrollInfo.Addr(), 3,
		uintptr(hwnd),
		uintptr(fnBar),
		uintptr(unsafe.Pointer(lpsi)))

	return ret != 0
}

func GetSubMenu(hMenu winuser.HMENU, nPos int32) winuser.HMENU {
	ret, _, _ := syscall.Syscall(getSubMenu.Addr(), 2,
		uintptr(hMenu),
		uintptr(nPos),
		0)

	return winuser.HMENU(ret)
}

func GetSysColor(nIndex int) uint32 {
	ret, _, _ := syscall.Syscall(getSysColor.Addr(), 1,
		uintptr(nIndex),
		0,
		0)

	return uint32(ret)
}

func GetSysColorBrush(nIndex int) gdi32.HBRUSH {
	ret, _, _ := syscall.Syscall(getSysColorBrush.Addr(), 1,
		uintptr(nIndex),
		0,
		0)

	return gdi32.HBRUSH(ret)
}

func GetSystemMenu(hWnd handle.HWND, revert bool) winuser.HMENU {
	ret, _, _ := syscall.Syscall(getSystemMenu.Addr(), 2,
		uintptr(hWnd),
		uintptr(win.BoolToBOOL(revert)),
		0)
	return winuser.HMENU(ret)
}

func GetSystemMetrics(nIndex int32) int32 {
	ret, _, _ := syscall.Syscall(getSystemMetrics.Addr(), 1,
		uintptr(nIndex),
		0,
		0)

	return int32(ret)
}

func GetSystemMetricsForDpi(nIndex int32, dpi uint32) int32 {
	if getSystemMetricsForDpi.Find() != nil {
		return GetSystemMetrics(nIndex)
	}

	ret, _, _ := syscall.Syscall(getSystemMetricsForDpi.Addr(), 2,
		uintptr(nIndex),
		uintptr(dpi),
		0)

	return int32(ret)
}

func GetWindow(hWnd handle.HWND, uCmd uint32) handle.HWND {
	ret, _, _ := syscall.Syscall(getWindow.Addr(), 2,
		uintptr(hWnd),
		uintptr(uCmd),
		0)

	return handle.HWND(ret)
}

func GetWindowLong(hWnd handle.HWND, index int32) int32 {
	ret, _, _ := syscall.Syscall(getWindowLong.Addr(), 2,
		uintptr(hWnd),
		uintptr(index),
		0)

	return int32(ret)
}

func GetWindowLongPtr(hWnd handle.HWND, index int32) uintptr {
	ret, _, _ := syscall.Syscall(getWindowLongPtr.Addr(), 2,
		uintptr(hWnd),
		uintptr(index),
		0)

	return ret
}

func GetWindowPlacement(hWnd handle.HWND, lpwndpl *WINDOWPLACEMENT) bool {
	ret, _, _ := syscall.Syscall(getWindowPlacement.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpwndpl)),
		0)

	return ret != 0
}

func GetWindowRect(hWnd handle.HWND, rect *gdi32.RECT) bool {
	ret, _, _ := syscall.Syscall(getWindowRect.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(rect)),
		0)

	return ret != 0
}

func InsertMenuItem(hMenu winuser.HMENU, uItem uint32, fByPosition bool, lpmii *winuser.MENUITEMINFO) bool {
	ret, _, _ := syscall.Syscall6(insertMenuItem.Addr(), 4,
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(win.BoolToBOOL(fByPosition)),
		uintptr(unsafe.Pointer(lpmii)),
		0,
		0)

	return ret != 0
}

func InvalidateRect(hWnd handle.HWND, lpRect *gdi32.RECT, bErase bool) bool {
	ret, _, _ := syscall.Syscall(invalidateRect.Addr(), 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(win.BoolToBOOL(bErase)))

	return ret != 0
}

func IsChild(hWndParent, hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(isChild.Addr(), 2,
		uintptr(hWndParent),
		uintptr(hWnd),
		0)

	return ret != 0
}

func IsClipboardFormatAvailable(format uint32) bool {
	ret, _, _ := syscall.Syscall(isClipboardFormatAvailable.Addr(), 1,
		uintptr(format),
		0,
		0)

	return ret != 0
}

func IsDialogMessage(hWnd handle.HWND, msg *MSG) bool {
	ret, _, _ := syscall.Syscall(isDialogMessage.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(msg)),
		0)

	return ret != 0
}

func IsIconic(hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(isIconic.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func IsWindowEnabled(hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(isWindowEnabled.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func IsWindowVisible(hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(isWindowVisible.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func IsZoomed(hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(isZoomed.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func KillTimer(hWnd handle.HWND, uIDEvent uintptr) bool {
	ret, _, _ := syscall.Syscall(killTimer.Addr(), 2,
		uintptr(hWnd),
		uIDEvent,
		0)

	return ret != 0
}

func LoadCursor(hInstance kernel32.HINSTANCE, lpCursorName *uint16) HCURSOR {
	ret, _, _ := syscall.Syscall(loadCursor.Addr(), 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpCursorName)),
		0)

	return HCURSOR(ret)
}

func LoadIcon(hInstance kernel32.HINSTANCE, lpIconName *uint16) HICON {
	ret, _, _ := syscall.Syscall(loadIcon.Addr(), 2,
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
		0)

	return HICON(ret)
}

func LoadImage(hinst kernel32.HINSTANCE, lpszName *uint16, uType uint32, cxDesired, cyDesired int32, fuLoad uint32) handle.HANDLE {
	ret, _, _ := syscall.Syscall6(loadImage.Addr(), 6,
		uintptr(hinst),
		uintptr(unsafe.Pointer(lpszName)),
		uintptr(uType),
		uintptr(cxDesired),
		uintptr(cyDesired),
		uintptr(fuLoad))

	return handle.HANDLE(ret)
}

func LoadMenu(hinst kernel32.HINSTANCE, name *uint16) winuser.HMENU {
	ret, _, _ := syscall.Syscall(loadMenu.Addr(), 2,
		uintptr(hinst),
		uintptr(unsafe.Pointer(name)),
		0)

	return winuser.HMENU(ret)
}

func LoadString(instRes kernel32.HINSTANCE, id uint32, buf *uint16, length int32) int32 {
	ret, _, _ := syscall.Syscall6(loadString.Addr(), 4,
		uintptr(instRes),
		uintptr(id),
		uintptr(unsafe.Pointer(buf)),
		uintptr(length),
		0,
		0)

	return int32(ret)
}

// Plays a waveform sound. uType is the sound to be played. The sounds are set by the user through the Sound control panel application.
// The following values can be used as a sound:
//
//	MB_ICONASTERISK (see MB_ICONINFORMATION)
//	MB_ICONEXCLAMATION (see MB_ICONWARNING)
//	MB_ICONERROR (The sound specified as the Windows Critical Stop sound)
//	MB_ICONHAND (See MB_ICONERROR)
//	MB_ICONINFORMATION (The sounds specified as the Windows Asterisk sound)
//	MB_ICONQUESTION (The sound specified as the Windows Question sound)
// 	MB_ICONSTOP (See MB_ICONERROR)
//	MB_ICONWARNING (The sounds specified as the Windows Exclamation sound)
//	MB_OK (The sound specified as the Windows Default Beep sound)
//
// The function will return true if the function succeeds, false if otherwise.
func MessageBeep(uType uint32) bool {
	ret, _, _ := syscall.Syscall(messageBeep.Addr(), 2,
		uintptr(uType),
		0,
		0)

	return ret != 0
}

func MessageBox(hWnd handle.HWND, lpText, lpCaption *uint16, uType uint32) int32 {
	ret, _, _ := syscall.Syscall6(messageBox.Addr(), 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(uType),
		0,
		0)

	return int32(ret)
}

func MonitorFromWindow(hwnd handle.HWND, dwFlags uint32) HMONITOR {
	ret, _, _ := syscall.Syscall(monitorFromWindow.Addr(), 2,
		uintptr(hwnd),
		uintptr(dwFlags),
		0)

	return HMONITOR(ret)
}

func MoveWindow(hWnd handle.HWND, x, y, width, height int32, repaint bool) bool {
	ret, _, _ := syscall.Syscall6(moveWindow.Addr(), 6,
		uintptr(hWnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(win.BoolToBOOL(repaint)))

	return ret != 0
}

func NotifyWinEvent(event uint32, hwnd handle.HWND, idObject, idChild int32) {
	syscall.Syscall6(notifyWinEvent.Addr(), 4,
		uintptr(event),
		uintptr(hwnd),
		uintptr(idObject),
		uintptr(idChild),
		0,
		0)
}

func UnregisterClass(name *uint16) bool {
	ret, _, _ := syscall.Syscall(unregisterClass.Addr(), 1,
		uintptr(unsafe.Pointer(name)),
		0,
		0)

	return ret != 0
}

func OpenClipboard(hWndNewOwner handle.HWND) bool {
	ret, _, _ := syscall.Syscall(openClipboard.Addr(), 1,
		uintptr(hWndNewOwner),
		0,
		0)

	return ret != 0
}

func PeekMessage(lpMsg *MSG, hWnd handle.HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := syscall.Syscall6(peekMessage.Addr(), 5,
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
		0)

	return ret != 0
}

func PostMessage(hWnd handle.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(postMessage.Addr(), 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func PostQuitMessage(exitCode int32) {
	syscall.Syscall(postQuitMessage.Addr(), 1,
		uintptr(exitCode),
		0,
		0)
}

const (
	// RedrawWindow() flags
	RDW_INVALIDATE    = 0x0001
	RDW_INTERNALPAINT = 0x0002
	RDW_ERASE         = 0x0004

	RDW_VALIDATE        = 0x0008
	RDW_NOINTERNALPAINT = 0x0010
	RDW_NOERASE         = 0x0020

	RDW_NOCHILDREN  = 0x0040
	RDW_ALLCHILDREN = 0x0080

	RDW_UPDATENOW = 0x0100
	RDW_ERASENOW  = 0x0200

	RDW_FRAME   = 0x0400
	RDW_NOFRAME = 0x0800
)

func RedrawWindow(hWnd handle.HWND, lprcUpdate *gdi32.RECT, hrgnUpdate gdi32.HRGN, flags uint32) bool {
	ret, _, _ := syscall.Syscall6(redrawWindow.Addr(), 4,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lprcUpdate)),
		uintptr(hrgnUpdate),
		uintptr(flags),
		0,
		0)

	return ret != 0
}

func RegisterClassEx(windowClass *WNDCLASSEX) kernel32.ATOM {
	ret, _, _ := syscall.Syscall(registerClassEx.Addr(), 1,
		uintptr(unsafe.Pointer(windowClass)),
		0,
		0)

	return kernel32.ATOM(ret)
}

func RegisterRawInputDevices(pRawInputDevices *RAWINPUTDEVICE, uiNumDevices uint32, cbSize uint32) bool {
	ret, _, _ := syscall.Syscall(registerRawInputDevices.Addr(), 3,
		uintptr(unsafe.Pointer(pRawInputDevices)),
		uintptr(uiNumDevices),
		uintptr(cbSize))

	return ret != 0
}

func RegisterWindowMessage(lpString *uint16) uint32 {
	ret, _, _ := syscall.Syscall(registerWindowMessage.Addr(), 1,
		uintptr(unsafe.Pointer(lpString)),
		0,
		0)

	return uint32(ret)
}

func ReleaseCapture() bool {
	ret, _, _ := syscall.Syscall(releaseCapture.Addr(), 0,
		0,
		0,
		0)

	return ret != 0
}

func ReleaseDC(hWnd handle.HWND, hDC gdi32.HDC) bool {
	ret, _, _ := syscall.Syscall(releaseDC.Addr(), 2,
		uintptr(hWnd),
		uintptr(hDC),
		0)

	return ret != 0
}

func RemoveMenu(hMenu winuser.HMENU, uPosition, uFlags uint32) bool {
	ret, _, _ := syscall.Syscall(removeMenu.Addr(), 3,
		uintptr(hMenu),
		uintptr(uPosition),
		uintptr(uFlags))

	return ret != 0
}

func ScreenToClient(hWnd handle.HWND, point *gdi32.POINT) bool {
	ret, _, _ := syscall.Syscall(screenToClient.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(point)),
		0)

	return ret != 0
}

func SendDlgItemMessage(hWnd handle.HWND, id int32, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(sendDlgItemMessage.Addr(), 5,
		uintptr(hWnd),
		uintptr(id),
		uintptr(msg),
		wParam,
		lParam,
		0)

	return ret
}

// pInputs expects a unsafe.Pointer to a slice of MOUSE_INPUT or KEYBD_INPUT or HARDWARE_INPUT structs.
func SendInput(nInputs uint32, pInputs unsafe.Pointer, cbSize int32) uint32 {
	ret, _, _ := syscall.Syscall(sendInput.Addr(), 3,
		uintptr(nInputs),
		uintptr(pInputs),
		uintptr(cbSize))

	return uint32(ret)
}

func SendMessage(hWnd handle.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(sendMessage.Addr(), 4,
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		0,
		0)

	return ret
}

func SetActiveWindow(hWnd handle.HWND) handle.HWND {
	ret, _, _ := syscall.Syscall(setActiveWindow.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return handle.HWND(ret)
}

func SetCapture(hWnd handle.HWND) handle.HWND {
	ret, _, _ := syscall.Syscall(setCapture.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return handle.HWND(ret)
}

func SetClipboardData(uFormat uint32, hMem handle.HANDLE) handle.HANDLE {
	ret, _, _ := syscall.Syscall(setClipboardData.Addr(), 2,
		uintptr(uFormat),
		uintptr(hMem),
		0)

	return handle.HANDLE(ret)
}

func SetCursor(hCursor HCURSOR) HCURSOR {
	ret, _, _ := syscall.Syscall(setCursor.Addr(), 1,
		uintptr(hCursor),
		0,
		0)

	return HCURSOR(ret)
}

func SetCursorPos(X, Y int32) bool {
	ret, _, _ := syscall.Syscall(setCursorPos.Addr(), 2,
		uintptr(X),
		uintptr(Y),
		0)

	return ret != 0
}

func SetFocus(hWnd handle.HWND) handle.HWND {
	ret, _, _ := syscall.Syscall(setFocus.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return handle.HWND(ret)
}

func SetForegroundWindow(hWnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(setForegroundWindow.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func SetLayeredWindowAttributes(hwnd handle.HWND, crKey gdi32.COLORREF, bAlpha byte, dwFlags uint32) bool {

	ret, _, _ := syscall.Syscall6(setLayeredWindowAttributes.Addr(), 4,
		uintptr(hwnd),
		uintptr(crKey),
		uintptr(bAlpha),
		uintptr(dwFlags),
		0,
		0)

	return ret != 0
}

func SetMenu(hWnd handle.HWND, hMenu winuser.HMENU) bool {
	ret, _, _ := syscall.Syscall(setMenu.Addr(), 2,
		uintptr(hWnd),
		uintptr(hMenu),
		0)

	return ret != 0
}

func SetMenuDefaultItem(hMenu winuser.HMENU, uItem uint32, fByPosition bool) bool {
	ret, _, _ := syscall.Syscall(setMenuDefaultItem.Addr(), 3,
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(win.BoolToBOOL(fByPosition)))

	return ret != 0
}

func SetMenuInfo(hmenu winuser.HMENU, lpcmi *winuser.MENUINFO) bool {
	ret, _, _ := syscall.Syscall(setMenuInfo.Addr(), 2,
		uintptr(hmenu),
		uintptr(unsafe.Pointer(lpcmi)),
		0)

	return ret != 0
}

func SetMenuItemBitmaps(hMenu winuser.HMENU, uPosition uint32, uFlags uint32, hBitmapUnchecked gdi32.HBITMAP, hBitmapChecked gdi32.HBITMAP) bool {
	ret, _, _ := syscall.Syscall6(setMenuItemBitmaps.Addr(), 5,
		uintptr(hMenu),
		uintptr(uPosition),
		uintptr(uFlags),
		uintptr(hBitmapUnchecked),
		uintptr(hBitmapChecked),
		0)

	return ret != 0
}

func SetMenuItemInfo(hMenu winuser.HMENU, uItem uint32, fByPosition bool, lpmii *winuser.MENUITEMINFO) bool {
	ret, _, _ := syscall.Syscall6(setMenuItemInfo.Addr(), 4,
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(win.BoolToBOOL(fByPosition)),
		uintptr(unsafe.Pointer(lpmii)),
		0,
		0)

	return ret != 0
}

func SetParent(hWnd handle.HWND, parentHWnd handle.HWND) handle.HWND {
	ret, _, _ := syscall.Syscall(setParent.Addr(), 2,
		uintptr(hWnd),
		uintptr(parentHWnd),
		0)

	return handle.HWND(ret)
}

func SetRect(lprc *gdi32.RECT, xLeft, yTop, xRight, yBottom uint32) win.BOOL {
	ret, _, _ := syscall.Syscall6(setRect.Addr(), 5,
		uintptr(unsafe.Pointer(lprc)),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(xRight),
		uintptr(yBottom),
		0)

	return win.BOOL(ret)
}

func SetScrollInfo(hwnd handle.HWND, fnBar int32, lpsi *SCROLLINFO, fRedraw bool) int32 {
	ret, _, _ := syscall.Syscall6(setScrollInfo.Addr(), 4,
		uintptr(hwnd),
		uintptr(fnBar),
		uintptr(unsafe.Pointer(lpsi)),
		uintptr(win.BoolToBOOL(fRedraw)),
		0,
		0)

	return int32(ret)
}

func SetTimer(hWnd handle.HWND, nIDEvent uintptr, uElapse uint32, lpTimerFunc uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(setTimer.Addr(), 4,
		uintptr(hWnd),
		nIDEvent,
		uintptr(uElapse),
		lpTimerFunc,
		0,
		0)

	return ret
}

func SetWinEventHook(eventMin uint32, eventMax uint32, hmodWinEventProc kernel32.HMODULE, callbackFunction win.WINEVENTPROC, idProcess uint32, idThread uint32, dwFlags uint32) (kernel32.HWINEVENTHOOK, error) {
	ret, _, err := syscall.Syscall9(setWinEventHook.Addr(), 7,
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		windows.NewCallback(callbackFunction),
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
		0, 0)

	if ret == 0 {
		return 0, err
	}

	return kernel32.HWINEVENTHOOK(ret), nil
}

func SetWindowLong(hWnd handle.HWND, index, value int32) int32 {
	ret, _, _ := syscall.Syscall(setWindowLong.Addr(), 3,
		uintptr(hWnd),
		uintptr(index),
		uintptr(value))

	return int32(ret)
}

func SetWindowLongPtr(hWnd handle.HWND, index int, value uintptr) uintptr {
	ret, _, _ := syscall.Syscall(setWindowLongPtr.Addr(), 3,
		uintptr(hWnd),
		uintptr(index),
		value)

	return ret
}

func SetWindowPlacement(hWnd handle.HWND, lpwndpl *WINDOWPLACEMENT) bool {
	ret, _, _ := syscall.Syscall(setWindowPlacement.Addr(), 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpwndpl)),
		0)

	return ret != 0
}

func SetWindowPos(hWnd, hWndInsertAfter handle.HWND, x, y, width, height int32, flags uint32) bool {
	ret, _, _ := syscall.Syscall9(setWindowPos.Addr(), 7,
		uintptr(hWnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
		0,
		0)

	return ret != 0
}

func ShowWindow(hWnd handle.HWND, nCmdShow int32) bool {
	ret, _, _ := syscall.Syscall(showWindow.Addr(), 2,
		uintptr(hWnd),
		uintptr(nCmdShow),
		0)

	return ret != 0
}

func SystemParametersInfo(uiAction, uiParam uint32, pvParam unsafe.Pointer, fWinIni uint32) bool {
	ret, _, _ := syscall.Syscall6(systemParametersInfo.Addr(), 4,
		uintptr(uiAction),
		uintptr(uiParam),
		uintptr(pvParam),
		uintptr(fWinIni),
		0,
		0)

	return ret != 0
}

func TrackMouseEvent(lpEventTrack *TRACKMOUSEEVENT) bool {
	ret, _, _ := syscall.Syscall(trackMouseEvent.Addr(), 1,
		uintptr(unsafe.Pointer(lpEventTrack)),
		0,
		0)

	return ret != 0
}

func TrackPopupMenu(hMenu winuser.HMENU, uFlags uint32, x, y int32, nReserved int32, hWnd handle.HWND, prcRect *gdi32.RECT) uint32 {
	ret, _, _ := syscall.Syscall9(trackPopupMenu.Addr(), 7,
		uintptr(hMenu),
		uintptr(uFlags),
		uintptr(x),
		uintptr(y),
		uintptr(nReserved),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(prcRect)),
		0,
		0)

	return uint32(ret)
}

func TrackPopupMenuEx(hMenu winuser.HMENU, fuFlags uint32, x, y int32, hWnd handle.HWND, lptpm *TPMPARAMS) win.BOOL {
	ret, _, _ := syscall.Syscall6(trackPopupMenuEx.Addr(), 6,
		uintptr(hMenu),
		uintptr(fuFlags),
		uintptr(x),
		uintptr(y),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lptpm)))

	return win.BOOL(ret)
}

func TranslateMessage(msg *MSG) bool {
	ret, _, _ := syscall.Syscall(translateMessage.Addr(), 1,
		uintptr(unsafe.Pointer(msg)),
		0,
		0)

	return ret != 0
}

func TranslateMessageWV2(msg *MSG) error {
	_, _, err := translateMessage.Call(uintptr(unsafe.Pointer(msg)))
	if err != nil && !errors.Is(err, syscall.Errno(0)) {
		return err
	}

	return nil
}

func UnhookWinEvent(hWinHookEvent kernel32.HWINEVENTHOOK) bool {
	ret, _, _ := syscall.Syscall(unhookWinEvent.Addr(), 1, uintptr(hWinHookEvent), 0, 0)
	return ret != 0
}

func UpdateWindow(hwnd handle.HWND) bool {
	ret, _, _ := syscall.Syscall(updateWindow.Addr(), 1,
		uintptr(hwnd),
		0,
		0)

	return ret != 0
}

func WindowFromDC(hDC gdi32.HDC) handle.HWND {
	ret, _, _ := syscall.Syscall(windowFromDC.Addr(), 1,
		uintptr(hDC),
		0,
		0)

	return handle.HWND(ret)
}

func WindowFromPoint(Point gdi32.POINT) handle.HWND {
	ret, _, _ := syscall.Syscall(windowFromPoint.Addr(), 2,
		uintptr(Point.X),
		uintptr(Point.Y),
		0)

	return handle.HWND(ret)
}
