// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oleaut32

type DISPID int32

type IDispatchVtbl struct {
	QueryInterface   uintptr
	AddRef           uintptr
	Release          uintptr
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

type IDispatch struct {
	LpVtbl *IDispatchVtbl
}

type VARTYPE uint16

type DISPPARAMS struct {
	Rgvarg            *VARIANTARG
	RgdispidNamedArgs *DISPID
	CArgs             int32
	CNamedArgs        int32
}
