// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package richedit

import (
	"github.com/Gipcomp/winapi/gdi32"
	"github.com/Gipcomp/winapi/handle"
	"github.com/Gipcomp/winapi/kernel32"
	"github.com/Gipcomp/winapi/user32"
	"github.com/Gipcomp/winapi/win"
)

type TEXTMODE int32

// EM_GETIMECOMPTEXT wparam structure
type IMECOMPTEXT struct {
	// count of bytes in the output buffer.
	Cb int32

	// value specifying the composition string type.
	//	Currently only support ICT_RESULTREADSTR
	Flags uint32
}

// Data type defining table rows for EM_INSERTTABLE
// Note: The Richedit.h is completely #pragma pack(4)-ed
type TABLEROWPARMS struct { // EM_INSERTTABLE wparam is a (TABLEROWPARMS *)
	CbRow        uint32 // Count of bytes in this structure
	CbCell       uint32 // Count of bytes in TABLECELLPARMS
	CCell        uint32 // Count of cells
	CRow         uint32 // Count of rows
	DxCellMargin int32  // Cell left/right margin (\trgaph)
	DxIndent     int32  // Row left (right if fRTL indent (similar to \trleft)
	DyHeight     int32  // Row height (\trrh)

	// nAlignment:3   Row alignment (like PARAFORMAT::bAlignment, \trql, trqr, \trqc)
	// fRTL:1         Display cells in RTL order (\rtlrow)
	// fKeep:1        Keep row together (\trkeep}
	// fKeepFollow:1  Keep row on same page as following row (\trkeepfollow)
	// fWrap:1        Wrap text to right/left (depending on bAlignment) (see \tdfrmtxtLeftN, \tdfrmtxtRightN)
	// fIdentCells:1  lparam points at single struct valid for all cells
	Flags uint32

	CpStartRow  int32  // cp where to insert table (-1 for selection cp) (can be used for either TRD by EM_GETTABLEPARMS)
	BTableLevel uint32 // Table nesting level (EM_GETTABLEPARMS only)
	ICell       uint32 // Index of cell to insert/delete (EM_SETTABLEPARMS only)
}

// Data type defining table cells for EM_INSERTTABLE
// Note: The Richedit.h is completely #pragma pack(4)-ed
type TABLECELLPARMS struct { // EM_INSERTTABLE lparam is a (TABLECELLPARMS *)
	DxWidth int32 // Cell width (\cellx)

	// nVertAlign:2   Vertical alignment (0/1/2 = top/center/bottom \clvertalt (def), \clvertalc, \clvertalb)
	// fMergeTop:1    Top cell for vertical merge (\clvmgf)
	// fMergePrev:1   Merge with cell above (\clvmrg)
	// fVertical:1    Display text top to bottom, right to left (\cltxtbrlv)
	// fMergeStart:1  Start set of horizontally merged cells (\clmgf)
	// fMergeCont:1   Merge with previous cell (\clmrg)
	Flags uint32

	WShading uint32 // Shading in .01%		(\clshdng) e.g., 10000 flips fore/back

	DxBrdrLeft   int32 // Left border width	(\clbrdrl\brdrwN) (in twips)
	DyBrdrTop    int32 // Top border width 	(\clbrdrt\brdrwN)
	DxBrdrRight  int32 // Right border width	(\clbrdrr\brdrwN)
	DyBrdrBottom int32 // Bottom border width	(\clbrdrb\brdrwN)

	CrBrdrLeft   gdi32.COLORREF // Left border color	(\clbrdrl\brdrcf)
	CrBrdrTop    gdi32.COLORREF // Top border color 	(\clbrdrt\brdrcf)
	CrBrdrRight  gdi32.COLORREF // Right border color	(\clbrdrr\brdrcf)
	CrBrdrBottom gdi32.COLORREF // Bottom border color	(\clbrdrb\brdrcf)
	CrBackPat    gdi32.COLORREF // Background color 	(\clcbpat)
	CrForePat    gdi32.COLORREF // Foreground color 	(\clcfpat)
}

// AutoCorrect callback
type AutoCorrectProc func(langid kernel32.LANGID, pszBefore *uint16, pszAfter *uint16, cchAfter int32, pcchReplaced *int32) int

// lparam for EM_INSERTIMAGE
type RICHEDIT_IMAGE_PARAMETERS struct {
	XWidth            int32 // Units are HIMETRIC
	YHeight           int32 // Units are HIMETRIC
	Ascent            int32 // Units are HIMETRIC
	Type              int32 // Valid values are TA_TOP, TA_BOTTOM and TA_BASELINE
	PwszAlternateText *uint16
	PIStream          uintptr
}

// Notification structure for EN_ENDCOMPOSITION
type ENDCOMPOSITIONNOTIFY struct {
	Nmhdr  user32.NMHDR
	DwCode uint32
}

type CHARFORMAT struct {
	CbSize          uint32
	DwMask          uint32
	DwEffects       uint32
	YHeight         int32
	YOffset         int32
	CrTextColor     gdi32.COLORREF
	BCharSet        byte
	BPitchAndFamily byte
	SzFaceName      [gdi32.LF_FACESIZE]uint16
}

type CHARFORMAT2 struct {
	CHARFORMAT
	WWeight         uint16         // Font weight (LOGFONT value)
	SSpacing        int16          // Amount to space between letters
	CrBackColor     gdi32.COLORREF // Background color
	Lcid            kernel32.LCID  // Locale ID
	DwCookie        uint32         // Client cookie opaque to RichEdit
	SStyle          int16          // Style handle
	WKerning        uint16         // Twip size above which to kern char pair
	BUnderlineType  byte           // Underline type
	BAnimation      byte           // Animated text like marching ants
	BRevAuthor      byte           // Revision author index
	BUnderlineColor byte           // Underline color
}

type CHARRANGE struct {
	CpMin int32
	CpMax int32
}

type TEXTRANGE struct {
	Chrg      CHARRANGE
	LpstrText *uint16 // Allocated by caller, zero terminated by RichEdit
}

type EDITSTREAM struct {
	DwCookie    uintptr // User value passed to callback as first parameter
	DwError     uint32  // Last error
	PfnCallback uintptr
}

type FINDTEXT struct {
	Chrg      CHARRANGE
	LpstrText *uint16
}

type FINDTEXTEX struct {
	chrg      CHARRANGE
	lpstrText *uint16
	chrgText  CHARRANGE
}

type FORMATRANGE struct {
	hdc       gdi32.HDC
	hdcTarget gdi32.HDC
	rc        gdi32.RECT
	rcPage    gdi32.RECT
	chrg      CHARRANGE
}

type PARAFORMAT struct {
	CbSize        uint32
	DwMask        uint32
	WNumbering    uint16
	WEffects      uint16
	DxStartIndent int32
	DxRightIndent int32
	DxOffset      int32
	WAlignment    uint16
	CTabCount     int16
	RgxTabs       [MAX_TAB_STOPS]int32
}

type PARAFORMAT2 struct {
	PARAFORMAT
	DySpaceBefore    int32  // Vertical spacing before para
	DySpaceAfter     int32  // Vertical spacing after para
	DyLineSpacing    int32  // Line spacing depending on Rule
	SStyle           int16  // Style handle
	BLineSpacingRule byte   // Rule for line spacing (see tom.doc)
	BOutlineLevel    byte   // Outline level
	WShadingWeight   uint16 // Shading in hundredths of a per cent
	WShadingStyle    uint16 // Nibble 0: style, 1: cfpat, 2: cbpat
	WNumberingStart  uint16 // Starting value for numbering
	WNumberingStyle  uint16 // Alignment, roman/arabic, (), ), ., etc.
	WNumberingTab    uint16 // Space bet FirstIndent & 1st-line text
	WBorderSpace     uint16 // Border-text spaces (nbl/bdr in pts)
	WBorderWidth     uint16 // Pen widths (nbl/bdr in half pts)
	WBorders         uint16 // Border styles (nibble/border)
}

type MSGFILTER struct {
	Nmhdr  user32.NMHDR
	Msg    uint32
	WParam uintptr
	LParam uintptr
}

type REQRESIZE struct {
	Nmhdr user32.NMHDR
	Rc    gdi32.RECT
}

type SELCHANGE struct {
	Nmhdr  user32.NMHDR
	Chrg   CHARRANGE
	Seltyp uint16
}

type GROUPTYPINGCHANGE struct {
	Nmhdr        user32.NMHDR
	FGroupTyping win.BOOL
}

type CLIPBOARDFORMAT struct {
	Nmhdr user32.NMHDR
	Cf    gdi32.CLIPFORMAT
}

type GETCONTEXTMENUEX struct {
	Chrg       CHARRANGE
	DwFlags    uint32
	Pt         gdi32.POINT
	PvReserved uintptr
}

type ENDROPFILES struct {
	Nmhdr      user32.NMHDR
	HDrop      handle.HANDLE
	Cp         int32
	FProtected win.BOOL
}

type ENPROTECTED struct {
	Nmhdr  user32.NMHDR
	Msg    uint32
	WParam uintptr
	LParam uintptr
	Chrg   CHARRANGE
}

type ENSAVECLIPBOARD struct {
	Nmhdr        user32.NMHDR
	CObjectCount int32
	Cch          int32
}

type ENOLEOPFAILED struct {
	Nmhdr user32.NMHDR
	Iob   int32
	LOper int32
	Hr    win.HRESULT
}

type OBJECTPOSITIONS struct {
	Nmhdr        user32.NMHDR
	CObjectCount int32
	PcpPositions *int32
}

type ENLINK struct {
	Nmhdr  user32.NMHDR
	Msg    uint32
	WParam uintptr
	LParam uintptr
	Chrg   CHARRANGE
}

type ENLOWFIRTF struct {
	Nmhdr     user32.NMHDR
	SzControl *byte
}

// PenWin specific
type ENCORRECTTEXT struct {
	Nmhdr  user32.NMHDR
	Chrg   CHARRANGE
	Seltyp uint16
}

// East Asia specific
type PUNCTUATION struct {
	ISize         uint32
	SzPunctuation *byte
}

// East Asia specific
type COMPCOLOR struct {
	CrText       gdi32.COLORREF
	CrBackground gdi32.COLORREF
	DwEffects    uint32
}

// Paste Special
type REPASTESPECIAL struct {
	DwAspect uint32
	DwParam  uintptr
}

//	UndoName info
type UNDONAMEID int32

// EM_SETTEXTEX info; this struct is passed in the wparam of the message
type SETTEXTEX struct {
	Flags    uint32 // Flags (see the ST_XXX defines)
	Codepage uint32 // Code page for translation (CP_ACP for sys default, 1200 for Unicode, -1 for control default)
}

// EM_GETTEXTEX info; this struct is passed in the wparam of the message
type GETTEXTEX struct {
	Cb            uint32    // Count of bytes in the string
	Flags         uint32    // Flags (see the GT_XXX defines
	Codepage      uint32    // Code page for translation (CP_ACP for sys default, 1200 for Unicode, -1 for control default)
	LpDefaultChar *byte     // Replacement for unmappable chars
	LpUsedDefChar *win.BOOL // Pointer to flag set when def char used
}

// EM_GETTEXTLENGTHEX info; this struct is passed in the wparam of the msg
type GETTEXTLENGTHEX struct {
	Flags    uint32 // Flags (see GTL_XXX defines)
	Codepage uint32 // Code page for translation (CP_ACP for default, 1200 for Unicode)
}

// BiDi specific features
type BIDIOPTIONS struct {
	CbSize   uint32
	WMask    uint16
	WEffects uint16
}

// khyph - Kind of hyphenation
type KHYPH int32

type HYPHRESULT struct {
	Khyph   KHYPH  // Kind of hyphenation
	IchHyph int32  // Character which was hyphenated
	ChHyph  uint16 // Depending on hyphenation type, character added, changed, etc.
}

type HYPHENATEINFO struct {
	CbSize          int16 // Size of HYPHENATEINFO structure
	DxHyphenateZone int16 // If a space character is closer to the margin than this value, don't hyphenate (in TWIPs)
	PfnHyphenate    uintptr
}
