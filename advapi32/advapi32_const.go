// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package advapi32

// Button control messages
const KEY_READ REGSAM = 0x20019
const KEY_WRITE REGSAM = 0x20006

const (
	HKEY_CLASSES_ROOT     HKEY = 0x80000000
	HKEY_CURRENT_USER     HKEY = 0x80000001
	HKEY_LOCAL_MACHINE    HKEY = 0x80000002
	HKEY_USERS            HKEY = 0x80000003
	HKEY_PERFORMANCE_DATA HKEY = 0x80000004
	HKEY_CURRENT_CONFIG   HKEY = 0x80000005
	HKEY_DYN_DATA         HKEY = 0x80000006
)

const (
	ERROR_NO_MORE_ITEMS = 259
)

const (
	REG_NONE      uint64 = 0 // No value type
	REG_SZ               = 1 // Unicode nul terminated string
	REG_EXPAND_SZ        = 2 // Unicode nul terminated string

	// (with environment variable references)
	REG_BINARY                     = 3 // Free form binary
	REG_DWORD                      = 4 // 32-bit number
	REG_DWORD_LITTLE_ENDIAN        = 4 // 32-bit number (same as REG_DWORD)
	REG_DWORD_BIG_ENDIAN           = 5 // 32-bit number
	REG_LINK                       = 6 // Symbolic Link (unicode)
	REG_MULTI_SZ                   = 7 // Multiple Unicode strings
	REG_RESOURCE_LIST              = 8 // Resource list in the resource map
	REG_FULL_RESOURCE_DESCRIPTOR   = 9 // Resource list in the hardware description
	REG_RESOURCE_REQUIREMENTS_LIST = 10
	REG_QWORD                      = 11 // 64-bit number
	REG_QWORD_LITTLE_ENDIAN        = 11 // 64-bit number (same as REG_QWORD)

)
