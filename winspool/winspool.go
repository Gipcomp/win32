// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winspool

import "golang.org/x/sys/windows"

var (
	// Library
	libwinspool *windows.LazyDLL

	// Functions
	deviceCapabilities *windows.LazyProc
	documentProperties *windows.LazyProc
	enumPrinters       *windows.LazyProc
	getDefaultPrinter  *windows.LazyProc
)

func init() {
	// Library
	libwinspool = windows.NewLazySystemDLL("winspool.drv")

	// Functions
	deviceCapabilities = libwinspool.NewProc("DeviceCapabilitiesW")
	documentProperties = libwinspool.NewProc("DocumentPropertiesW")
	enumPrinters = libwinspool.NewProc("EnumPrintersW")
	getDefaultPrinter = libwinspool.NewProc("GetDefaultPrinterW")
}
