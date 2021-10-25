// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package ole32

import "golang.org/x/sys/windows"

var (
	IID_IClassFactory             = IID{0x00000001, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IConnectionPointContainer = IID{0xB196B284, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
	IID_IOleClientSite            = IID{0x00000118, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IOleInPlaceObject         = IID{0x00000113, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IOleInPlaceSite           = IID{0x00000119, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IOleObject                = IID{0x00000112, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IUnknown                  = IID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
)

var (
	// Library
	libole32 *windows.LazyDLL

	// Functions
	coCreateInstance      *windows.LazyProc
	coGetClassObject      *windows.LazyProc
	coInitializeEx        *windows.LazyProc
	coTaskMemFree         *windows.LazyProc
	coUninitialize        *windows.LazyProc
	oleInitialize         *windows.LazyProc
	oleSetContainedObject *windows.LazyProc
	oleUninitialize       *windows.LazyProc
	CoInitializeExCall    *windows.LazyProc
)

func init() {
	// Library
	libole32 = windows.NewLazySystemDLL("ole32.dll")

	// Functions
	coCreateInstance = libole32.NewProc("CoCreateInstance")
	coGetClassObject = libole32.NewProc("CoGetClassObject")
	coInitializeEx = libole32.NewProc("CoInitializeEx")
	coTaskMemFree = libole32.NewProc("CoTaskMemFree")
	coUninitialize = libole32.NewProc("CoUninitialize")
	oleInitialize = libole32.NewProc("OleInitialize")
	oleSetContainedObject = libole32.NewProc("OleSetContainedObject")
	oleUninitialize = libole32.NewProc("OleUninitialize")
}
