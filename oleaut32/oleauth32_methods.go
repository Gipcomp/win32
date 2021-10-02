// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package oleaut32

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/D4v1dW3bb/winapi/win"
)

func StringToBSTR(value string) *uint16 /*BSTR*/ {
	// IMPORTANT: Don't forget to free the BSTR value when no longer needed!
	return SysAllocString(value)
}

func BSTRToString(value *uint16 /*BSTR*/) string {
	// ISSUE: Is this really ok?
	bstrArrPtr := (*[200000000]uint16)(unsafe.Pointer(value))

	bstrSlice := make([]uint16, SysStringLen(value))
	copy(bstrSlice, bstrArrPtr[:])

	return syscall.UTF16ToString(bstrSlice)
}

func IntToVariantI4(value int32) *VAR_I4 {
	return &VAR_I4{vt: VT_I4, lVal: value}
}

func VariantI4ToInt(value *VAR_I4) int32 {
	return value.lVal
}

func BoolToVariantBool(value bool) *VAR_BOOL {
	return &VAR_BOOL{vt: VT_BOOL, boolVal: VARIANT_BOOL(win.BoolToBOOL(value))}
}

func VariantBoolToBool(value *VAR_BOOL) bool {
	return value.boolVal != 0
}

func StringToVariantBSTR(value string) *VAR_BSTR {
	// IMPORTANT: Don't forget to free the BSTR value when no longer needed!
	return &VAR_BSTR{vt: VT_BSTR, bstrVal: StringToBSTR(value)}
}

func VariantBSTRToString(value *VAR_BSTR) string {
	return BSTRToString(value.bstrVal)
}

func (v *VARIANT) MustLong() int32 {
	value, err := v.Long()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) Long() (int32, error) {
	if v.Vt != VT_I4 {
		return 0, fmt.Errorf("Error: Long() v.Vt !=  VT_I4, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_I4)(unsafe.Pointer(v))
	return p.lVal, nil
}

func (v *VARIANT) SetLong(value int32) {
	v.Vt = VT_I4
	p := (*VAR_I4)(unsafe.Pointer(v))
	p.lVal = value
}

func (v *VARIANT) MustULong() uint32 {
	value, err := v.ULong()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) ULong() (uint32, error) {
	if v.Vt != VT_UI4 {
		return 0, fmt.Errorf("Error: ULong() v.Vt !=  VT_UI4, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_UI4)(unsafe.Pointer(v))
	return p.ulVal, nil
}

func (v *VARIANT) SetULong(value uint32) {
	v.Vt = VT_UI4
	p := (*VAR_UI4)(unsafe.Pointer(v))
	p.ulVal = value
}

func (v *VARIANT) MustBool() VARIANT_BOOL {
	value, err := v.Bool()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) Bool() (VARIANT_BOOL, error) {
	if v.Vt != VT_BOOL {
		return VARIANT_FALSE, fmt.Errorf("Error: Bool() v.Vt !=  VT_BOOL, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_BOOL)(unsafe.Pointer(v))
	return p.boolVal, nil
}

func (v *VARIANT) SetBool(value VARIANT_BOOL) {
	v.Vt = VT_BOOL
	p := (*VAR_BOOL)(unsafe.Pointer(v))
	p.boolVal = value
}

func (v *VARIANT) MustBSTR() *uint16 {
	value, err := v.BSTR()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) BSTR() (*uint16, error) {
	if v.Vt != VT_BSTR {
		return nil, fmt.Errorf("Error: BSTR() v.Vt !=  VT_BSTR, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_BSTR)(unsafe.Pointer(v))
	return p.bstrVal, nil
}

func (v *VARIANT) SetBSTR(value *uint16) {
	v.Vt = VT_BSTR
	p := (*VAR_BSTR)(unsafe.Pointer(v))
	p.bstrVal = value
}

func (v *VARIANT) MustPDispatch() *IDispatch {
	value, err := v.PDispatch()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) PDispatch() (*IDispatch, error) {
	if v.Vt != VT_DISPATCH {
		return nil, fmt.Errorf("Error: PDispatch() v.Vt !=  VT_DISPATCH, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_PDISP)(unsafe.Pointer(v))
	return p.pdispVal, nil
}

func (v *VARIANT) SetPDispatch(value *IDispatch) {
	v.Vt = VT_DISPATCH
	p := (*VAR_PDISP)(unsafe.Pointer(v))
	p.pdispVal = value
}

func (v *VARIANT) MustPVariant() *VARIANT {
	value, err := v.PVariant()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) PVariant() (*VARIANT, error) {
	if v.Vt != VT_BYREF|VT_VARIANT {
		return nil, fmt.Errorf("Error: PVariant() v.Vt !=  VT_BYREF|VT_VARIANT, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_PVAR)(unsafe.Pointer(v))
	return p.pvarVal, nil
}

func (v *VARIANT) SetPVariant(value *VARIANT) {
	v.Vt = VT_BYREF | VT_VARIANT
	p := (*VAR_PVAR)(unsafe.Pointer(v))
	p.pvarVal = value
}

func (v *VARIANT) MustPBool() *VARIANT_BOOL {
	value, err := v.PBool()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) PBool() (*VARIANT_BOOL, error) {
	if v.Vt != VT_BYREF|VT_BOOL {
		return nil, fmt.Errorf("Error: PBool() v.Vt !=  VT_BYREF|VT_BOOL, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_PBOOL)(unsafe.Pointer(v))
	return p.pboolVal, nil
}

func (v *VARIANT) SetPBool(value *VARIANT_BOOL) {
	v.Vt = VT_BYREF | VT_BOOL
	p := (*VAR_PBOOL)(unsafe.Pointer(v))
	p.pboolVal = value
}

func (v *VARIANT) MustPPDispatch() **IDispatch {
	value, err := v.PPDispatch()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) PPDispatch() (**IDispatch, error) {
	if v.Vt != VT_BYREF|VT_DISPATCH {
		return nil, fmt.Errorf("PPDispatch() v.Vt !=  VT_BYREF|VT_DISPATCH, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_PPDISP)(unsafe.Pointer(v))
	return p.ppdispVal, nil
}

func (v *VARIANT) SetPPDispatch(value **IDispatch) {
	v.Vt = VT_BYREF | VT_DISPATCH
	p := (*VAR_PPDISP)(unsafe.Pointer(v))
	p.ppdispVal = value
}

func (v *VARIANT) MustPSafeArray() *SAFEARRAY {
	value, err := v.PSafeArray()
	if err != nil {
		panic(err)
	}
	return value
}

func (v *VARIANT) PSafeArray() (*SAFEARRAY, error) {
	if (v.Vt & VT_ARRAY) != VT_ARRAY {
		return nil, fmt.Errorf("Error: PSafeArray() (v.Vt & VT_ARRAY) != VT_ARRAY, ptr=%p, value=%+v", v, v)
	}
	p := (*VAR_PSAFEARRAY)(unsafe.Pointer(v))
	return p.parray, nil
}

func (v *VARIANT) SetPSafeArray(value *SAFEARRAY, elementVt VARTYPE) {
	v.Vt = VT_ARRAY | elementVt
	p := (*VAR_PSAFEARRAY)(unsafe.Pointer(v))
	p.parray = value
}
