package hrtim

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type HRTIM_Periph struct {
}

func (p *HRTIM_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var HRTIM1 = (*HRTIM_Periph)(unsafe.Pointer(uintptr(mmap.HRTIM1_BASE)))