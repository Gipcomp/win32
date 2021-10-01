// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oleacc

import "syscall"

type AnnoScope int

type MSAAPROPID syscall.GUID

type IAccPropServerVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	GetPropValue   uintptr
}

type IAccPropServer struct {
	LpVtbl *IAccPropServerVtbl
}

type IAccPropServicesVtbl struct {
	QueryInterface               uintptr
	AddRef                       uintptr
	Release                      uintptr
	SetPropValue                 uintptr
	SetPropServer                uintptr
	ClearProps                   uintptr
	SetHwndProp                  uintptr
	SetHwndPropStr               uintptr
	SetHwndPropServer            uintptr
	ClearHwndProps               uintptr
	ComposeHwndIdentityString    uintptr
	DecomposeHwndIdentityString  uintptr
	SetHmenuProp                 uintptr
	SetHmenuPropStr              uintptr
	SetHmenuPropServer           uintptr
	ClearHmenuProps              uintptr
	ComposeHmenuIdentityString   uintptr
	DecomposeHmenuIdentityString uintptr
}

type IAccPropServices struct {
	LpVtbl *IAccPropServicesVtbl
}
