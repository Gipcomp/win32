// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package shdocvw

type DWebBrowserEvents2Vtbl struct {
	QueryInterface   uintptr
	AddRef           uintptr
	Release          uintptr
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

type DWebBrowserEvents2 struct {
	LpVtbl *DWebBrowserEvents2Vtbl
}

type IWebBrowser2Vtbl struct {
	QueryInterface           uintptr
	AddRef                   uintptr
	Release                  uintptr
	GetTypeInfoCount         uintptr
	GetTypeInfo              uintptr
	GetIDsOfNames            uintptr
	Invoke                   uintptr
	GoBack                   uintptr
	GoForward                uintptr
	GoHome                   uintptr
	GoSearch                 uintptr
	Navigate                 uintptr
	Refresh                  uintptr
	Refresh2                 uintptr
	Stop                     uintptr
	Get_Application          uintptr
	Get_Parent               uintptr
	Get_Container            uintptr
	Get_Document             uintptr
	Get_TopLevelContainer    uintptr
	Get_Type                 uintptr
	Get_Left                 uintptr
	Put_Left                 uintptr
	Get_Top                  uintptr
	Put_Top                  uintptr
	Get_Width                uintptr
	Put_Width                uintptr
	Get_Height               uintptr
	Put_Height               uintptr
	Get_LocationName         uintptr
	Get_LocationURL          uintptr
	Get_Busy                 uintptr
	Quit                     uintptr
	ClientToWindow           uintptr
	PutProperty              uintptr
	GetProperty              uintptr
	Get_Name                 uintptr
	Get_HWND                 uintptr
	Get_FullName             uintptr
	Get_Path                 uintptr
	Get_Visible              uintptr
	Put_Visible              uintptr
	Get_StatusBar            uintptr
	Put_StatusBar            uintptr
	Get_StatusText           uintptr
	Put_StatusText           uintptr
	Get_ToolBar              uintptr
	Put_ToolBar              uintptr
	Get_MenuBar              uintptr
	Put_MenuBar              uintptr
	Get_FullScreen           uintptr
	Put_FullScreen           uintptr
	Navigate2                uintptr
	QueryStatusWB            uintptr
	ExecWB                   uintptr
	ShowBrowserBar           uintptr
	Get_ReadyState           uintptr
	Get_Offline              uintptr
	Put_Offline              uintptr
	Get_Silent               uintptr
	Put_Silent               uintptr
	Get_RegisterAsBrowser    uintptr
	Put_RegisterAsBrowser    uintptr
	Get_RegisterAsDropTarget uintptr
	Put_RegisterAsDropTarget uintptr
	Get_TheaterMode          uintptr
	Put_TheaterMode          uintptr
	Get_AddressBar           uintptr
	Put_AddressBar           uintptr
	Get_Resizable            uintptr
	Put_Resizable            uintptr
}

type IWebBrowser2 struct {
	LpVtbl *IWebBrowser2Vtbl
}

type IDocHostUIHandlerVtbl struct {
	QueryInterface        uintptr
	AddRef                uintptr
	Release               uintptr
	ShowContextMenu       uintptr
	GetHostInfo           uintptr
	ShowUI                uintptr
	HideUI                uintptr
	UpdateUI              uintptr
	EnableModeless        uintptr
	OnDocWindowActivate   uintptr
	OnFrameWindowActivate uintptr
	ResizeBorder          uintptr
	TranslateAccelerator  uintptr
	GetOptionKeyPath      uintptr
	GetDropTarget         uintptr
	GetExternal           uintptr
	TranslateUrl          uintptr
	FilterDataObject      uintptr
}

type IDocHostUIHandler struct {
	LpVtbl *IDocHostUIHandlerVtbl
}

type DOCHOSTUIINFO struct {
	CbSize        uint32
	DwFlags       uint32
	DwDoubleClick uint32
	PchHostCss    *uint16
	PchHostNS     *uint16
}

type IOleInPlaceActiveObjectVtbl struct {
	QueryInterface        uintptr
	AddRef                uintptr
	Release               uintptr
	GetWindow             uintptr
	ContextSensitiveHelp  uintptr
	TranslateAccelerator  uintptr
	OnFrameWindowActivate uintptr
	OnDocWindowActivate   uintptr
	ResizeBorder          uintptr
	EnableModeless        uintptr
}

type IOleInPlaceActiveObject struct {
	LpVtbl *IOleInPlaceActiveObjectVtbl
}
