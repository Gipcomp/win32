// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package gdi32

import "golang.org/x/sys/windows"

var (
	// Library
	libgdi32   *windows.LazyDLL
	libmsimg32 *windows.LazyDLL

	// Functions
	abortDoc                *windows.LazyProc
	addFontResourceEx       *windows.LazyProc
	addFontMemResourceEx    *windows.LazyProc
	alphaBlend              *windows.LazyProc
	bitBlt                  *windows.LazyProc
	choosePixelFormat       *windows.LazyProc
	closeEnhMetaFile        *windows.LazyProc
	combineRgn              *windows.LazyProc
	copyEnhMetaFile         *windows.LazyProc
	createBitmap            *windows.LazyProc
	createCompatibleBitmap  *windows.LazyProc
	createBrushIndirect     *windows.LazyProc
	createCompatibleDC      *windows.LazyProc
	createDC                *windows.LazyProc
	createDIBSection        *windows.LazyProc
	createFontIndirect      *windows.LazyProc
	createEnhMetaFile       *windows.LazyProc
	createIC                *windows.LazyProc
	createPatternBrush      *windows.LazyProc
	createRectRgn           *windows.LazyProc
	deleteDC                *windows.LazyProc
	deleteEnhMetaFile       *windows.LazyProc
	deleteObject            *windows.LazyProc
	ellipse                 *windows.LazyProc
	endDoc                  *windows.LazyProc
	endPage                 *windows.LazyProc
	excludeClipRect         *windows.LazyProc
	extCreatePen            *windows.LazyProc
	fillRgn                 *windows.LazyProc
	gdiFlush                *windows.LazyProc
	getBkColor              *windows.LazyProc
	getDeviceCaps           *windows.LazyProc
	getDIBits               *windows.LazyProc
	getEnhMetaFile          *windows.LazyProc
	getEnhMetaFileHeader    *windows.LazyProc
	getObject               *windows.LazyProc
	getPixel                *windows.LazyProc
	getRgnBox               *windows.LazyProc
	getStockObject          *windows.LazyProc
	getTextColor            *windows.LazyProc
	getTextExtentExPoint    *windows.LazyProc
	getTextExtentPoint32    *windows.LazyProc
	getTextMetrics          *windows.LazyProc
	getViewportOrgEx        *windows.LazyProc
	gradientFill            *windows.LazyProc
	intersectClipRect       *windows.LazyProc
	lineTo                  *windows.LazyProc
	moveToEx                *windows.LazyProc
	playEnhMetaFile         *windows.LazyProc
	polyline                *windows.LazyProc
	rectangle               *windows.LazyProc
	removeFontResourceEx    *windows.LazyProc
	removeFontMemResourceEx *windows.LazyProc
	resetDC                 *windows.LazyProc
	restoreDC               *windows.LazyProc
	roundRect               *windows.LazyProc
	selectObject            *windows.LazyProc
	setBkColor              *windows.LazyProc
	setBkMode               *windows.LazyProc
	setBrushOrgEx           *windows.LazyProc
	setDIBits               *windows.LazyProc
	setPixel                *windows.LazyProc
	setPixelFormat          *windows.LazyProc
	setStretchBltMode       *windows.LazyProc
	setTextColor            *windows.LazyProc
	setViewportOrgEx        *windows.LazyProc
	saveDC                  *windows.LazyProc
	startDoc                *windows.LazyProc
	startPage               *windows.LazyProc
	stretchBlt              *windows.LazyProc
	swapBuffers             *windows.LazyProc
	textOut                 *windows.LazyProc
	transparentBlt          *windows.LazyProc
)

func init() {
	// Library
	libgdi32 = windows.NewLazySystemDLL("gdi32.dll")
	libmsimg32 = windows.NewLazySystemDLL("msimg32.dll")

	// Functions
	abortDoc = libgdi32.NewProc("AbortDoc")
	addFontResourceEx = libgdi32.NewProc("AddFontResourceExW")
	addFontMemResourceEx = libgdi32.NewProc("AddFontMemResourceEx")
	bitBlt = libgdi32.NewProc("BitBlt")
	choosePixelFormat = libgdi32.NewProc("ChoosePixelFormat")
	closeEnhMetaFile = libgdi32.NewProc("CloseEnhMetaFile")
	combineRgn = libgdi32.NewProc("CombineRgn")
	copyEnhMetaFile = libgdi32.NewProc("CopyEnhMetaFileW")
	createBitmap = libgdi32.NewProc("CreateBitmap")
	createCompatibleBitmap = libgdi32.NewProc("CreateCompatibleBitmap")
	createBrushIndirect = libgdi32.NewProc("CreateBrushIndirect")
	createCompatibleDC = libgdi32.NewProc("CreateCompatibleDC")
	createDC = libgdi32.NewProc("CreateDCW")
	createDIBSection = libgdi32.NewProc("CreateDIBSection")
	createEnhMetaFile = libgdi32.NewProc("CreateEnhMetaFileW")
	createFontIndirect = libgdi32.NewProc("CreateFontIndirectW")
	createIC = libgdi32.NewProc("CreateICW")
	createPatternBrush = libgdi32.NewProc("CreatePatternBrush")
	createRectRgn = libgdi32.NewProc("CreateRectRgn")
	deleteDC = libgdi32.NewProc("DeleteDC")
	deleteEnhMetaFile = libgdi32.NewProc("DeleteEnhMetaFile")
	deleteObject = libgdi32.NewProc("DeleteObject")
	ellipse = libgdi32.NewProc("Ellipse")
	endDoc = libgdi32.NewProc("EndDoc")
	endPage = libgdi32.NewProc("EndPage")
	excludeClipRect = libgdi32.NewProc("ExcludeClipRect")
	extCreatePen = libgdi32.NewProc("ExtCreatePen")
	fillRgn = libgdi32.NewProc("FillRgn")
	gdiFlush = libgdi32.NewProc("GdiFlush")
	getBkColor = libgdi32.NewProc("GetBkColor")
	getDeviceCaps = libgdi32.NewProc("GetDeviceCaps")
	getDIBits = libgdi32.NewProc("GetDIBits")
	getEnhMetaFile = libgdi32.NewProc("GetEnhMetaFileW")
	getEnhMetaFileHeader = libgdi32.NewProc("GetEnhMetaFileHeader")
	getObject = libgdi32.NewProc("GetObjectW")
	getPixel = libgdi32.NewProc("GetPixel")
	getRgnBox = libgdi32.NewProc("GetRgnBox")
	getStockObject = libgdi32.NewProc("GetStockObject")
	getTextColor = libgdi32.NewProc("GetTextColor")
	getTextExtentExPoint = libgdi32.NewProc("GetTextExtentExPointW")
	getTextExtentPoint32 = libgdi32.NewProc("GetTextExtentPoint32W")
	getTextMetrics = libgdi32.NewProc("GetTextMetricsW")
	getViewportOrgEx = libgdi32.NewProc("GetViewportOrgEx")
	intersectClipRect = libgdi32.NewProc("IntersectClipRect")
	lineTo = libgdi32.NewProc("LineTo")
	moveToEx = libgdi32.NewProc("MoveToEx")
	playEnhMetaFile = libgdi32.NewProc("PlayEnhMetaFile")
	polyline = libgdi32.NewProc("Polyline")
	rectangle = libgdi32.NewProc("Rectangle")
	removeFontResourceEx = libgdi32.NewProc("RemoveFontResourceExW")
	removeFontMemResourceEx = libgdi32.NewProc("RemoveFontMemResourceEx")
	resetDC = libgdi32.NewProc("ResetDCW")
	restoreDC = libgdi32.NewProc("RestoreDC")
	roundRect = libgdi32.NewProc("RoundRect")
	saveDC = libgdi32.NewProc("SaveDC")
	selectObject = libgdi32.NewProc("SelectObject")
	setBkColor = libgdi32.NewProc("SetBkColor")
	setBkMode = libgdi32.NewProc("SetBkMode")
	setBrushOrgEx = libgdi32.NewProc("SetBrushOrgEx")
	setDIBits = libgdi32.NewProc("SetDIBits")
	setPixel = libgdi32.NewProc("SetPixel")
	setPixelFormat = libgdi32.NewProc("SetPixelFormat")
	setStretchBltMode = libgdi32.NewProc("SetStretchBltMode")
	setTextColor = libgdi32.NewProc("SetTextColor")
	setViewportOrgEx = libgdi32.NewProc("SetViewportOrgEx")
	startDoc = libgdi32.NewProc("StartDocW")
	startPage = libgdi32.NewProc("StartPage")
	stretchBlt = libgdi32.NewProc("StretchBlt")
	swapBuffers = libgdi32.NewProc("SwapBuffers")
	textOut = libgdi32.NewProc("TextOutW")

	alphaBlend = libmsimg32.NewProc("AlphaBlend")
	gradientFill = libmsimg32.NewProc("GradientFill")
	transparentBlt = libmsimg32.NewProc("TransparentBlt")
}
