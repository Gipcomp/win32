// Copyright 2013 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package pdh

import "golang.org/x/sys/windows"

var (
	// Library
	libpdhDll *windows.LazyDLL

	// Functions
	pdh_AddCounterW               *windows.LazyProc
	pdh_AddEnglishCounterW        *windows.LazyProc
	pdh_CloseQuery                *windows.LazyProc
	pdh_CollectQueryData          *windows.LazyProc
	pdh_GetFormattedCounterValue  *windows.LazyProc
	pdh_GetFormattedCounterArrayW *windows.LazyProc
	pdh_OpenQuery                 *windows.LazyProc
	pdh_ValidatePathW             *windows.LazyProc
)

func init() {
	// Library
	libpdhDll = windows.NewLazySystemDLL("pdh.dll")

	// Functions
	pdh_AddCounterW = libpdhDll.NewProc("PdhAddCounterW")
	pdh_AddEnglishCounterW = libpdhDll.NewProc("PdhAddEnglishCounterW")
	pdh_CloseQuery = libpdhDll.NewProc("PdhCloseQuery")
	pdh_CollectQueryData = libpdhDll.NewProc("PdhCollectQueryData")
	pdh_GetFormattedCounterValue = libpdhDll.NewProc("PdhGetFormattedCounterValue")
	pdh_GetFormattedCounterArrayW = libpdhDll.NewProc("PdhGetFormattedCounterArrayW")
	pdh_OpenQuery = libpdhDll.NewProc("PdhOpenQuery")
	pdh_ValidatePathW = libpdhDll.NewProc("PdhValidatePathW")
}
