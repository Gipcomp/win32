// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package advapi32

import "golang.org/x/sys/windows"

var (
	// Library
	libadvapi32 *windows.LazyDLL

	// Functions
	regCloseKey     *windows.LazyProc
	regOpenKeyEx    *windows.LazyProc
	regQueryValueEx *windows.LazyProc
	regEnumValue    *windows.LazyProc
	regSetValueEx   *windows.LazyProc
)

func init() {
	// Library
	libadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	// Functions
	regCloseKey = libadvapi32.NewProc("RegCloseKey")
	regOpenKeyEx = libadvapi32.NewProc("RegOpenKeyExW")
	regQueryValueEx = libadvapi32.NewProc("RegQueryValueExW")
	regEnumValue = libadvapi32.NewProc("RegEnumValueW")
	regSetValueEx = libadvapi32.NewProc("RegSetValueExW")
}
