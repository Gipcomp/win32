// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package gdi32

import (
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/handle"
)

func RGB(r, g, b byte) COLORREF {
	return COLORREF(r) | (COLORREF(g) << 8) | (COLORREF(b) << 16)
}

func AbortDoc(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(abortDoc.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return int32(ret)
}

func AddFontResourceEx(lpszFilename *uint16, fl uint32, pdv unsafe.Pointer) int32 {
	ret, _, _ := syscall.Syscall(addFontResourceEx.Addr(), 3,
		uintptr(unsafe.Pointer(lpszFilename)),
		uintptr(fl),
		uintptr(pdv))

	return int32(ret)
}

func AddFontMemResourceEx(pFileView uintptr, cjSize uint32, pvReserved unsafe.Pointer, pNumFonts *uint32) handle.HANDLE {
	ret, _, _ := syscall.Syscall6(addFontMemResourceEx.Addr(), 4,
		pFileView,
		uintptr(cjSize),
		uintptr(pvReserved),
		uintptr(unsafe.Pointer(pNumFonts)),
		0,
		0)

	return handle.HANDLE(ret)
}

func AlphaBlend(hdcDest HDC, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc HDC, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc int32, ftn BLENDFUNCTION) bool {
	ret, _, _ := syscall.Syscall12(alphaBlend.Addr(), 11,
		uintptr(hdcDest),
		uintptr(nXOriginDest),
		uintptr(nYOriginDest),
		uintptr(nWidthDest),
		uintptr(nHeightDest),
		uintptr(hdcSrc),
		uintptr(nXOriginSrc),
		uintptr(nYOriginSrc),
		uintptr(nWidthSrc),
		uintptr(nHeightSrc),
		uintptr(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ftn))))),
		0)

	return ret != 0
}

func BitBlt(hdcDest HDC, nXDest, nYDest, nWidth, nHeight int32, hdcSrc HDC, nXSrc, nYSrc int32, dwRop uint32) bool {
	ret, _, _ := syscall.Syscall9(bitBlt.Addr(), 9,
		uintptr(hdcDest),
		uintptr(nXDest),
		uintptr(nYDest),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hdcSrc),
		uintptr(nXSrc),
		uintptr(nYSrc),
		uintptr(dwRop))

	return ret != 0
}

func ChoosePixelFormat(hdc HDC, ppfd *PIXELFORMATDESCRIPTOR) int32 {
	ret, _, _ := syscall.Syscall(choosePixelFormat.Addr(), 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(ppfd)),
		0)

	return int32(ret)
}

func CloseEnhMetaFile(hdc HDC) HENHMETAFILE {
	ret, _, _ := syscall.Syscall(closeEnhMetaFile.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return HENHMETAFILE(ret)
}

func CombineRgn(hrgnDest, hrgnSrc1, hrgnSrc2 HRGN, fnCombineMode int32) int32 {
	ret, _, _ := syscall.Syscall6(combineRgn.Addr(), 4,
		uintptr(hrgnDest),
		uintptr(hrgnSrc1),
		uintptr(hrgnSrc2),
		uintptr(fnCombineMode),
		0,
		0)

	return int32(ret)
}

func CopyEnhMetaFile(hemfSrc HENHMETAFILE, lpszFile *uint16) HENHMETAFILE {
	ret, _, _ := syscall.Syscall(copyEnhMetaFile.Addr(), 2,
		uintptr(hemfSrc),
		uintptr(unsafe.Pointer(lpszFile)),
		0)

	return HENHMETAFILE(ret)
}

func CreateBitmap(nWidth, nHeight int32, cPlanes, cBitsPerPel uint32, lpvBits unsafe.Pointer) HBITMAP {
	ret, _, _ := syscall.Syscall6(createBitmap.Addr(), 5,
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(cPlanes),
		uintptr(cBitsPerPel),
		uintptr(lpvBits),
		0)

	return HBITMAP(ret)
}

func CreateCompatibleBitmap(hdc HDC, nWidth, nHeight int32) HBITMAP {
	ret, _, _ := syscall.Syscall(createCompatibleBitmap.Addr(), 3,
		uintptr(hdc),
		uintptr(nWidth),
		uintptr(nHeight))

	return HBITMAP(ret)
}

func CreateBrushIndirect(lplb *LOGBRUSH) HBRUSH {
	ret, _, _ := syscall.Syscall(createBrushIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lplb)),
		0,
		0)

	return HBRUSH(ret)
}

func CreateCompatibleDC(hdc HDC) HDC {
	ret, _, _ := syscall.Syscall(createCompatibleDC.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return HDC(ret)
}

func CreateDC(lpszDriver, lpszDevice, lpszOutput *uint16, lpInitData *DEVMODE) HDC {
	ret, _, _ := syscall.Syscall6(createDC.Addr(), 4,
		uintptr(unsafe.Pointer(lpszDriver)),
		uintptr(unsafe.Pointer(lpszDevice)),
		uintptr(unsafe.Pointer(lpszOutput)),
		uintptr(unsafe.Pointer(lpInitData)),
		0,
		0)

	return HDC(ret)
}

func CreateDIBSection(hdc HDC, pbmih *BITMAPINFOHEADER, iUsage uint32, ppvBits *unsafe.Pointer, hSection handle.HANDLE, dwOffset uint32) HBITMAP {
	ret, _, _ := syscall.Syscall6(createDIBSection.Addr(), 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(pbmih)),
		uintptr(iUsage),
		uintptr(unsafe.Pointer(ppvBits)),
		uintptr(hSection),
		uintptr(dwOffset))

	return HBITMAP(ret)
}

func CreateEnhMetaFile(hdcRef HDC, lpFilename *uint16, lpRect *RECT, lpDescription *uint16) HDC {
	ret, _, _ := syscall.Syscall6(createEnhMetaFile.Addr(), 4,
		uintptr(hdcRef),
		uintptr(unsafe.Pointer(lpFilename)),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(unsafe.Pointer(lpDescription)),
		0,
		0)

	return HDC(ret)
}

func CreateFontIndirect(lplf *LOGFONT) HFONT {
	ret, _, _ := syscall.Syscall(createFontIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lplf)),
		0,
		0)

	return HFONT(ret)
}

func CreateIC(lpszDriver, lpszDevice, lpszOutput *uint16, lpdvmInit *DEVMODE) HDC {
	ret, _, _ := syscall.Syscall6(createIC.Addr(), 4,
		uintptr(unsafe.Pointer(lpszDriver)),
		uintptr(unsafe.Pointer(lpszDevice)),
		uintptr(unsafe.Pointer(lpszOutput)),
		uintptr(unsafe.Pointer(lpdvmInit)),
		0,
		0)

	return HDC(ret)
}

func CreatePatternBrush(hbmp HBITMAP) HBRUSH {
	ret, _, _ := syscall.Syscall(createPatternBrush.Addr(), 1,
		uintptr(hbmp),
		0,
		0)

	return HBRUSH(ret)
}

func CreateRectRgn(nLeftRect, nTopRect, nRightRect, nBottomRect int32) HRGN {
	ret, _, _ := syscall.Syscall6(createRectRgn.Addr(), 4,
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		0,
		0)

	return HRGN(ret)
}

func DeleteDC(hdc HDC) bool {
	ret, _, _ := syscall.Syscall(deleteDC.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return ret != 0
}

func DeleteEnhMetaFile(hemf HENHMETAFILE) bool {
	ret, _, _ := syscall.Syscall(deleteEnhMetaFile.Addr(), 1,
		uintptr(hemf),
		0,
		0)

	return ret != 0
}

func DeleteObject(hObject HGDIOBJ) bool {
	ret, _, _ := syscall.Syscall(deleteObject.Addr(), 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}

func Ellipse(hdc HDC, nLeftRect, nTopRect, nRightRect, nBottomRect int32) bool {
	ret, _, _ := syscall.Syscall6(ellipse.Addr(), 5,
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		0)

	return ret != 0
}

func EndDoc(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(endDoc.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return int32(ret)
}

func EndPage(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(endPage.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return int32(ret)
}

func ExcludeClipRect(hdc HDC, nLeftRect, nTopRect, nRightRect, nBottomRect int32) int32 {
	ret, _, _ := syscall.Syscall6(excludeClipRect.Addr(), 5,
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		0)

	return int32(ret)
}

func ExtCreatePen(dwPenStyle, dwWidth uint32, lplb *LOGBRUSH, dwStyleCount uint32, lpStyle *uint32) HPEN {
	ret, _, _ := syscall.Syscall6(extCreatePen.Addr(), 5,
		uintptr(dwPenStyle),
		uintptr(dwWidth),
		uintptr(unsafe.Pointer(lplb)),
		uintptr(dwStyleCount),
		uintptr(unsafe.Pointer(lpStyle)),
		0)

	return HPEN(ret)
}

func FillRgn(hdc HDC, hrgn HRGN, hbr HBRUSH) bool {
	ret, _, _ := syscall.Syscall(fillRgn.Addr(), 3,
		uintptr(hdc),
		uintptr(hrgn),
		uintptr(hbr))

	return ret != 0
}

func GdiFlush() bool {
	ret, _, _ := syscall.Syscall(gdiFlush.Addr(), 0,
		0,
		0,
		0)

	return ret != 0
}

func GetBkColor(hdc HDC) COLORREF {
	ret, _, _ := syscall.Syscall(getBkColor.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return COLORREF(ret)
}

func GetDeviceCaps(hdc HDC, nIndex int32) int32 {
	ret, _, _ := syscall.Syscall(getDeviceCaps.Addr(), 2,
		uintptr(hdc),
		uintptr(nIndex),
		0)

	return int32(ret)
}

func GetDIBits(hdc HDC, hbmp HBITMAP, uStartScan uint32, cScanLines uint32, lpvBits *byte, lpbi *BITMAPINFO, uUsage uint32) int32 {
	ret, _, _ := syscall.Syscall9(getDIBits.Addr(), 7,
		uintptr(hdc),
		uintptr(hbmp),
		uintptr(uStartScan),
		uintptr(cScanLines),
		uintptr(unsafe.Pointer(lpvBits)),
		uintptr(unsafe.Pointer(lpbi)),
		uintptr(uUsage),
		0,
		0)
	return int32(ret)
}

func GetEnhMetaFile(lpszMetaFile *uint16) HENHMETAFILE {
	ret, _, _ := syscall.Syscall(getEnhMetaFile.Addr(), 1,
		uintptr(unsafe.Pointer(lpszMetaFile)),
		0,
		0)

	return HENHMETAFILE(ret)
}

func GetEnhMetaFileHeader(hemf HENHMETAFILE, cbBuffer uint32, lpemh *ENHMETAHEADER) uint32 {
	ret, _, _ := syscall.Syscall(getEnhMetaFileHeader.Addr(), 3,
		uintptr(hemf),
		uintptr(cbBuffer),
		uintptr(unsafe.Pointer(lpemh)))

	return uint32(ret)
}

func GetObject(hgdiobj HGDIOBJ, cbBuffer uintptr, lpvObject unsafe.Pointer) int32 {
	ret, _, _ := syscall.Syscall(getObject.Addr(), 3,
		uintptr(hgdiobj),
		uintptr(cbBuffer),
		uintptr(lpvObject))

	return int32(ret)
}

func GetPixel(hdc HDC, nXPos, nYPos int32) COLORREF {
	ret, _, _ := syscall.Syscall(getPixel.Addr(), 3,
		uintptr(hdc),
		uintptr(nXPos),
		uintptr(nYPos))

	return COLORREF(ret)
}

func GetRgnBox(hrgn HRGN, lprc *RECT) int32 {
	ret, _, _ := syscall.Syscall(getRgnBox.Addr(), 2,
		uintptr(hrgn),
		uintptr(unsafe.Pointer(lprc)),
		0)

	return int32(ret)
}

func GetStockObject(fnObject int32) HGDIOBJ {
	ret, _, _ := syscall.Syscall(getStockObject.Addr(), 1,
		uintptr(fnObject),
		0,
		0)

	return HGDIOBJ(ret)
}

func GetTextColor(hdc HDC) COLORREF {
	ret, _, _ := syscall.Syscall(getTextColor.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return COLORREF(ret)
}

func GetTextExtentExPoint(hdc HDC, lpszStr *uint16, cchString, nMaxExtent int32, lpnFit, alpDx *int32, lpSize *SIZE) bool {
	ret, _, _ := syscall.Syscall9(getTextExtentExPoint.Addr(), 7,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpszStr)),
		uintptr(cchString),
		uintptr(nMaxExtent),
		uintptr(unsafe.Pointer(lpnFit)),
		uintptr(unsafe.Pointer(alpDx)),
		uintptr(unsafe.Pointer(lpSize)),
		0,
		0)

	return ret != 0
}

func GetTextExtentPoint32(hdc HDC, lpString *uint16, c int32, lpSize *SIZE) bool {
	ret, _, _ := syscall.Syscall6(getTextExtentPoint32.Addr(), 4,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(c),
		uintptr(unsafe.Pointer(lpSize)),
		0,
		0)

	return ret != 0
}

func GetTextMetrics(hdc HDC, lptm *TEXTMETRIC) bool {
	ret, _, _ := syscall.Syscall(getTextMetrics.Addr(), 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lptm)),
		0)

	return ret != 0
}

func GetViewportOrgEx(hdc HDC, lpPoint *POINT) bool {
	ret, _, _ := syscall.Syscall(getViewportOrgEx.Addr(), 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpPoint)),
		0)

	return ret != 0
}

func GradientFill(hdc HDC, pVertex *TRIVERTEX, nVertex uint32, pMesh unsafe.Pointer, nMesh, ulMode uint32) bool {
	ret, _, _ := syscall.Syscall6(gradientFill.Addr(), 6,
		uintptr(hdc),
		uintptr(unsafe.Pointer(pVertex)),
		uintptr(nVertex),
		uintptr(pMesh),
		uintptr(nMesh),
		uintptr(ulMode))

	return ret != 0
}

func IntersectClipRect(hdc HDC, nLeftRect, nTopRect, nRightRect, nBottomRect int32) int32 {
	ret, _, _ := syscall.Syscall6(intersectClipRect.Addr(), 5,
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		0)

	return int32(ret)
}

func LineTo(hdc HDC, nXEnd, nYEnd int32) bool {
	ret, _, _ := syscall.Syscall(lineTo.Addr(), 3,
		uintptr(hdc),
		uintptr(nXEnd),
		uintptr(nYEnd))

	return ret != 0
}

func MoveToEx(hdc HDC, x, y int, lpPoint *POINT) bool {
	ret, _, _ := syscall.Syscall6(moveToEx.Addr(), 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lpPoint)),
		0,
		0)

	return ret != 0
}

func PlayEnhMetaFile(hdc HDC, hemf HENHMETAFILE, lpRect *RECT) bool {
	ret, _, _ := syscall.Syscall(playEnhMetaFile.Addr(), 3,
		uintptr(hdc),
		uintptr(hemf),
		uintptr(unsafe.Pointer(lpRect)))

	return ret != 0
}

func Polyline(hdc HDC, lppt unsafe.Pointer, cPoints int32) bool {
	ret, _, _ := syscall.Syscall(polyline.Addr(), 3,
		uintptr(hdc),
		uintptr(lppt),
		uintptr(cPoints))

	return ret != 0
}

func Rectangle_(hdc HDC, nLeftRect, nTopRect, nRightRect, nBottomRect int32) bool {
	ret, _, _ := syscall.Syscall6(rectangle.Addr(), 5,
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		0)

	return ret != 0
}

func RemoveFontResourceEx(lpszFilename *uint16, fl uint32, pdv unsafe.Pointer) bool {
	ret, _, _ := syscall.Syscall(removeFontResourceEx.Addr(), 3,
		uintptr(unsafe.Pointer(lpszFilename)),
		uintptr(fl),
		uintptr(pdv))

	return ret != 0
}

func RemoveFontMemResourceEx(h handle.HANDLE) bool {
	ret, _, _ := syscall.Syscall(removeFontMemResourceEx.Addr(), 1,
		uintptr(h),
		0,
		0)

	return ret != 0
}

func ResetDC(hdc HDC, lpInitData *DEVMODE) HDC {
	ret, _, _ := syscall.Syscall(resetDC.Addr(), 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpInitData)),
		0)

	return HDC(ret)
}

func RestoreDC(hdc HDC, nSaveDC int32) bool {
	ret, _, _ := syscall.Syscall(restoreDC.Addr(), 2,
		uintptr(hdc),
		uintptr(nSaveDC),
		0)
	return ret != 0
}

func RoundRect(hdc HDC, nLeftRect, nTopRect, nRightRect, nBottomRect, nWidth, nHeight int32) bool {
	ret, _, _ := syscall.Syscall9(roundRect.Addr(), 7,
		uintptr(hdc),
		uintptr(nLeftRect),
		uintptr(nTopRect),
		uintptr(nRightRect),
		uintptr(nBottomRect),
		uintptr(nWidth),
		uintptr(nHeight),
		0,
		0)

	return ret != 0
}

func SaveDC(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(saveDC.Addr(), 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret)
}

func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
	ret, _, _ := syscall.Syscall(selectObject.Addr(), 2,
		uintptr(hdc),
		uintptr(hgdiobj),
		0)

	return HGDIOBJ(ret)
}

func SetBkColor(hdc HDC, crColor COLORREF) COLORREF {
	ret, _, _ := syscall.Syscall(setBkColor.Addr(), 2,
		uintptr(hdc),
		uintptr(crColor),
		0)

	return COLORREF(ret)
}

func SetBkMode(hdc HDC, iBkMode int32) int32 {
	ret, _, _ := syscall.Syscall(setBkMode.Addr(), 2,
		uintptr(hdc),
		uintptr(iBkMode),
		0)

	return int32(ret)
}

func SetBrushOrgEx(hdc HDC, nXOrg, nYOrg int32, lppt *POINT) bool {
	ret, _, _ := syscall.Syscall6(setBrushOrgEx.Addr(), 4,
		uintptr(hdc),
		uintptr(nXOrg),
		uintptr(nYOrg),
		uintptr(unsafe.Pointer(lppt)),
		0,
		0)

	return ret != 0
}

func SetDIBits(hdc HDC, hbmp HBITMAP, uStartScan, cScanLines uint32, lpvBits *byte, lpbmi *BITMAPINFO, fuColorUse uint32) int32 {
	ret, _, _ := syscall.Syscall9(setDIBits.Addr(), 7,
		uintptr(hdc),
		uintptr(hbmp),
		uintptr(uStartScan),
		uintptr(cScanLines),
		uintptr(unsafe.Pointer(lpvBits)),
		uintptr(unsafe.Pointer(lpbmi)),
		uintptr(fuColorUse),
		0,
		0)

	return int32(ret)
}

func SetPixel(hdc HDC, X, Y int32, crColor COLORREF) COLORREF {
	ret, _, _ := syscall.Syscall6(setPixel.Addr(), 4,
		uintptr(hdc),
		uintptr(X),
		uintptr(Y),
		uintptr(crColor),
		0,
		0)

	return COLORREF(ret)
}

func SetPixelFormat(hdc HDC, iPixelFormat int32, ppfd *PIXELFORMATDESCRIPTOR) bool {
	ret, _, _ := syscall.Syscall(setPixelFormat.Addr(), 3,
		uintptr(hdc),
		uintptr(iPixelFormat),
		uintptr(unsafe.Pointer(ppfd)))

	return ret != 0
}

func SetStretchBltMode(hdc HDC, iStretchMode int32) int32 {
	ret, _, _ := syscall.Syscall(setStretchBltMode.Addr(), 2,
		uintptr(hdc),
		uintptr(iStretchMode),
		0)

	return int32(ret)
}

func SetTextColor(hdc HDC, crColor COLORREF) COLORREF {
	ret, _, _ := syscall.Syscall(setTextColor.Addr(), 2,
		uintptr(hdc),
		uintptr(crColor),
		0)

	return COLORREF(ret)
}

func SetViewportOrgEx(hdc HDC, x, y int32, lpPoint *POINT) COLORREF {
	ret, _, _ := syscall.Syscall6(setViewportOrgEx.Addr(), 4,
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(lpPoint)),
		0,
		0)

	return COLORREF(ret)
}

func StartDoc(hdc HDC, lpdi *DOCINFO) int32 {
	ret, _, _ := syscall.Syscall(startDoc.Addr(), 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(lpdi)),
		0)

	return int32(ret)
}

func StartPage(hdc HDC) int32 {
	ret, _, _ := syscall.Syscall(startPage.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return int32(ret)
}

func StretchBlt(hdcDest HDC, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int32, hdcSrc HDC, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc int32, dwRop uint32) bool {
	ret, _, _ := syscall.Syscall12(stretchBlt.Addr(), 11,
		uintptr(hdcDest),
		uintptr(nXOriginDest),
		uintptr(nYOriginDest),
		uintptr(nWidthDest),
		uintptr(nHeightDest),
		uintptr(hdcSrc),
		uintptr(nXOriginSrc),
		uintptr(nYOriginSrc),
		uintptr(nWidthSrc),
		uintptr(nHeightSrc),
		uintptr(dwRop),
		0)

	return ret != 0
}

func SwapBuffers(hdc HDC) bool {
	ret, _, _ := syscall.Syscall(swapBuffers.Addr(), 1,
		uintptr(hdc),
		0,
		0)

	return ret != 0
}

func TextOut(hdc HDC, nXStart, nYStart int32, lpString *uint16, cchString int32) bool {
	ret, _, _ := syscall.Syscall6(textOut.Addr(), 5,
		uintptr(hdc),
		uintptr(nXStart),
		uintptr(nYStart),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(cchString),
		0)
	return ret != 0
}

func TransparentBlt(hdcDest HDC, xoriginDest, yoriginDest, wDest, hDest int32, hdcSrc HDC, xoriginSrc, yoriginSrc, wSrc, hSrc int32, crTransparent uint32) bool {
	ret, _, _ := syscall.Syscall12(transparentBlt.Addr(), 11,
		uintptr(hdcDest),
		uintptr(xoriginDest),
		uintptr(yoriginDest),
		uintptr(wDest),
		uintptr(hDest),
		uintptr(hdcSrc),
		uintptr(xoriginSrc),
		uintptr(yoriginSrc),
		uintptr(wSrc),
		uintptr(hSrc),
		uintptr(crTransparent),
		0)

	return ret != 0
}
