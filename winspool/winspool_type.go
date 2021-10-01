// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winspool

type PRINTER_INFO_4 struct {
	PPrinterName *uint16
	PServerName  *uint16
	Attributes   uint32
}
