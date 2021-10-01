// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package user32

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	// Library
	libuser32 *windows.LazyDLL

	// Functions
	addClipboardFormatListener  *windows.LazyProc
	adjustWindowRect            *windows.LazyProc
	attachThreadInput           *windows.LazyProc
	animateWindow               *windows.LazyProc
	beginDeferWindowPos         *windows.LazyProc
	beginPaint                  *windows.LazyProc
	bringWindowToTop            *windows.LazyProc
	callWindowProc              *windows.LazyProc
	changeWindowMessageFilterEx *windows.LazyProc
	checkMenuRadioItem          *windows.LazyProc
	clientToScreen              *windows.LazyProc
	closeClipboard              *windows.LazyProc
	createDialogParam           *windows.LazyProc
	createIconIndirect          *windows.LazyProc
	createMenu                  *windows.LazyProc
	createPopupMenu             *windows.LazyProc
	createWindowEx              *windows.LazyProc
	deferWindowPos              *windows.LazyProc
	defWindowProc               *windows.LazyProc
	deleteMenu                  *windows.LazyProc
	destroyIcon                 *windows.LazyProc
	destroyMenu                 *windows.LazyProc
	destroyWindow               *windows.LazyProc
	dialogBoxParam              *windows.LazyProc
	dispatchMessage             *windows.LazyProc
	drawIconEx                  *windows.LazyProc
	drawMenuBar                 *windows.LazyProc
	drawFocusRect               *windows.LazyProc
	drawTextEx                  *windows.LazyProc
	emptyClipboard              *windows.LazyProc
	enableMenuItem              *windows.LazyProc
	enableWindow                *windows.LazyProc
	endDeferWindowPos           *windows.LazyProc
	endDialog                   *windows.LazyProc
	endPaint                    *windows.LazyProc
	enumChildWindows            *windows.LazyProc
	findWindow                  *windows.LazyProc
	getActiveWindow             *windows.LazyProc
	getAncestor                 *windows.LazyProc
	getCaretPos                 *windows.LazyProc
	getClassName                *windows.LazyProc
	getClientRect               *windows.LazyProc
	getClipboardData            *windows.LazyProc
	getCursorPos                *windows.LazyProc
	getDC                       *windows.LazyProc
	getDesktopWindow            *windows.LazyProc
	getDlgItem                  *windows.LazyProc
	getDpiForWindow             *windows.LazyProc
	getFocus                    *windows.LazyProc
	getForegroundWindow         *windows.LazyProc
	getIconInfo                 *windows.LazyProc
	getKeyState                 *windows.LazyProc
	getMenuCheckMarkDimensions  *windows.LazyProc
	getMenuInfo                 *windows.LazyProc
	getMenuItemCount            *windows.LazyProc
	getMenuItemID               *windows.LazyProc
	getMenuItemInfo             *windows.LazyProc
	getMessage                  *windows.LazyProc
	getMonitorInfo              *windows.LazyProc
	getParent                   *windows.LazyProc
	getRawInputData             *windows.LazyProc
	getScrollInfo               *windows.LazyProc
	getSubMenu                  *windows.LazyProc
	getSysColor                 *windows.LazyProc
	getSysColorBrush            *windows.LazyProc
	getSystemMenu               *windows.LazyProc
	getSystemMetrics            *windows.LazyProc
	getSystemMetricsForDpi      *windows.LazyProc
	getWindow                   *windows.LazyProc
	getWindowLong               *windows.LazyProc
	getWindowLongPtr            *windows.LazyProc
	getWindowPlacement          *windows.LazyProc
	getWindowRect               *windows.LazyProc
	getWindowThreadProcessId    *windows.LazyProc
	insertMenuItem              *windows.LazyProc
	invalidateRect              *windows.LazyProc
	isChild                     *windows.LazyProc
	isClipboardFormatAvailable  *windows.LazyProc
	isDialogMessage             *windows.LazyProc
	isIconic                    *windows.LazyProc
	isWindowEnabled             *windows.LazyProc
	isWindowVisible             *windows.LazyProc
	isZoomed                    *windows.LazyProc
	killTimer                   *windows.LazyProc
	loadCursor                  *windows.LazyProc
	loadIcon                    *windows.LazyProc
	loadImage                   *windows.LazyProc
	loadMenu                    *windows.LazyProc
	loadString                  *windows.LazyProc
	messageBeep                 *windows.LazyProc
	messageBox                  *windows.LazyProc
	monitorFromWindow           *windows.LazyProc
	moveWindow                  *windows.LazyProc
	notifyWinEvent              *windows.LazyProc
	unregisterClass             *windows.LazyProc
	openClipboard               *windows.LazyProc
	peekMessage                 *windows.LazyProc
	postMessage                 *windows.LazyProc
	postQuitMessage             *windows.LazyProc
	redrawWindow                *windows.LazyProc
	registerClassEx             *windows.LazyProc
	registerRawInputDevices     *windows.LazyProc
	registerWindowMessage       *windows.LazyProc
	releaseCapture              *windows.LazyProc
	releaseDC                   *windows.LazyProc
	removeMenu                  *windows.LazyProc
	screenToClient              *windows.LazyProc
	sendDlgItemMessage          *windows.LazyProc
	sendInput                   *windows.LazyProc
	sendMessage                 *windows.LazyProc
	setActiveWindow             *windows.LazyProc
	setCapture                  *windows.LazyProc
	setClipboardData            *windows.LazyProc
	setCursor                   *windows.LazyProc
	setCursorPos                *windows.LazyProc
	setFocus                    *windows.LazyProc
	setForegroundWindow         *windows.LazyProc
	setMenu                     *windows.LazyProc
	setMenuDefaultItem          *windows.LazyProc
	setMenuInfo                 *windows.LazyProc
	setMenuItemBitmaps          *windows.LazyProc
	setMenuItemInfo             *windows.LazyProc
	setParent                   *windows.LazyProc
	setRect                     *windows.LazyProc
	setScrollInfo               *windows.LazyProc
	setTimer                    *windows.LazyProc
	setWinEventHook             *windows.LazyProc
	setWindowLong               *windows.LazyProc
	setWindowLongPtr            *windows.LazyProc
	setWindowPlacement          *windows.LazyProc
	setWindowPos                *windows.LazyProc
	showWindow                  *windows.LazyProc
	systemParametersInfo        *windows.LazyProc
	trackMouseEvent             *windows.LazyProc
	trackPopupMenu              *windows.LazyProc
	trackPopupMenuEx            *windows.LazyProc
	translateMessage            *windows.LazyProc
	unhookWinEvent              *windows.LazyProc
	updateWindow                *windows.LazyProc
	windowFromDC                *windows.LazyProc
	windowFromPoint             *windows.LazyProc
)

func init() {
	is64bit := unsafe.Sizeof(uintptr(0)) == 8

	// Library
	libuser32 = windows.NewLazySystemDLL("user32.dll")

	// Functions
	addClipboardFormatListener = libuser32.NewProc("AddClipboardFormatListener")
	adjustWindowRect = libuser32.NewProc("AdjustWindowRect")
	attachThreadInput = libuser32.NewProc("AttachThreadInput")
	animateWindow = libuser32.NewProc("AnimateWindow")
	beginDeferWindowPos = libuser32.NewProc("BeginDeferWindowPos")
	beginPaint = libuser32.NewProc("BeginPaint")
	bringWindowToTop = libuser32.NewProc("BringWindowToTop")
	callWindowProc = libuser32.NewProc("CallWindowProcW")
	changeWindowMessageFilterEx = libuser32.NewProc("ChangeWindowMessageFilterEx")
	checkMenuRadioItem = libuser32.NewProc("CheckMenuRadioItem")
	clientToScreen = libuser32.NewProc("ClientToScreen")
	closeClipboard = libuser32.NewProc("CloseClipboard")
	createDialogParam = libuser32.NewProc("CreateDialogParamW")
	createIconIndirect = libuser32.NewProc("CreateIconIndirect")
	createMenu = libuser32.NewProc("CreateMenu")
	createPopupMenu = libuser32.NewProc("CreatePopupMenu")
	createWindowEx = libuser32.NewProc("CreateWindowExW")
	deferWindowPos = libuser32.NewProc("DeferWindowPos")
	defWindowProc = libuser32.NewProc("DefWindowProcW")
	deleteMenu = libuser32.NewProc("DeleteMenu")
	destroyIcon = libuser32.NewProc("DestroyIcon")
	destroyMenu = libuser32.NewProc("DestroyMenu")
	destroyWindow = libuser32.NewProc("DestroyWindow")
	dialogBoxParam = libuser32.NewProc("DialogBoxParamW")
	dispatchMessage = libuser32.NewProc("DispatchMessageW")
	drawIconEx = libuser32.NewProc("DrawIconEx")
	drawFocusRect = libuser32.NewProc("DrawFocusRect")
	drawMenuBar = libuser32.NewProc("DrawMenuBar")
	drawTextEx = libuser32.NewProc("DrawTextExW")
	emptyClipboard = libuser32.NewProc("EmptyClipboard")
	enableMenuItem = libuser32.NewProc("EnableMenuItem")
	enableWindow = libuser32.NewProc("EnableWindow")
	endDeferWindowPos = libuser32.NewProc("EndDeferWindowPos")
	endDialog = libuser32.NewProc("EndDialog")
	endPaint = libuser32.NewProc("EndPaint")
	enumChildWindows = libuser32.NewProc("EnumChildWindows")
	findWindow = libuser32.NewProc("FindWindowW")
	getActiveWindow = libuser32.NewProc("GetActiveWindow")
	getAncestor = libuser32.NewProc("GetAncestor")
	getCaretPos = libuser32.NewProc("GetCaretPos")
	getClassName = libuser32.NewProc("GetClassNameW")
	getClientRect = libuser32.NewProc("GetClientRect")
	getClipboardData = libuser32.NewProc("GetClipboardData")
	getCursorPos = libuser32.NewProc("GetCursorPos")
	getDC = libuser32.NewProc("GetDC")
	getDesktopWindow = libuser32.NewProc("GetDesktopWindow")
	getDlgItem = libuser32.NewProc("GetDlgItem")
	getDpiForWindow = libuser32.NewProc("GetDpiForWindow")
	getFocus = libuser32.NewProc("GetFocus")
	getForegroundWindow = libuser32.NewProc("GetForegroundWindow")
	getIconInfo = libuser32.NewProc("GetIconInfo")
	getKeyState = libuser32.NewProc("GetKeyState")
	getMenuCheckMarkDimensions = libuser32.NewProc("GetMenuCheckMarkDimensions")
	getMenuInfo = libuser32.NewProc("GetMenuInfo")
	getMenuItemCount = libuser32.NewProc("GetMenuItemCount")
	getMenuItemID = libuser32.NewProc("GetMenuItemID")
	getMenuItemInfo = libuser32.NewProc("GetMenuItemInfoW")
	getMessage = libuser32.NewProc("GetMessageW")
	getMonitorInfo = libuser32.NewProc("GetMonitorInfoW")
	getParent = libuser32.NewProc("GetParent")
	getRawInputData = libuser32.NewProc("GetRawInputData")
	getScrollInfo = libuser32.NewProc("GetScrollInfo")
	getSubMenu = libuser32.NewProc("GetSubMenu")
	getSysColor = libuser32.NewProc("GetSysColor")
	getSysColorBrush = libuser32.NewProc("GetSysColorBrush")
	getSystemMenu = libuser32.NewProc("GetSystemMenu")
	getSystemMetrics = libuser32.NewProc("GetSystemMetrics")
	getSystemMetricsForDpi = libuser32.NewProc("GetSystemMetricsForDpi")
	getWindow = libuser32.NewProc("GetWindow")
	getWindowLong = libuser32.NewProc("GetWindowLongW")
	// On 32 bit GetWindowLongPtrW is not available
	if is64bit {
		getWindowLongPtr = libuser32.NewProc("GetWindowLongPtrW")
	} else {
		getWindowLongPtr = libuser32.NewProc("GetWindowLongW")
	}
	getWindowPlacement = libuser32.NewProc("GetWindowPlacement")
	getWindowRect = libuser32.NewProc("GetWindowRect")
	getWindowThreadProcessId = libuser32.NewProc("GetWindowThreadProcessId")
	insertMenuItem = libuser32.NewProc("InsertMenuItemW")
	invalidateRect = libuser32.NewProc("InvalidateRect")
	isChild = libuser32.NewProc("IsChild")
	isClipboardFormatAvailable = libuser32.NewProc("IsClipboardFormatAvailable")
	isDialogMessage = libuser32.NewProc("IsDialogMessageW")
	isIconic = libuser32.NewProc("IsIconic")
	isWindowEnabled = libuser32.NewProc("IsWindowEnabled")
	isWindowVisible = libuser32.NewProc("IsWindowVisible")
	isZoomed = libuser32.NewProc("IsZoomed")
	killTimer = libuser32.NewProc("KillTimer")
	loadCursor = libuser32.NewProc("LoadCursorW")
	loadIcon = libuser32.NewProc("LoadIconW")
	loadImage = libuser32.NewProc("LoadImageW")
	loadMenu = libuser32.NewProc("LoadMenuW")
	loadString = libuser32.NewProc("LoadStringW")
	messageBeep = libuser32.NewProc("MessageBeep")
	messageBox = libuser32.NewProc("MessageBoxW")
	monitorFromWindow = libuser32.NewProc("MonitorFromWindow")
	moveWindow = libuser32.NewProc("MoveWindow")
	notifyWinEvent = libuser32.NewProc("NotifyWinEvent")
	unregisterClass = libuser32.NewProc("UnregisterClassW")
	openClipboard = libuser32.NewProc("OpenClipboard")
	peekMessage = libuser32.NewProc("PeekMessageW")
	postMessage = libuser32.NewProc("PostMessageW")
	postQuitMessage = libuser32.NewProc("PostQuitMessage")
	redrawWindow = libuser32.NewProc("RedrawWindow")
	registerClassEx = libuser32.NewProc("RegisterClassExW")
	registerRawInputDevices = libuser32.NewProc("RegisterRawInputDevices")
	registerWindowMessage = libuser32.NewProc("RegisterWindowMessageW")
	releaseCapture = libuser32.NewProc("ReleaseCapture")
	releaseDC = libuser32.NewProc("ReleaseDC")
	removeMenu = libuser32.NewProc("RemoveMenu")
	screenToClient = libuser32.NewProc("ScreenToClient")
	sendDlgItemMessage = libuser32.NewProc("SendDlgItemMessageW")
	sendInput = libuser32.NewProc("SendInput")
	sendMessage = libuser32.NewProc("SendMessageW")
	setActiveWindow = libuser32.NewProc("SetActiveWindow")
	setCapture = libuser32.NewProc("SetCapture")
	setClipboardData = libuser32.NewProc("SetClipboardData")
	setCursor = libuser32.NewProc("SetCursor")
	setCursorPos = libuser32.NewProc("SetCursorPos")
	setFocus = libuser32.NewProc("SetFocus")
	setForegroundWindow = libuser32.NewProc("SetForegroundWindow")
	setMenu = libuser32.NewProc("SetMenu")
	setMenuDefaultItem = libuser32.NewProc("SetMenuDefaultItem")
	setMenuInfo = libuser32.NewProc("SetMenuInfo")
	setMenuItemBitmaps = libuser32.NewProc("SetMenuItemBitmaps")
	setMenuItemInfo = libuser32.NewProc("SetMenuItemInfoW")
	setRect = libuser32.NewProc("SetRect")
	setParent = libuser32.NewProc("SetParent")
	setScrollInfo = libuser32.NewProc("SetScrollInfo")
	setTimer = libuser32.NewProc("SetTimer")
	setWinEventHook = libuser32.NewProc("SetWinEventHook")
	setWindowLong = libuser32.NewProc("SetWindowLongW")
	// On 32 bit SetWindowLongPtrW is not available
	if is64bit {
		setWindowLongPtr = libuser32.NewProc("SetWindowLongPtrW")
	} else {
		setWindowLongPtr = libuser32.NewProc("SetWindowLongW")
	}
	setWindowPlacement = libuser32.NewProc("SetWindowPlacement")
	setWindowPos = libuser32.NewProc("SetWindowPos")
	showWindow = libuser32.NewProc("ShowWindow")
	systemParametersInfo = libuser32.NewProc("SystemParametersInfoW")
	trackMouseEvent = libuser32.NewProc("TrackMouseEvent")
	trackPopupMenu = libuser32.NewProc("TrackPopupMenu")
	trackPopupMenuEx = libuser32.NewProc("TrackPopupMenuEx")
	translateMessage = libuser32.NewProc("TranslateMessage")
	unhookWinEvent = libuser32.NewProc("UnhookWinEvent")
	updateWindow = libuser32.NewProc("UpdateWindow")
	windowFromDC = libuser32.NewProc("WindowFromDC")
	windowFromPoint = libuser32.NewProc("WindowFromPoint")
}
