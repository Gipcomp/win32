// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package tom

import (
	"syscall"
	"unsafe"

	"github.com/Gipcomp/winapi/kernel32"
	"github.com/Gipcomp/winapi/oaidl"
	"github.com/Gipcomp/winapi/ole32"
	"github.com/Gipcomp/winapi/oleaut32"
	"github.com/Gipcomp/winapi/win"
)

func (obj *ITextDocument) QueryInterface(riid ole32.REFIID, ppvObject *unsafe.Pointer) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.QueryInterface, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)))
	return win.HRESULT(ret)
}

func (obj *ITextDocument) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.AddRef, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *ITextDocument) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *ITextDocument) GetTypeInfoCount(pctinfo *uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetTypeInfoCount, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pctinfo)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) GetTypeInfo(iTInfo uint32, lcid kernel32.LCID, ppTInfo **oaidl.ITypeInfo) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.GetTypeInfo, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iTInfo),
		uintptr(lcid),
		uintptr(unsafe.Pointer(ppTInfo)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) GetIDsOfNames(riid ole32.REFIID, rgszNames **uint16, cNames uint32, lcid kernel32.LCID, rgDispId *oleaut32.DISPID) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.GetIDsOfNames, 6,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(rgszNames)),
		uintptr(cNames),
		uintptr(lcid),
		uintptr(unsafe.Pointer(rgDispId)))
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Invoke(dispIdMember oleaut32.DISPID, riid ole32.REFIID, lcid kernel32.LCID, wFlags uint16, pDispParams *oleaut32.DISPPARAMS, pVarResult *oleaut32.VARIANT, pExcepInfo *oaidl.EXCEPINFO, puArgErr *uint32) win.HRESULT {
	ret, _, _ := syscall.Syscall9(obj.LpVtbl.Invoke, 9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(dispIdMember),
		uintptr(unsafe.Pointer(riid)),
		uintptr(lcid),
		uintptr(wFlags),
		uintptr(unsafe.Pointer(pDispParams)),
		uintptr(unsafe.Pointer(pVarResult)),
		uintptr(unsafe.Pointer(pExcepInfo)),
		uintptr(unsafe.Pointer(puArgErr)))
	return win.HRESULT(ret)
}

func (obj *ITextDocument) GetName(pName **uint16 /*BSTR*/) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetName, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pName)),
		0)
	return win.HRESULT(ret)

}

func (obj *ITextDocument) GetSelection(ppSel **ITextSelection) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetSelection, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppSel)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) GetStoryCount(pCount *int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetStoryCount, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pCount)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) GetStoryRanges(ppStories **ITextStoryRanges) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetStoryRanges, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppStories)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) GetSaved(pValue *int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetSaved, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pValue)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) SetSaved(Value int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetSaved, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(Value),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) GetDefaultTabStop(pValue *float32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetDefaultTabStop, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pValue)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) SetDefaultTabStop(Value float32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetDefaultTabStop, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(Value),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) New() win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.New, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Open(pVar *oleaut32.VARIANT, Flags int32, CodePage int32) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.Open, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pVar)),
		uintptr(Flags),
		uintptr(CodePage),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Save(pVar *oleaut32.VARIANT, Flags int32, CodePage int32) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.Save, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pVar)),
		uintptr(Flags),
		uintptr(CodePage),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Freeze(pCount *int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Freeze, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pCount)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Unfreeze(pCount *int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Freeze, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pCount)),
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) BeginEditCollection() win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.BeginEditCollection, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) EndEditCollection() win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.EndEditCollection, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Undo(Count int32, pCount *int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Undo, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(Count),
		uintptr(unsafe.Pointer(pCount)))
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Redo(Count int32, pCount *int32) win.HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Redo, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(Count),
		uintptr(unsafe.Pointer(pCount)))
	return win.HRESULT(ret)
}

func (obj *ITextDocument) Range(cpActive int32, cpAnchor int32, ppRange **ITextRange) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.Range, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(cpActive),
		uintptr(cpAnchor),
		uintptr(unsafe.Pointer(ppRange)),
		0,
		0)
	return win.HRESULT(ret)
}

func (obj *ITextDocument) RangeFromPoint(x int32, y int32, ppRange **ITextRange) win.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.LpVtbl.RangeFromPoint, 4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(ppRange)),
		0,
		0)
	return win.HRESULT(ret)
}
