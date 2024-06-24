// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package richedit

import "github.com/gfphoenix/win32/user32"

const (
	// NOTE:  MSFTEDIT.DLL only registers MSFTEDIT_CLASS.  If an application wants
	// to use the following RichEdit classes, it needs to load riched20.dll.
	// Otherwise, CreateWindow with RICHEDIT_CLASS will fail.
	// This also applies to any dialog that uses RICHEDIT_CLASS
	// RichEdit 2.0 Window Class
	MSFTEDIT_CLASS = "RICHEDIT50W"
	RICHEDIT_CLASS = "RichEdit20W"
)

// RichEdit messages
const (
	EM_CANPASTE           = user32.WM_USER + 50
	EM_DISPLAYBAND        = user32.WM_USER + 51
	EM_EXGETSEL           = user32.WM_USER + 52
	EM_EXLIMITTEXT        = user32.WM_USER + 53
	EM_EXLINEFROMCHAR     = user32.WM_USER + 54
	EM_EXSETSEL           = user32.WM_USER + 55
	EM_FINDTEXT           = user32.WM_USER + 56
	EM_FORMATRANGE        = user32.WM_USER + 57
	EM_GETCHARFORMAT      = user32.WM_USER + 58
	EM_GETEVENTMASK       = user32.WM_USER + 59
	EM_GETOLEINTERFACE    = user32.WM_USER + 60
	EM_GETPARAFORMAT      = user32.WM_USER + 61
	EM_GETSELTEXT         = user32.WM_USER + 62
	EM_HIDESELECTION      = user32.WM_USER + 63
	EM_PASTESPECIAL       = user32.WM_USER + 64
	EM_REQUESTRESIZE      = user32.WM_USER + 65
	EM_SELECTIONTYPE      = user32.WM_USER + 66
	EM_SETBKGNDCOLOR      = user32.WM_USER + 67
	EM_SETCHARFORMAT      = user32.WM_USER + 68
	EM_SETEVENTMASK       = user32.WM_USER + 69
	EM_SETOLECALLBACK     = user32.WM_USER + 70
	EM_SETPARAFORMAT      = user32.WM_USER + 71
	EM_SETTARGETDEVICE    = user32.WM_USER + 72
	EM_STREAMIN           = user32.WM_USER + 73
	EM_STREAMOUT          = user32.WM_USER + 74
	EM_GETTEXTRANGE       = user32.WM_USER + 75
	EM_FINDWORDBREAK      = user32.WM_USER + 76
	EM_SETOPTIONS         = user32.WM_USER + 77
	EM_GETOPTIONS         = user32.WM_USER + 78
	EM_FINDTEXTEX         = user32.WM_USER + 79
	EM_GETWORDBREAKPROCEX = user32.WM_USER + 80
	EM_SETWORDBREAKPROCEX = user32.WM_USER + 81
)

// RichEdit 2.0 messages
const (
	EM_SETUNDOLIMIT    = user32.WM_USER + 82
	EM_REDO            = user32.WM_USER + 84
	EM_CANREDO         = user32.WM_USER + 85
	EM_GETUNDONAME     = user32.WM_USER + 86
	EM_GETREDONAME     = user32.WM_USER + 87
	EM_STOPGROUPTYPING = user32.WM_USER + 88

	EM_SETTEXTMODE = user32.WM_USER + 89
	EM_GETTEXTMODE = user32.WM_USER + 90
)

const (
	TM_PLAINTEXT       TEXTMODE = 1
	TM_RICHTEXT                 = 2 // Default behavior
	TM_SINGLELEVELUNDO          = 4
	TM_MULTILEVELUNDO           = 8 // Default behavior
	TM_SINGLECODEPAGE           = 16
	TM_MULTICODEPAGE            = 32 // Default behavior
)

const (
	EM_AUTOURLDETECT = user32.WM_USER + 91
)

// RichEdit 8.0 messages
const (
	AURL_ENABLEURL          = 1
	AURL_ENABLEEMAILADDR    = 2
	AURL_ENABLETELNO        = 4
	AURL_ENABLEEAURLS       = 8
	AURL_ENABLEDRIVELETTERS = 16
	AURL_DISABLEMIXEDLGC    = 32 // Disable mixed Latin Greek Cyrillic IDNs
)

const (
	EM_GETAUTOURLDETECT = user32.WM_USER + 92
	EM_SETPALETTE       = user32.WM_USER + 93
	EM_GETTEXTEX        = user32.WM_USER + 94
	EM_GETTEXTLENGTHEX  = user32.WM_USER + 95
	EM_SHOWSCROLLBAR    = user32.WM_USER + 96
	EM_SETTEXTEX        = user32.WM_USER + 97
)

// East Asia specific messages
const (
	EM_SETPUNCTUATION  = user32.WM_USER + 100
	EM_GETPUNCTUATION  = user32.WM_USER + 101
	EM_SETWORDWRAPMODE = user32.WM_USER + 102
	EM_GETWORDWRAPMODE = user32.WM_USER + 103
	EM_SETIMECOLOR     = user32.WM_USER + 104
	EM_GETIMECOLOR     = user32.WM_USER + 105
	EM_SETIMEOPTIONS   = user32.WM_USER + 106
	EM_GETIMEOPTIONS   = user32.WM_USER + 107
	EM_CONVPOSITION    = user32.WM_USER + 108
)

const (
	EM_SETLANGOPTIONS = user32.WM_USER + 120
	EM_GETLANGOPTIONS = user32.WM_USER + 121
	EM_GETIMECOMPMODE = user32.WM_USER + 122

	EM_FINDTEXTW   = user32.WM_USER + 123
	EM_FINDTEXTEXW = user32.WM_USER + 124
)

// RE3.0 FE messages
const (
	EM_RECONVERSION   = user32.WM_USER + 125
	EM_SETIMEMODEBIAS = user32.WM_USER + 126
	EM_GETIMEMODEBIAS = user32.WM_USER + 127
)

// BiDi specific messages
const (
	EM_SETBIDIOPTIONS = user32.WM_USER + 200
	EM_GETBIDIOPTIONS = user32.WM_USER + 201

	EM_SETTYPOGRAPHYOPTIONS = user32.WM_USER + 202
	EM_GETTYPOGRAPHYOPTIONS = user32.WM_USER + 203
)

// Extended edit style specific messages
const (
	EM_SETEDITSTYLE = user32.WM_USER + 204
	EM_GETEDITSTYLE = user32.WM_USER + 205
)

// Extended edit style masks
const (
	SES_EMULATESYSEDIT    = 1
	SES_BEEPONMAXTEXT     = 2
	SES_EXTENDBACKCOLOR   = 4
	SES_MAPCPS            = 8 // Obsolete (never used)
	SES_HYPERLINKTOOLTIPS = 8
	SES_EMULATE10         = 16 // Obsolete (never used)
	SES_DEFAULTLATINLIGA  = 16
	SES_USECRLF           = 32 // Obsolete (never used)
	SES_NOFOCUSLINKNOTIFY = 32
	SES_USEAIMM           = 64
	SES_NOIME             = 128

	SES_ALLOWBEEPS         = 256
	SES_UPPERCASE          = 512
	SES_LOWERCASE          = 1024
	SES_NOINPUTSEQUENCECHK = 2048
	SES_BIDI               = 4096
	SES_SCROLLONKILLFOCUS  = 8192
	SES_XLTCRCRLFTOCR      = 16384
	SES_DRAFTMODE          = 32768

	SES_USECTF               = 0x00010000
	SES_HIDEGRIDLINES        = 0x00020000
	SES_USEATFONT            = 0x00040000
	SES_CUSTOMLOOK           = 0x00080000
	SES_LBSCROLLNOTIFY       = 0x00100000
	SES_CTFALLOWEMBED        = 0x00200000
	SES_CTFALLOWSMARTTAG     = 0x00400000
	SES_CTFALLOWPROOFING     = 0x00800000
	SES_LOGICALCARET         = 0x01000000
	SES_WORDDRAGDROP         = 0x02000000
	SES_SMARTDRAGDROP        = 0x04000000
	SES_MULTISELECT          = 0x08000000
	SES_CTFNOLOCK            = 0x10000000
	SES_NOEALINEHEIGHTADJUST = 0x20000000
	SES_MAX                  = 0x20000000
)

// Options for EM_SETLANGOPTIONS and EM_GETLANGOPTIONS
const (
	IMF_AUTOKEYBOARD        = 0x0001
	IMF_AUTOFONT            = 0x0002
	IMF_IMECANCELCOMPLETE   = 0x0004 // High completes comp string when aborting, low cancels
	IMF_IMEALWAYSSENDNOTIFY = 0x0008
	IMF_AUTOFONTSIZEADJUST  = 0x0010
	IMF_UIFONTS             = 0x0020
	IMF_NOIMPLICITLANG      = 0x0040
	IMF_DUALFONT            = 0x0080
	IMF_NOKBDLIDFIXUP       = 0x0200
	IMF_NORTFFONTSUBSTITUTE = 0x0400
	IMF_SPELLCHECKING       = 0x0800
	IMF_TKBPREDICTION       = 0x1000
	IMF_IMEUIINTEGRATION    = 0x2000
)

// Values for EM_GETIMECOMPMODE
const (
	ICM_NOTOPEN    = 0x0000
	ICM_LEVEL3     = 0x0001
	ICM_LEVEL2     = 0x0002
	ICM_LEVEL2_5   = 0x0003
	ICM_LEVEL2_SUI = 0x0004
	ICM_CTF        = 0x0005
)

// Options for EM_SETTYPOGRAPHYOPTIONS
const (
	TO_ADVANCEDTYPOGRAPHY   = 0x0001
	TO_SIMPLELINEBREAK      = 0x0002
	TO_DISABLECUSTOMTEXTOUT = 0x0004
	TO_ADVANCEDLAYOUT       = 0x0008
)

// Pegasus outline mode messages (RE 3.0)
const (
	// Outline mode message
	EM_OUTLINE = user32.WM_USER + 220

	// Message for getting and restoring scroll pos
	EM_GETSCROLLPOS = user32.WM_USER + 221
	EM_SETSCROLLPOS = user32.WM_USER + 222

	// Change fontsize in current selection by wParam
	EM_SETFONTSIZE = user32.WM_USER + 223
	EM_GETZOOM     = user32.WM_USER + 224
	EM_SETZOOM     = user32.WM_USER + 225
	EM_GETVIEWKIND = user32.WM_USER + 226
	EM_SETVIEWKIND = user32.WM_USER + 227
)

// RichEdit 4.0 messages
const (
	EM_GETPAGE          = user32.WM_USER + 228
	EM_SETPAGE          = user32.WM_USER + 229
	EM_GETHYPHENATEINFO = user32.WM_USER + 230
	EM_SETHYPHENATEINFO = user32.WM_USER + 231

	EM_GETPAGEROTATE    = user32.WM_USER + 235
	EM_SETPAGEROTATE    = user32.WM_USER + 236
	EM_GETCTFMODEBIAS   = user32.WM_USER + 237
	EM_SETCTFMODEBIAS   = user32.WM_USER + 238
	EM_GETCTFOPENSTATUS = user32.WM_USER + 240
	EM_SETCTFOPENSTATUS = user32.WM_USER + 241
	EM_GETIMECOMPTEXT   = user32.WM_USER + 242
	EM_ISIME            = user32.WM_USER + 243
	EM_GETIMEPROPERTY   = user32.WM_USER + 244
)

// These messages control what rich edit does when it comes accross
// OLE objects during RTF stream in.  Normally rich edit queries the client
// application only after OleLoad has been called.  With these messages it is possible to
// set the rich edit control to a mode where it will query the client application before
// OleLoad is called
const (
	EM_GETQUERYRTFOBJ = user32.WM_USER + 269
	EM_SETQUERYRTFOBJ = user32.WM_USER + 270
)

// EM_SETPAGEROTATE wparam values
const (
	EPR_0   = 0 // Text flows left to right and top to bottom
	EPR_270 = 1 // Text flows top to bottom and right to left
	EPR_180 = 2 // Text flows right to left and bottom to top
	EPR_90  = 3 // Text flows bottom to top and left to right
	EPR_SE  = 5 // Text flows top to bottom and left to right (Mongolian text layout)
)

// EM_SETCTFMODEBIAS wparam values
const (
	CTFMODEBIAS_DEFAULT               = 0x0000
	CTFMODEBIAS_FILENAME              = 0x0001
	CTFMODEBIAS_NAME                  = 0x0002
	CTFMODEBIAS_READING               = 0x0003
	CTFMODEBIAS_DATETIME              = 0x0004
	CTFMODEBIAS_CONVERSATION          = 0x0005
	CTFMODEBIAS_NUMERIC               = 0x0006
	CTFMODEBIAS_HIRAGANA              = 0x0007
	CTFMODEBIAS_KATAKANA              = 0x0008
	CTFMODEBIAS_HANGUL                = 0x0009
	CTFMODEBIAS_HALFWIDTHKATAKANA     = 0x000A
	CTFMODEBIAS_FULLWIDTHALPHANUMERIC = 0x000B
	CTFMODEBIAS_HALFWIDTHALPHANUMERIC = 0x000C
)

// EM_SETIMEMODEBIAS lparam values
const (
	IMF_SMODE_PLAURALCLAUSE = 0x0001
	IMF_SMODE_NONE          = 0x0002
)

const ICT_RESULTREADSTR = 1

// Outline mode wparam values
const (
	// Enter normal mode,  lparam ignored
	EMO_EXIT = 0

	// Enter outline mode, lparam ignored
	EMO_ENTER = 1

	// LOWORD(lparam) == 0 ==>
	//	promote  to body-text
	// LOWORD(lparam) != 0 ==>
	//	promote/demote current selection
	//	by indicated number of levels
	EMO_PROMOTE = 2

	// HIWORD(lparam) = EMO_EXPANDSELECTION
	//	-> expands selection to level
	//	indicated in LOWORD(lparam)
	//	LOWORD(lparam) = -1/+1 corresponds
	//	to collapse/expand button presses
	//	in winword (other values are
	//	equivalent to having pressed these
	//	buttons more than once)
	//	HIWORD(lparam) = EMO_EXPANDDOCUMENT
	//	-> expands whole document to
	//	indicated level
	EMO_EXPAND = 3

	// LOWORD(lparam) != 0 -> move current
	//	selection up/down by indicated amount
	EMO_MOVESELECTION = 4

	// Returns VM_NORMAL or VM_OUTLINE
	EMO_GETVIEWMODE = 5
)

// EMO_EXPAND options
const (
	EMO_EXPANDSELECTION = 0
	EMO_EXPANDDOCUMENT  = 1
)

const (
	// Agrees with RTF \viewkindN
	VM_NORMAL = 4

	VM_OUTLINE = 2

	// Screen page view (not print layout)
	VM_PAGE = 9
)

// New messages as of Win8
const (
	EM_INSERTTABLE = user32.WM_USER + 232
)

const (
	EM_GETAUTOCORRECTPROC  = user32.WM_USER + 233
	EM_SETAUTOCORRECTPROC  = user32.WM_USER + 234
	EM_CALLAUTOCORRECTPROC = user32.WM_USER + 255
)

const (
	ATP_NOCHANGE       = 0
	ATP_CHANGE         = 1
	ATP_NODELIMITER    = 2
	ATP_REPLACEALLTEXT = 4
)

const (
	EM_GETTABLEPARMS = user32.WM_USER + 265

	EM_SETEDITSTYLEEX = user32.WM_USER + 275
	EM_GETEDITSTYLEEX = user32.WM_USER + 276
)

// wparam values for EM_SETEDITSTYLEEX/EM_GETEDITSTYLEEX
// All unused bits are reserved.
const (
	SES_EX_NOTABLE            = 0x00000004
	SES_EX_NOMATH             = 0x00000040
	SES_EX_HANDLEFRIENDLYURL  = 0x00000100
	SES_EX_NOTHEMING          = 0x00080000
	SES_EX_NOACETATESELECTION = 0x00100000
	SES_EX_USESINGLELINE      = 0x00200000
	SES_EX_MULTITOUCH         = 0x08000000 // Only works under Win8+
	SES_EX_HIDETEMPFORMAT     = 0x10000000
	SES_EX_USEMOUSEWPARAM     = 0x20000000 // Use wParam when handling WM_MOUSEMOVE message and do not call GetAsyncKeyState
)

const (
	EM_GETSTORYTYPE = user32.WM_USER + 290
	EM_SETSTORYTYPE = user32.WM_USER + 291

	EM_GETELLIPSISMODE = user32.WM_USER + 305
	EM_SETELLIPSISMODE = user32.WM_USER + 306
)

// uint32: *lparam for EM_GETELLIPSISMODE, lparam for EM_SETELLIPSISMODE
const (
	ELLIPSIS_MASK = 0x00000003 // all meaningful bits
	ELLIPSIS_NONE = 0x00000000 // ellipsis disabled
	ELLIPSIS_END  = 0x00000001 // ellipsis at the end (forced break)
	ELLIPSIS_WORD = 0x00000003 // ellipsis at the end (word break)
)

const (
	EM_SETTABLEPARMS = user32.WM_USER + 307

	EM_GETTOUCHOPTIONS  = user32.WM_USER + 310
	EM_SETTOUCHOPTIONS  = user32.WM_USER + 311
	EM_INSERTIMAGE      = user32.WM_USER + 314
	EM_SETUIANAME       = user32.WM_USER + 320
	EM_GETELLIPSISSTATE = user32.WM_USER + 322
)

// Values for EM_SETTOUCHOPTIONS/EM_GETTOUCHOPTIONS
const (
	RTO_SHOWHANDLES    = 1
	RTO_DISABLEHANDLES = 2
	RTO_READINGMODE    = 3
)

// New notifications
const (
	EN_MSGFILTER         = 0x0700
	EN_REQUESTRESIZE     = 0x0701
	EN_SELCHANGE         = 0x0702
	EN_DROPFILES         = 0x0703
	EN_PROTECTED         = 0x0704
	EN_CORRECTTEXT       = 0x0705 // PenWin specific
	EN_STOPNOUNDO        = 0x0706
	EN_IMECHANGE         = 0x0707 // East Asia specific
	EN_SAVECLIPBOARD     = 0x0708
	EN_OLEOPFAILED       = 0x0709
	EN_OBJECTPOSITIONS   = 0x070a
	EN_LINK              = 0x070b
	EN_DRAGDROPDONE      = 0x070c
	EN_PARAGRAPHEXPANDED = 0x070d
	EN_PAGECHANGE        = 0x070e
	EN_LOWFIRTF          = 0x070f
	EN_ALIGNLTR          = 0x0710 // BiDi specific notification
	EN_ALIGNRTL          = 0x0711 // BiDi specific notification
	EN_CLIPFORMAT        = 0x0712
	EN_STARTCOMPOSITION  = 0x0713
	EN_ENDCOMPOSITION    = 0x0714
)

// Constants for ENDCOMPOSITIONNOTIFY dwCode
const (
	ECN_ENDCOMPOSITION = 0x0001
	ECN_NEWTEXT        = 0x0002
)

// Event notification masks
const (
	ENM_NONE              = 0x00000000
	ENM_CHANGE            = 0x00000001
	ENM_UPDATE            = 0x00000002
	ENM_SCROLL            = 0x00000004
	ENM_SCROLLEVENTS      = 0x00000008
	ENM_DRAGDROPDONE      = 0x00000010
	ENM_PARAGRAPHEXPANDED = 0x00000020
	ENM_PAGECHANGE        = 0x00000040
	ENM_CLIPFORMAT        = 0x00000080
	ENM_KEYEVENTS         = 0x00010000
	ENM_MOUSEEVENTS       = 0x00020000
	ENM_REQUESTRESIZE     = 0x00040000
	ENM_SELCHANGE         = 0x00080000
	ENM_DROPFILES         = 0x00100000
	ENM_PROTECTED         = 0x00200000
	ENM_CORRECTTEXT       = 0x00400000 // PenWin specific
	ENM_IMECHANGE         = 0x00800000 // Used by RE1.0 compatibility
	ENM_LANGCHANGE        = 0x01000000
	ENM_OBJECTPOSITIONS   = 0x02000000
	ENM_LINK              = 0x04000000
	ENM_LOWFIRTF          = 0x08000000
	ENM_STARTCOMPOSITION  = 0x10000000
	ENM_ENDCOMPOSITION    = 0x20000000
	ENM_GROUPTYPINGCHANGE = 0x40000000
	ENM_HIDELINKTOOLTIP   = 0x80000000
)

// New edit control styles
const (
	ES_SAVESEL         = 0x00008000
	ES_SUNKEN          = 0x00004000
	ES_DISABLENOSCROLL = 0x00002000
	ES_SELECTIONBAR    = 0x01000000 // Same as WS_MAXIMIZE, but that doesn't make sense so we re-use the value
	ES_NOOLEDRAGDROP   = 0x00000008 // Same as ES_UPPERCASE, but re-used to completely disable OLE drag'n'drop
)

// Obsolete Edit Style
const (
	ES_EX_NOCALLOLEINIT = 0x00000000 // Not supported in RE 2.0/3.0
)

// These flags are used in FE Windows
const (
	ES_VERTICAL = 0x00400000 // Not supported in RE 2.0/3.0
	ES_NOIME    = 0x00080000
	ES_SELFIME  = 0x00040000
)

// Edit control options
const (
	ECO_AUTOWORDSELECTION = 0x00000001
	ECO_AUTOVSCROLL       = 0x00000040
	ECO_AUTOHSCROLL       = 0x00000080
	ECO_NOHIDESEL         = 0x00000100
	ECO_READONLY          = 0x00000800
	ECO_WANTRETURN        = 0x00001000
	ECO_SAVESEL           = 0x00008000
	ECO_SELECTIONBAR      = 0x01000000
	ECO_VERTICAL          = 0x00400000 // FE specific
)

// ECO operations
const (
	ECOOP_SET = 0x0001
	ECOOP_OR  = 0x0002
	ECOOP_AND = 0x0003
	ECOOP_XOR = 0x0004
)

// New word break function actions
const (
	WB_CLASSIFY      = 3
	WB_MOVEWORDLEFT  = 4
	WB_MOVEWORDRIGHT = 5
	WB_LEFTBREAK     = 6
	WB_RIGHTBREAK    = 7
)

// East Asia specific flags
const (
	WB_MOVEWORDPREV = 4
	WB_MOVEWORDNEXT = 5
	WB_PREVBREAK    = 6
	WB_NEXTBREAK    = 7

	PC_FOLLOWING  = 1
	PC_LEADING    = 2
	PC_OVERFLOW   = 3
	PC_DELIMITER  = 4
	WBF_WORDWRAP  = 0x010
	WBF_WORDBREAK = 0x020
	WBF_OVERFLOW  = 0x040
	WBF_LEVEL1    = 0x080
	WBF_LEVEL2    = 0x100
	WBF_CUSTOM    = 0x200
)

// East Asia specific flags
const (
	IMF_FORCENONE         = 0x0001
	IMF_FORCEENABLE       = 0x0002
	IMF_FORCEDISABLE      = 0x0004
	IMF_CLOSESTATUSWINDOW = 0x0008
	IMF_VERTICAL          = 0x0020
	IMF_FORCEACTIVE       = 0x0040
	IMF_FORCEINACTIVE     = 0x0080
	IMF_FORCEREMEMBER     = 0x0100
	IMF_MULTIPLEEDIT      = 0x0400
)

// Word break flags (used with WB_CLASSIFY)
const (
	WBF_CLASS      byte = 0x0F
	WBF_ISWHITE    byte = 0x10
	WBF_BREAKLINE  byte = 0x20
	WBF_BREAKAFTER byte = 0x40
)

// CHARFORMAT masks
const (
	CFM_BOLD      = 0x00000001
	CFM_ITALIC    = 0x00000002
	CFM_UNDERLINE = 0x00000004
	CFM_STRIKEOUT = 0x00000008
	CFM_PROTECTED = 0x00000010
	CFM_LINK      = 0x00000020 // Exchange hyperlink extension
	CFM_SIZE      = 0x80000000
	CFM_COLOR     = 0x40000000
	CFM_FACE      = 0x20000000
	CFM_OFFSET    = 0x10000000
	CFM_CHARSET   = 0x08000000
)

// CHARFORMAT effects
const (
	CFE_BOLD      = 0x00000001
	CFE_ITALIC    = 0x00000002
	CFE_UNDERLINE = 0x00000004
	CFE_STRIKEOUT = 0x00000008
	CFE_PROTECTED = 0x00000010
	CFE_LINK      = 0x00000020
	CFE_AUTOCOLOR = 0x40000000 // NOTE: this corresponds to CFM_COLOR, which controls it

	// Masks and effects defined for CHARFORMAT2 -- an (*) indicates that the data is stored by RichEdit 2.0/3.0, but not displayed
	CFM_SMALLCAPS = 0x00000040 // (*)
	CFM_ALLCAPS   = 0x00000080 // Displayed by 3.0
	CFM_HIDDEN    = 0x00000100 // Hidden by 3.0
	CFM_OUTLINE   = 0x00000200 // (*)
	CFM_SHADOW    = 0x00000400 // (*)
	CFM_EMBOSS    = 0x00000800 // (*)
	CFM_IMPRINT   = 0x00001000 // (*)
	CFM_DISABLED  = 0x00002000
	CFM_REVISED   = 0x00004000

	CFM_REVAUTHOR     = 0x00008000
	CFE_SUBSCRIPT     = 0x00010000 // Superscript and subscript are
	CFE_SUPERSCRIPT   = 0x00020000 //	mutually exclusive
	CFM_ANIMATION     = 0x00040000 // (*)
	CFM_STYLE         = 0x00080000 // (*)
	CFM_KERNING       = 0x00100000
	CFM_SPACING       = 0x00200000 // Displayed by 3.0
	CFM_WEIGHT        = 0x00400000
	CFM_UNDERLINETYPE = 0x00800000 // Many displayed by 3.0
	CFM_COOKIE        = 0x01000000 // RE 6.0
	CFM_LCID          = 0x02000000
	CFM_BACKCOLOR     = 0x04000000 // Higher mask bits defined above

	CFM_SUBSCRIPT   = (CFE_SUBSCRIPT | CFE_SUPERSCRIPT)
	CFM_SUPERSCRIPT = CFM_SUBSCRIPT

	// CHARFORMAT "ALL" masks
	CFM_EFFECTS  = CFM_BOLD | CFM_ITALIC | CFM_UNDERLINE | CFM_COLOR | CFM_STRIKEOUT | CFE_PROTECTED | CFM_LINK
	CFM_ALL      = CFM_EFFECTS | CFM_SIZE | CFM_FACE | CFM_OFFSET | CFM_CHARSET
	CFM_EFFECTS2 = CFM_EFFECTS | CFM_DISABLED | CFM_SMALLCAPS | CFM_ALLCAPS | CFM_HIDDEN | CFM_OUTLINE | CFM_SHADOW | CFM_EMBOSS | CFM_IMPRINT | CFM_REVISED | CFM_SUBSCRIPT | CFM_SUPERSCRIPT | CFM_BACKCOLOR
	CFM_ALL2     = CFM_ALL | CFM_EFFECTS2 | CFM_BACKCOLOR | CFM_LCID | CFM_UNDERLINETYPE | CFM_WEIGHT | CFM_REVAUTHOR | CFM_SPACING | CFM_KERNING | CFM_STYLE | CFM_ANIMATION | CFM_COOKIE

	CFE_SMALLCAPS = CFM_SMALLCAPS
	CFE_ALLCAPS   = CFM_ALLCAPS
	CFE_HIDDEN    = CFM_HIDDEN
	CFE_OUTLINE   = CFM_OUTLINE
	CFE_SHADOW    = CFM_SHADOW
	CFE_EMBOSS    = CFM_EMBOSS
	CFE_IMPRINT   = CFM_IMPRINT
	CFE_DISABLED  = CFM_DISABLED
	CFE_REVISED   = CFM_REVISED

	// CFE_AUTOCOLOR and CFE_AUTOBACKCOLOR correspond to CFM_COLOR and
	// CFM_BACKCOLOR, respectively, which control them
	CFE_AUTOBACKCOLOR = CFM_BACKCOLOR

	CFM_FONTBOUND     = 0x00100000
	CFM_LINKPROTECTED = 0x00800000 // Word hyperlink field
	CFM_EXTENDED      = 0x02000000
	CFM_MATHNOBUILDUP = 0x08000000
	CFM_MATH          = 0x10000000
	CFM_MATHORDINARY  = 0x20000000

	CFM_ALLEFFECTS = (CFM_EFFECTS2 | CFM_FONTBOUND | CFM_EXTENDED | CFM_MATHNOBUILDUP | CFM_MATH | CFM_MATHORDINARY)

	CFE_FONTBOUND     = 0x00100000 // Font chosen by binder, not user
	CFE_LINKPROTECTED = 0x00800000
	CFE_EXTENDED      = 0x02000000
	CFE_MATHNOBUILDUP = 0x08000000
	CFE_MATH          = 0x10000000
	CFE_MATHORDINARY  = 0x20000000

	// Underline types. RE 1.0 displays only CFU_UNDERLINE
	CFU_CF1UNDERLINE             = 0xFF // Map charformat's bit underline to CF2
	CFU_INVERT                   = 0xFE // For IME composition fake a selection
	CFU_UNDERLINETHICKLONGDASH   = 18   // (*) display as dash
	CFU_UNDERLINETHICKDOTTED     = 17   // (*) display as dot
	CFU_UNDERLINETHICKDASHDOTDOT = 16   // (*) display as dash dot dot
	CFU_UNDERLINETHICKDASHDOT    = 15   // (*) display as dash dot
	CFU_UNDERLINETHICKDASH       = 14   // (*) display as dash
	CFU_UNDERLINELONGDASH        = 13   // (*) display as dash
	CFU_UNDERLINEHEAVYWAVE       = 12   // (*) display as wave
	CFU_UNDERLINEDOUBLEWAVE      = 11   // (*) display as wave
	CFU_UNDERLINEHAIRLINE        = 10   // (*) display as single
	CFU_UNDERLINETHICK           = 9
	CFU_UNDERLINEWAVE            = 8
	CFU_UNDERLINEDASHDOTDOT      = 7
	CFU_UNDERLINEDASHDOT         = 6
	CFU_UNDERLINEDASH            = 5
	CFU_UNDERLINEDOTTED          = 4
	CFU_UNDERLINEDOUBLE          = 3 // (*) display as single
	CFU_UNDERLINEWORD            = 2 // (*) display as single
	CFU_UNDERLINE                = 1
	CFU_UNDERLINENONE            = 0
)

const YHeightCharPtsMost = 1638

const (
	// EM_SETCHARFORMAT wParam masks
	SCF_SELECTION       = 0x0001
	SCF_WORD            = 0x0002
	SCF_DEFAULT         = 0x0000 // Set default charformat or paraformat
	SCF_ALL             = 0x0004 // Not valid with SCF_SELECTION or SCF_WORD
	SCF_USEUIRULES      = 0x0008 // Modifier for SCF_SELECTION; says that came from a toolbar, etc., and  UI formatting rules should be instead of literal formatting
	SCF_ASSOCIATEFONT   = 0x0010 // Associate fontname with bCharSet (one possible for each of Western, ME, FE, Thai)
	SCF_NOKBUPDATE      = 0x0020 // Do not update KB layout for this change even if autokeyboard is on
	SCF_ASSOCIATEFONT2  = 0x0040 // Associate plane-2 (surrogate) font
	SCF_SMARTFONT       = 0x0080 // Apply font only if it can handle script (5.0)
	SCF_CHARREPFROMLCID = 0x0100 // Get character repertoire from lcid (5.0)

	SPF_DONTSETDEFAULT = 0x0002 // Suppress setting default on empty control
	SPF_SETDEFAULT     = 0x0004 // Set the default paraformat
)

const (
	// Stream formats. Flags are all in low word, since high word gives possible codepage choice.
	SF_TEXT      = 0x0001
	SF_RTF       = 0x0002
	SF_RTFNOOBJS = 0x0003 // Write only
	SF_TEXTIZED  = 0x0004 // Write only

	SF_UNICODE        = 0x0010 // Unicode file (UCS2 little endian)
	SF_USECODEPAGE    = 0x0020 // CodePage given by high word
	SF_NCRFORNONASCII = 0x40   // Output \uN for nonASCII
	SFF_WRITEXTRAPAR  = 0x80   // Output \par at end

	// Flag telling stream operations to operate on selection only
	// EM_STREAMIN	replaces current selection
	// EM_STREAMOUT streams out current selection
	SFF_SELECTION = 0x8000

	// Flag telling stream operations to ignore some FE control words having to do with FE word breaking and horiz vs vertical text.
	// Not used in RichEdit 2.0 and later
	SFF_PLAINRTF = 0x4000

	// Flag telling file stream output (SFF_SELECTION flag not set) to persist // \viewscaleN control word.
	SFF_PERSISTVIEWSCALE = 0x2000

	// Flag telling file stream input with SFF_SELECTION flag not set not to // close the document
	SFF_KEEPDOCINFO = 0x1000

	// Flag telling stream operations to output in Pocket Word format
	SFF_PWD = 0x0800

	// 3-bit field specifying the value of N - 1 to use for \rtfN or \pwdN
	SF_RTFVAL = 0x0700
)

// All paragraph measurements are in twips
const (
	MAX_TAB_STOPS   = 32
	LDefaultTab     = 720
	MAX_TABLE_CELLS = 63
)

const (
	// PARAFORMAT mask values
	PFM_STARTINDENT  = 0x00000001
	PFM_RIGHTINDENT  = 0x00000002
	PFM_OFFSET       = 0x00000004
	PFM_ALIGNMENT    = 0x00000008
	PFM_TABSTOPS     = 0x00000010
	PFM_NUMBERING    = 0x00000020
	PFM_OFFSETINDENT = 0x80000000

	// PARAFORMAT 2.0 masks and effects
	PFM_SPACEBEFORE    = 0x00000040
	PFM_SPACEAFTER     = 0x00000080
	PFM_LINESPACING    = 0x00000100
	PFM_STYLE          = 0x00000400
	PFM_BORDER         = 0x00000800 // (*)
	PFM_SHADING        = 0x00001000 // (*)
	PFM_NUMBERINGSTYLE = 0x00002000 // RE 3.0
	PFM_NUMBERINGTAB   = 0x00004000 // RE 3.0
	PFM_NUMBERINGSTART = 0x00008000 // RE 3.0

	PFM_RTLPARA         = 0x00010000
	PFM_KEEP            = 0x00020000 // (*)
	PFM_KEEPNEXT        = 0x00040000 // (*)
	PFM_PAGEBREAKBEFORE = 0x00080000 // (*)
	PFM_NOLINENUMBER    = 0x00100000 // (*)
	PFM_NOWIDOWCONTROL  = 0x00200000 // (*)
	PFM_DONOTHYPHEN     = 0x00400000 // (*)
	PFM_SIDEBYSIDE      = 0x00800000 // (*)

	// The following two paragraph-format properties are read only
	PFM_COLLAPSED         = 0x01000000 // RE 3.0
	PFM_OUTLINELEVEL      = 0x02000000 // RE 3.0
	PFM_BOX               = 0x04000000 // RE 3.0
	PFM_RESERVED2         = 0x08000000 // RE 4.0
	PFM_TABLEROWDELIMITER = 0x10000000 // RE 4.0
	PFM_TEXTWRAPPINGBREAK = 0x20000000 // RE 3.0
	PFM_TABLE             = 0x40000000 // RE 3.0

	// PARAFORMAT "ALL" masks
	PFM_ALL = PFM_STARTINDENT | PFM_RIGHTINDENT | PFM_OFFSET | PFM_ALIGNMENT | PFM_TABSTOPS | PFM_NUMBERING | PFM_OFFSETINDENT | PFM_RTLPARA

	// Note: PARAFORMAT has no effects (BiDi RichEdit 1.0 does have PFE_RTLPARA)
	PFM_EFFECTS = PFM_RTLPARA | PFM_KEEP | PFM_KEEPNEXT | PFM_TABLE | PFM_PAGEBREAKBEFORE | PFM_NOLINENUMBER | PFM_NOWIDOWCONTROL | PFM_DONOTHYPHEN | PFM_SIDEBYSIDE | PFM_TABLE | PFM_TABLEROWDELIMITER

	PFM_ALL2 = PFM_ALL | PFM_EFFECTS | PFM_SPACEBEFORE | PFM_SPACEAFTER | PFM_LINESPACING | PFM_STYLE | PFM_SHADING | PFM_BORDER | PFM_NUMBERINGTAB | PFM_NUMBERINGSTART | PFM_NUMBERINGSTYLE

	PFE_RTLPARA           = PFM_RTLPARA >> 16
	PFE_KEEP              = PFM_KEEP >> 16              // (*)
	PFE_KEEPNEXT          = PFM_KEEPNEXT >> 16          // (*)
	PFE_PAGEBREAKBEFORE   = PFM_PAGEBREAKBEFORE >> 16   // (*)
	PFE_NOLINENUMBER      = PFM_NOLINENUMBER >> 16      // (*)
	PFE_NOWIDOWCONTROL    = PFM_NOWIDOWCONTROL >> 16    // (*)
	PFE_DONOTHYPHEN       = PFM_DONOTHYPHEN >> 16       // (*)
	PFE_SIDEBYSIDE        = PFM_SIDEBYSIDE >> 16        // (*)
	PFE_TEXTWRAPPINGBREAK = PFM_TEXTWRAPPINGBREAK >> 16 // (*)

	// The following four effects are read only
	PFE_COLLAPSED         = PFM_COLLAPSED >> 16         // (+)
	PFE_BOX               = PFM_BOX >> 16               // (+)
	PFE_TABLE             = PFM_TABLE >> 16             // Inside table row. RE 3.0
	PFE_TABLEROWDELIMITER = PFM_TABLEROWDELIMITER >> 16 // Table row start. RE 4.0

	// PARAFORMAT numbering options
	PFN_BULLET = 1 // tomListBullet

	// PARAFORMAT2 wNumbering options
	PFN_ARABIC   = 2 // tomListNumberAsArabic:	0, 1, 2,	...
	PFN_LCLETTER = 3 // tomListNumberAsLCLetter: a, b, c,	...
	PFN_UCLETTER = 4 // tomListNumberAsUCLetter: A, B, C,	...
	PFN_LCROMAN  = 5 // tomListNumberAsLCRoman:	i, ii, iii, ...
	PFN_UCROMAN  = 6 // tomListNumberAsUCRoman:	I, II, III, ...

	// PARAFORMAT2 wNumberingStyle options
	PFNS_PAREN    = 0x000 // default, e.g.,				  1)
	PFNS_PARENS   = 0x100 // tomListParentheses/256, e.g., (1)
	PFNS_PERIOD   = 0x200 // tomListPeriod/256, e.g., 	  1.
	PFNS_PLAIN    = 0x300 // tomListPlain/256, e.g.,		  1
	PFNS_NONUMBER = 0x400 // Used for continuation w/o number

	PFNS_NEWNUMBER = 0x8000 // Start new number with wNumberingStart
	// (can be combined with other PFNS_xxx)
	// PARAFORMAT alignment options
	PFA_LEFT   = 1
	PFA_RIGHT  = 2
	PFA_CENTER = 3

	// PARAFORMAT2 alignment options
	PFA_JUSTIFY        = 4 // New paragraph-alignment option 2.0 (*)
	PFA_FULL_INTERWORD = 4 // These are supported in 3.0 with advanced
)

const (
	SEL_EMPTY       = 0x0000
	SEL_TEXT        = 0x0001
	SEL_OBJECT      = 0x0002
	SEL_MULTICHAR   = 0x0004
	SEL_MULTIOBJECT = 0x0008
)

const (
	// Used with IRichEditOleCallback::GetContextMenu, this flag will be passed as a "selection type".  It indicates that a context menu for a right-mouse drag drop should be generated.  The IOleObject parameter will really be the IDataObject for the drop
	GCM_RIGHTMOUSEDROP = 0x8000
)

const (
	// bits for GETCONTEXTMENUEX::dwFlags
	GCMF_GRIPPER   = 0x00000001
	GCMF_SPELLING  = 0x00000002 // pSpellingSuggestions is valid and points to the list of spelling suggestions
	GCMF_TOUCHMENU = 0x00004000
	GCMF_MOUSEMENU = 0x00002000
)

const OLEOP_DOVERB = 1

const (
	// Clipboard formats - use as parameter to RegisterClipboardFormat()
	CF_RTF       = "Rich Text Format"
	CF_RTFNOOBJS = "Rich Text Format Without Objects"
	CF_RETEXTOBJ = "RichEdit Text and Objects"
)

const (
	UID_UNKNOWN   UNDONAMEID = 0
	UID_TYPING               = 1
	UID_DELETE               = 2
	UID_DRAGDROP             = 3
	UID_CUT                  = 4
	UID_PASTE                = 5
	UID_AUTOTABLE            = 6
)

const (
	// Flags for the SETEXTEX data structure
	ST_DEFAULT   = 0
	ST_KEEPUNDO  = 1
	ST_SELECTION = 2
	ST_NEuint16S = 4
	ST_UNICODE   = 8
)

const (
	// Flags for the GETEXTEX data structure
	GT_DEFAULT      = 0
	GT_USECRLF      = 1
	GT_SELECTION    = 2
	GT_RAWTEXT      = 4
	GT_NOHIDDENTEXT = 8
)

const (
	// Flags for the GETTEXTLENGTHEX data structure
	GTL_DEFAULT  = 0  // Do default (return # of chars)
	GTL_USECRLF  = 1  // Compute answer using CRLFs for paragraphs
	GTL_PRECISE  = 2  // Compute a precise answer
	GTL_CLOSE    = 4  // Fast computation of a "close" answer
	GTL_NUMCHARS = 8  // Return number of characters
	GTL_NUMBYTES = 16 // Return number of _bytes_
)

const (
	// BIDIOPTIONS masks
	BOM_NEUTRALOVERRIDE  = 0x0004 // Override neutral layout (obsolete)
	BOM_CONTEXTREADING   = 0x0008 // Context reading order
	BOM_CONTEXTALIGNMENT = 0x0010 // Context alignment
	BOM_LEGACYBIDICLASS  = 0x0040 // Legacy Bidi classification (obsolete)
	BOM_UNICODEBIDI      = 0x0080 // Use Unicode BiDi algorithm

	// BIDIOPTIONS effects
	BOE_NEUTRALOVERRIDE  = 0x0004 // Override neutral layout (obsolete)
	BOE_CONTEXTREADING   = 0x0008 // Context reading order
	BOE_CONTEXTALIGNMENT = 0x0010 // Context alignment
	BOE_FORCERECALC      = 0x0020 // Force recalc and redraw
	BOE_LEGACYBIDICLASS  = 0x0040 // Legacy Bidi classification (obsolete)
	BOE_UNICODEBIDI      = 0x0080 // Use Unicode BiDi algorithm

	// Additional EM_FINDTEXT[EX] flags
	FR_MATCHDIAC      = 0x20000000
	FR_MATCHKASHIDA   = 0x40000000
	FR_MATCHALEFHAMZA = 0x80000000

	// UNICODE embedding character
	WCH_EMBEDDING uint16 = 0xFFFC
)

const (
	KhyphNil          KHYPH = iota // No Hyphenation
	KhyphNormal                    // Normal Hyphenation
	KhyphAddBefore                 // Add letter before hyphen
	KhyphChangeBefore              // Change letter before hyphen
	KhyphDeleteBefore              // Delete letter before hyphen
	KhyphChangeAfter               // Change letter after hyphen
	KhyphDelAndChange              // Delete letter before hyphen and change letter preceding hyphen
)

const (
	// Additional class for Richedit 6.0
	RICHEDIT60_CLASS = "RICHEDIT60W"
)
