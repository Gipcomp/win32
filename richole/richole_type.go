// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package richole

import (
	"github.com/D4v1dW3bb/winapi/gdi32"
	"github.com/D4v1dW3bb/winapi/objidl"
	"github.com/D4v1dW3bb/winapi/ole32"
)

type REOBJECT struct {
	cbStruct uint32                // Size of structure
	cp       int32                 // Character position of object
	clsid    ole32.CLSID           // Class ID of object
	poleobj  *ole32.IOleObject     // OLE object interface
	pstg     *objidl.IStorage      // Associated storage interface
	polesite *ole32.IOleClientSite // Associated client site interface
	sizel    gdi32.SIZE            // Size of object (may be 0,0)
	dvaspect uint32                // Display aspect to use
	dwFlags  uint32                // Object status flags
	dwUser   uint32                // Dword for user's use
}

type IRichEditOleVtbl struct {
	ole32.IUnknownVtbl
	GetClientSite        uintptr
	GetObjectCount       uintptr
	GetLinkCount         uintptr
	GetObject            uintptr
	InsertObject         uintptr
	ConvertObject        uintptr
	ActivateAs           uintptr
	SetHostNames         uintptr
	SetLinkAvailable     uintptr
	SetDvaspect          uintptr
	HandsOffStorage      uintptr
	SaveCompleted        uintptr
	InPlaceDeactivate    uintptr
	ContextSensitiveHelp uintptr
	GetClipboardData     uintptr
	ImportDataObject     uintptr
}

type IRichEditOle struct {
	LpVtbl *IRichEditOleVtbl
}
