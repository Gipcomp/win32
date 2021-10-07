// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oaidl

import "github.com/Gipcomp/win32/ole32"

type SCODE int32

type EXCEPINFO struct {
	wCode             uint16
	wReserved         uint16
	bstrSource        *uint16 /*BSTR*/
	bstrDescription   *uint16 /*BSTR*/
	bstrHelpFile      *uint16 /*BSTR*/
	dwHelpContext     uint32
	pvReserved        uintptr
	pfnDeferredFillIn uintptr
	scode             SCODE
}

type ITypeInfoVtbl struct {
	ole32.IUnknownVtbl
	GetTypeAttr          uintptr
	GetTypeComp          uintptr
	GetFuncDesc          uintptr
	GetVarDesc           uintptr
	GetNames             uintptr
	GetRefTypeOfImplType uintptr
	GetImplTypeFlags     uintptr
	GetIDsOfNames        uintptr
	Invoke               uintptr
	GetDocumentation     uintptr
	GetDllEntry          uintptr
	GetRefTypeInfo       uintptr
	AddressOfMember      uintptr
	CreateInstance       uintptr
	GetMops              uintptr
	GetContainingTypeLib uintptr
	ReleaseTypeAttr      uintptr
	ReleaseFuncDesc      uintptr
	ReleaseVarDesc       uintptr
}

type ITypeInfo struct {
	LpVtbl *ITypeInfoVtbl
}
