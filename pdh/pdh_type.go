// Copyright 2013 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package pdh

import "github.com/gfphoenix/win32/handle"

type (
	PDH_HQUERY   handle.HANDLE // query handle
	PDH_HCOUNTER handle.HANDLE // counter handle
)

// Union specialization for double values
type PDH_FMT_COUNTERVALUE_DOUBLE struct {
	CStatus     uint32
	DoubleValue float64
}

// Union specialization for 64 bit integer values
type PDH_FMT_COUNTERVALUE_LARGE struct {
	CStatus    uint32
	LargeValue int64
}

// Union specialization for long values
type PDH_FMT_COUNTERVALUE_LONG struct {
	CStatus   uint32
	LongValue int32
	padding   [4]byte
}

// Union specialization for double values, used by PdhGetFormattedCounterArrayDouble()
type PDH_FMT_COUNTERVALUE_ITEM_DOUBLE struct {
	SzName   *uint16 // pointer to a string
	FmtValue PDH_FMT_COUNTERVALUE_DOUBLE
}

// Union specialization for 'large' values, used by PdhGetFormattedCounterArrayLarge()
type PDH_FMT_COUNTERVALUE_ITEM_LARGE struct {
	SzName   *uint16 // pointer to a string
	FmtValue PDH_FMT_COUNTERVALUE_LARGE
}

// Union specialization for long values, used by PdhGetFormattedCounterArrayLong()
type PDH_FMT_COUNTERVALUE_ITEM_LONG struct {
	SzName   *uint16 // pointer to a string
	FmtValue PDH_FMT_COUNTERVALUE_LONG
}
