// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package ole32

import (
	"syscall"

	"github.com/gfphoenix/win32/handle"
	"github.com/gfphoenix/win32/user32"
	"github.com/gfphoenix/win32/win"
)

type IID syscall.GUID
type CLSID syscall.GUID
type REFIID *IID
type REFCLSID *CLSID

type IClassFactoryVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	CreateInstance uintptr
	LockServer     uintptr
}

type IClassFactory struct {
	LpVtbl *IClassFactoryVtbl
}

type IConnectionPointVtbl struct {
	QueryInterface              uintptr
	AddRef                      uintptr
	Release                     uintptr
	GetConnectionInterface      uintptr
	GetConnectionPointContainer uintptr
	Advise                      uintptr
	Unadvise                    uintptr
	EnumConnections             uintptr
}

type IConnectionPoint struct {
	LpVtbl *IConnectionPointVtbl
}

type IConnectionPointContainerVtbl struct {
	QueryInterface       uintptr
	AddRef               uintptr
	Release              uintptr
	EnumConnectionPoints uintptr
	FindConnectionPoint  uintptr
}

type IConnectionPointContainer struct {
	LpVtbl *IConnectionPointContainerVtbl
}

type IOleClientSiteVtbl struct {
	QueryInterface         uintptr
	AddRef                 uintptr
	Release                uintptr
	SaveObject             uintptr
	GetMoniker             uintptr
	GetContainer           uintptr
	ShowObject             uintptr
	OnShowWindow           uintptr
	RequestNewObjectLayout uintptr
}

type IOleClientSite struct {
	LpVtbl *IOleClientSiteVtbl
}

type IOleInPlaceFrameVtbl struct {
	QueryInterface       uintptr
	AddRef               uintptr
	Release              uintptr
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
	GetBorder            uintptr
	RequestBorderSpace   uintptr
	SetBorderSpace       uintptr
	SetActiveObject      uintptr
	InsertMenus          uintptr
	SetMenu              uintptr
	RemoveMenus          uintptr
	SetStatusText        uintptr
	EnableModeless       uintptr
	TranslateAccelerator uintptr
}

type IOleInPlaceFrame struct {
	LpVtbl *IOleInPlaceFrameVtbl
}

type IOleInPlaceObjectVtbl struct {
	QueryInterface       uintptr
	AddRef               uintptr
	Release              uintptr
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
	InPlaceDeactivate    uintptr
	UIDeactivate         uintptr
	SetObjectRects       uintptr
	ReactivateAndUndo    uintptr
}

type IOleInPlaceObject struct {
	LpVtbl *IOleInPlaceObjectVtbl
}

type IOleInPlaceSiteVtbl struct {
	QueryInterface       uintptr
	AddRef               uintptr
	Release              uintptr
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
	CanInPlaceActivate   uintptr
	OnInPlaceActivate    uintptr
	OnUIActivate         uintptr
	GetWindowContext     uintptr
	Scroll               uintptr
	OnUIDeactivate       uintptr
	OnInPlaceDeactivate  uintptr
	DiscardUndoState     uintptr
	DeactivateAndUndo    uintptr
	OnPosRectChange      uintptr
}

type IOleInPlaceSite struct {
	LpVtbl *IOleInPlaceSiteVtbl
}

type IOleObjectVtbl struct {
	QueryInterface   uintptr
	AddRef           uintptr
	Release          uintptr
	SetClientSite    uintptr
	GetClientSite    uintptr
	SetHostNames     uintptr
	Close            uintptr
	SetMoniker       uintptr
	GetMoniker       uintptr
	InitFromData     uintptr
	GetClipboardData uintptr
	DoVerb           uintptr
	EnumVerbs        uintptr
	Update           uintptr
	IsUpToDate       uintptr
	GetUserClassID   uintptr
	GetUserType      uintptr
	SetExtent        uintptr
	GetExtent        uintptr
	Advise           uintptr
	Unadvise         uintptr
	EnumAdvise       uintptr
	GetMiscStatus    uintptr
	SetColorScheme   uintptr
}

type IOleObject struct {
	LpVtbl *IOleObjectVtbl
}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type IUnknown struct {
	LpVtbl *IUnknownVtbl
}

type OLEINPLACEFRAMEINFO struct {
	Cb            uint32
	FMDIApp       win.BOOL
	HwndFrame     handle.HWND
	Haccel        user32.HACCEL
	CAccelEntries uint32
}

type COAUTHIDENTITY struct {
	User           *uint16
	UserLength     uint32
	Domain         *uint16
	DomainLength   uint32
	Password       *uint16
	PasswordLength uint32
	Flags          uint32
}

type COAUTHINFO struct {
	dwAuthnSvc           uint32
	dwAuthzSvc           uint32
	pwszServerPrincName  *uint16
	dwAuthnLevel         uint32
	dwImpersonationLevel uint32
	pAuthIdentityData    *COAUTHIDENTITY
	dwCapabilities       uint32
}

type COSERVERINFO struct {
	dwReserved1 uint32
	pwszName    *uint16
	pAuthInfo   *COAUTHINFO
	dwReserved2 uint32
}
