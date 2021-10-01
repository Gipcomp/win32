// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oleacc

import "github.com/D4v1dW3bb/winapi/ole32"

var (
	PROPID_ACC_NAME             = MSAAPROPID{0x608d3df8, 0x8128, 0x4aa7, [8]byte{0xa4, 0x28, 0xf5, 0x5e, 0x49, 0x26, 0x72, 0x91}}
	PROPID_ACC_VALUE            = MSAAPROPID{0x123fe443, 0x211a, 0x4615, [8]byte{0x95, 0x27, 0xc4, 0x5a, 0x7e, 0x93, 0x71, 0x7a}}
	PROPID_ACC_DESCRIPTION      = MSAAPROPID{0x4d48dfe4, 0xbd3f, 0x491f, [8]byte{0xa6, 0x48, 0x49, 0x2d, 0x6f, 0x20, 0xc5, 0x88}}
	PROPID_ACC_ROLE             = MSAAPROPID{0xcb905ff2, 0x7bd1, 0x4c05, [8]byte{0xb3, 0xc8, 0xe6, 0xc2, 0x41, 0x36, 0x4d, 0x70}}
	PROPID_ACC_STATE            = MSAAPROPID{0xa8d4d5b0, 0x0a21, 0x42d0, [8]byte{0xa5, 0xc0, 0x51, 0x4e, 0x98, 0x4f, 0x45, 0x7b}}
	PROPID_ACC_HELP             = MSAAPROPID{0xc831e11f, 0x44db, 0x4a99, [8]byte{0x97, 0x68, 0xcb, 0x8f, 0x97, 0x8b, 0x72, 0x31}}
	PROPID_ACC_KEYBOARDSHORTCUT = MSAAPROPID{0x7d9bceee, 0x7d1e, 0x4979, [8]byte{0x93, 0x82, 0x51, 0x80, 0xf4, 0x17, 0x2c, 0x34}}
	PROPID_ACC_DEFAULTACTION    = MSAAPROPID{0x180c072b, 0xc27f, 0x43c7, [8]byte{0x99, 0x22, 0xf6, 0x35, 0x62, 0xa4, 0x63, 0x2b}}
	PROPID_ACC_HELPTOPIC        = MSAAPROPID{0x787d1379, 0x8ede, 0x440b, [8]byte{0x8a, 0xec, 0x11, 0xf7, 0xbf, 0x90, 0x30, 0xb3}}
	PROPID_ACC_FOCUS            = MSAAPROPID{0x6eb335df, 0x1c29, 0x4127, [8]byte{0xb1, 0x2c, 0xde, 0xe9, 0xfd, 0x15, 0x7f, 0x2b}}
	PROPID_ACC_SELECTION        = MSAAPROPID{0xb99d073c, 0xd731, 0x405b, [8]byte{0x90, 0x61, 0xd9, 0x5e, 0x8f, 0x84, 0x29, 0x84}}
	PROPID_ACC_PARENT           = MSAAPROPID{0x474c22b6, 0xffc2, 0x467a, [8]byte{0xb1, 0xb5, 0xe9, 0x58, 0xb4, 0x65, 0x73, 0x30}}
	PROPID_ACC_NAV_UP           = MSAAPROPID{0x016e1a2b, 0x1a4e, 0x4767, [8]byte{0x86, 0x12, 0x33, 0x86, 0xf6, 0x69, 0x35, 0xec}}
	PROPID_ACC_NAV_DOWN         = MSAAPROPID{0x031670ed, 0x3cdf, 0x48d2, [8]byte{0x96, 0x13, 0x13, 0x8f, 0x2d, 0xd8, 0xa6, 0x68}}
	PROPID_ACC_NAV_LEFT         = MSAAPROPID{0x228086cb, 0x82f1, 0x4a39, [8]byte{0x87, 0x05, 0xdc, 0xdc, 0x0f, 0xff, 0x92, 0xf5}}
	PROPID_ACC_NAV_RIGHT        = MSAAPROPID{0xcd211d9f, 0xe1cb, 0x4fe5, [8]byte{0xa7, 0x7c, 0x92, 0x0b, 0x88, 0x4d, 0x09, 0x5b}}
	PROPID_ACC_NAV_PREV         = MSAAPROPID{0x776d3891, 0xc73b, 0x4480, [8]byte{0xb3, 0xf6, 0x07, 0x6a, 0x16, 0xa1, 0x5a, 0xf6}}
	PROPID_ACC_NAV_NEXT         = MSAAPROPID{0x1cdc5455, 0x8cd9, 0x4c92, [8]byte{0xa3, 0x71, 0x39, 0x39, 0xa2, 0xfe, 0x3e, 0xee}}
	PROPID_ACC_NAV_FIRSTCHILD   = MSAAPROPID{0xcfd02558, 0x557b, 0x4c67, [8]byte{0x84, 0xf9, 0x2a, 0x09, 0xfc, 0xe4, 0x07, 0x49}}
	PROPID_ACC_NAV_LASTCHILD    = MSAAPROPID{0x302ecaa5, 0x48d5, 0x4f8d, [8]byte{0xb6, 0x71, 0x1a, 0x8d, 0x20, 0xa7, 0x78, 0x32}}
	PROPID_ACC_ROLEMAP          = MSAAPROPID{0xf79acda2, 0x140d, 0x4fe6, [8]byte{0x89, 0x14, 0x20, 0x84, 0x76, 0x32, 0x82, 0x69}}
	PROPID_ACC_VALUEMAP         = MSAAPROPID{0xda1c3d79, 0xfc5c, 0x420e, [8]byte{0xb3, 0x99, 0x9d, 0x15, 0x33, 0x54, 0x9e, 0x75}}
	PROPID_ACC_STATEMAP         = MSAAPROPID{0x43946c5e, 0x0ac0, 0x4042, [8]byte{0xb5, 0x25, 0x07, 0xbb, 0xdb, 0xe1, 0x7f, 0xa7}}
	PROPID_ACC_DESCRIPTIONMAP   = MSAAPROPID{0x1ff1435f, 0x8a14, 0x477b, [8]byte{0xb2, 0x26, 0xa0, 0xab, 0xe2, 0x79, 0x97, 0x5d}}
	PROPID_ACC_DODEFAULTACTION  = MSAAPROPID{0x1ba09523, 0x2e3b, 0x49a6, [8]byte{0xa0, 0x59, 0x59, 0x68, 0x2a, 0x3c, 0x48, 0xfd}}
)

var (
	IID_IAccPropServer    = ole32.IID{0x76c0dbbb, 0x15e0, 0x4e7b, [8]byte{0xb6, 0x1b, 0x20, 0xee, 0xea, 0x20, 0x01, 0xe0}}
	IID_IAccPropServices  = ole32.IID{0x6e26e776, 0x04f0, 0x495d, [8]byte{0x80, 0xe4, 0x33, 0x30, 0x35, 0x2e, 0x31, 0x69}}
	CLSID_AccPropServices = ole32.CLSID{0xb5f8350b, 0x0548, 0x48b1, [8]byte{0xa6, 0xee, 0x88, 0xbd, 0x00, 0xb4, 0xa5, 0xe7}}
)
