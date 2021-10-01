// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package comctl32

import "golang.org/x/sys/windows"

var (
	// Library
	libcomctl32 *windows.LazyDLL

	// Functions
	imageList_Add         *windows.LazyProc
	imageList_AddMasked   *windows.LazyProc
	imageList_Create      *windows.LazyProc
	imageList_Destroy     *windows.LazyProc
	imageList_DrawEx      *windows.LazyProc
	imageList_ReplaceIcon *windows.LazyProc
	initCommonControlsEx  *windows.LazyProc
	loadIconMetric        *windows.LazyProc
	loadIconWithScaleDown *windows.LazyProc
)

func init() {
	// Library
	libcomctl32 = windows.NewLazySystemDLL("comctl32.dll")

	// Functions
	imageList_Add = libcomctl32.NewProc("ImageList_Add")
	imageList_AddMasked = libcomctl32.NewProc("ImageList_AddMaskd")
	imageList_Create = libcomctl32.NewProc("ImageList_Create")
	imageList_Destroy = libcomctl32.NewProc("ImageList_Destroy")
	imageList_DrawEx = libcomctl32.NewProc("ImageList_DrawEx")
	imageList_ReplaceIcon = libcomctl32.NewProc("ImageListReplaceIcon")
	initCommonControlsEx = libcomctl32.NewProc("InitCommonControlsEx")
	loadIconMetric = libcomctl32.NewProc("LoadIconMetric")

	loadIconWithScaleDown = libcomctl32.NewProc("LoadIconWithScaleDown")
}
