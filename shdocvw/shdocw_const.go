// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shdocvw

const (
	DOCHOSTUIDBLCLK_DEFAULT        = 0
	DOCHOSTUIDBLCLK_SHOWPROPERTIES = 1
	DOCHOSTUIDBLCLK_SHOWCODE       = 2
)

const (
	DOCHOSTUIFLAG_DIALOG                     = 0x1
	DOCHOSTUIFLAG_DISABLE_HELP_MENU          = 0x2
	DOCHOSTUIFLAG_NO3DBORDER                 = 0x4
	DOCHOSTUIFLAG_SCROLL_NO                  = 0x8
	DOCHOSTUIFLAG_DISABLE_SCRIPT_INACTIVE    = 0x10
	DOCHOSTUIFLAG_OPENNEWWIN                 = 0x20
	DOCHOSTUIFLAG_DISABLE_OFFSCREEN          = 0x40
	DOCHOSTUIFLAG_FLAT_SCROLLBAR             = 0x80
	DOCHOSTUIFLAG_DIV_BLOCKDEFAULT           = 0x100
	DOCHOSTUIFLAG_ACTIVATE_CLIENTHIT_ONLY    = 0x200
	DOCHOSTUIFLAG_OVERRIDEBEHAVIORFACTORY    = 0x400
	DOCHOSTUIFLAG_CODEPAGELINKEDFONTS        = 0x800
	DOCHOSTUIFLAG_URL_ENCODING_DISABLE_UTF8  = 0x1000
	DOCHOSTUIFLAG_URL_ENCODING_ENABLE_UTF8   = 0x2000
	DOCHOSTUIFLAG_ENABLE_FORMS_AUTOCOMPLETE  = 0x4000
	DOCHOSTUIFLAG_ENABLE_INPLACE_NAVIGATION  = 0x10000
	DOCHOSTUIFLAG_IME_ENABLE_RECONVERSION    = 0x20000
	DOCHOSTUIFLAG_THEME                      = 0x40000
	DOCHOSTUIFLAG_NOTHEME                    = 0x80000
	DOCHOSTUIFLAG_NOPICS                     = 0x100000
	DOCHOSTUIFLAG_NO3DOUTERBORDER            = 0x200000
	DOCHOSTUIFLAG_DISABLE_EDIT_NS_FIXUP      = 0x400000
	DOCHOSTUIFLAG_LOCAL_MACHINE_ACCESS_CHECK = 0x800000
	DOCHOSTUIFLAG_DISABLE_UNTRUSTEDPROTOCOL  = 0x1000000
)

// BrowserNavConstants
const (
	NavOpenInNewWindow       = 0x1
	NavNoHistory             = 0x2
	NavNoReadFromCache       = 0x4
	NavNoWriteToCache        = 0x8
	NavAllowAutosearch       = 0x10
	NavBrowserBar            = 0x20
	NavHyperlink             = 0x40
	NavEnforceRestricted     = 0x80
	NavNewWindowsManaged     = 0x0100
	NavUntrustedForDownload  = 0x0200
	NavTrustedForActiveX     = 0x0400
	NavOpenInNewTab          = 0x0800
	NavOpenInBackgroundTab   = 0x1000
	NavKeepWordWheelText     = 0x2000
	NavVirtualTab            = 0x4000
	NavBlockRedirectsXDomain = 0x8000
	NavOpenNewForegroundTab  = 0x10000
)
