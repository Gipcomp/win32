// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package gdiplus

func (s GpStatus) String() string {
	switch s {
	case Ok:
		return "Ok"

	case GenericError:
		return "GenericError"

	case InvalidParameter:
		return "InvalidParameter"

	case OutOfMemory:
		return "OutOfMemory"

	case ObjectBusy:
		return "ObjectBusy"

	case InsufficientBuffer:
		return "InsufficientBuffer"

	case NotImplemented:
		return "NotImplemented"

	case Win32Error:
		return "Win32Error"

	case WrongState:
		return "WrongState"

	case Aborted:
		return "Aborted"

	case FileNotFound:
		return "FileNotFound"

	case ValueOverflow:
		return "ValueOverflow"

	case AccessDenied:
		return "AccessDenied"

	case UnknownImageFormat:
		return "UnknownImageFormat"

	case FontFamilyNotFound:
		return "FontFamilyNotFound"

	case FontStyleNotFound:
		return "FontStyleNotFound"

	case NotTrueTypeFont:
		return "NotTrueTypeFont"

	case UnsupportedGdiplusVersion:
		return "UnsupportedGdiplusVersion"

	case GdiplusNotInitialized:
		return "GdiplusNotInitialized"

	case PropertyNotFound:
		return "PropertyNotFound"

	case PropertyNotSupported:
		return "PropertyNotSupported"

	case ProfileNotFound:
		return "ProfileNotFound"
	}

	return "Unknown Status Value"
}
