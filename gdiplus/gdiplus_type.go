// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package gdiplus

import "github.com/Gipcomp/win32/win"

type GpStatus int32

type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread win.BOOL
	SuppressExternalCodecs   win.BOOL
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type GpImage struct{}

type GpBitmap GpImage

type ARGB uint32
