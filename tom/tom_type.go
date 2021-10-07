// Copyright 2011 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package tom

import "github.com/Gipcomp/win32/oleaut32"

type TomConstants uint32

type OBJECTTYPE int32

type ITextRangeVtbl struct {
	oleaut32.IDispatchVtbl
	GetText           uintptr
	SetText           uintptr
	GetChar           uintptr
	SetChar           uintptr
	GetDuplicate      uintptr
	GetFormattedText  uintptr
	SetFormattedText  uintptr
	GetStart          uintptr
	SetStart          uintptr
	GetEnd            uintptr
	SetEnd            uintptr
	GetFont           uintptr
	SetFont           uintptr
	GetPara           uintptr
	SetPara           uintptr
	GetStoryLength    uintptr
	GetStoryType      uintptr
	Collapse          uintptr
	Expand            uintptr
	GetIndex          uintptr
	SetIndex          uintptr
	SetRange          uintptr
	InRange           uintptr
	InStory           uintptr
	IsEqual           uintptr
	Select            uintptr
	StartOf           uintptr
	EndOf             uintptr
	Move              uintptr
	MoveStart         uintptr
	MoveEnd           uintptr
	MoveWhile         uintptr
	MoveStartWhile    uintptr
	MoveEndWhile      uintptr
	MoveUntil         uintptr
	MoveStartUntil    uintptr
	MoveEndUntil      uintptr
	FindText          uintptr
	FindTextStart     uintptr
	FindTextEnd       uintptr
	Delete            uintptr
	Cut               uintptr
	Copy              uintptr
	Paste             uintptr
	CanPaste          uintptr
	CanEdit           uintptr
	ChangeCase        uintptr
	GetPoint          uintptr
	SetPoint          uintptr
	ScrollIntoView    uintptr
	GetEmbeddedObject uintptr
}

type ITextRange struct {
	LpVtbl *ITextRangeVtbl
}

type ITextSelectionVtbl struct {
	ITextRangeVtbl
	GetFlags  uintptr
	SetFlags  uintptr
	GetType   uintptr
	MoveLeft  uintptr
	MoveRight uintptr
	MoveUp    uintptr
	MoveDown  uintptr
	HomeKey   uintptr
	EndKey    uintptr
	TypeText  uintptr
}

type ITextSelection struct {
	LpVtbl *ITextSelectionVtbl
}

type ITextDocumentVtbl struct {
	oleaut32.IDispatchVtbl
	GetName             uintptr
	GetSelection        uintptr
	GetStoryCount       uintptr
	GetStoryRanges      uintptr
	GetSaved            uintptr
	SetSaved            uintptr
	GetDefaultTabStop   uintptr
	SetDefaultTabStop   uintptr
	New                 uintptr
	Open                uintptr
	Save                uintptr
	Freeze              uintptr
	Unfreeze            uintptr
	BeginEditCollection uintptr
	EndEditCollection   uintptr
	Undo                uintptr
	Redo                uintptr
	Range               uintptr
	RangeFromPoint      uintptr
}

type ITextStoryRangesVtbl struct {
	oleaut32.IDispatchVtbl
	NewEnum  uintptr
	Item     uintptr
	GetCount uintptr
}

type ITextStoryRanges struct {
	LpVtbl *ITextStoryRangesVtbl
}

type ITextDocument struct {
	LpVtbl *ITextDocumentVtbl
}
