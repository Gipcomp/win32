// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package advapi32

import "github.com/gfphoenix/win32/handle"

type (
	ACCESS_MASK uint32
	HKEY        handle.HANDLE
	REGSAM      ACCESS_MASK
)
