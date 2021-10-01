// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"github.com/D4v1dW3bb/winapi/handle"
	"github.com/D4v1dW3bb/winapi/kernel32"
)

type (
	BOOL    int32
	HRESULT int32
)

type WINEVENTPROC func(hWinEventHook kernel32.HWINEVENTHOOK, event uint32, hwnd handle.HWND, idObject int32, idChild int32, idEventThread uint32, dwmsEventTime uint32) uintptr
