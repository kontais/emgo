// +build f40_41xxx

package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type FSMC_Bank2_Periph struct {
	PCR2  PCR2
	SR2   SR2
	PMEM2 PMEM2
	PATT2 PATT2
	_     uint32
	ECCR2 ECCR2
}

func (p *FSMC_Bank2_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var FSMC_Bank2 = (*FSMC_Bank2_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank2_R_BASE)))

type PCR2_Bits uint32

type PCR2 struct{ mmio.U32 }

func (r *PCR2) Bits(mask PCR2_Bits) PCR2_Bits { return PCR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *PCR2) StoreBits(mask, b PCR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCR2) SetBits(mask PCR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PCR2) ClearBits(mask PCR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PCR2) Load() PCR2_Bits               { return PCR2_Bits(r.U32.Load()) }
func (r *PCR2) Store(b PCR2_Bits)             { r.U32.Store(uint32(b)) }

type PCR2_Mask struct{ mmio.UM32 }

func (rm PCR2_Mask) Load() PCR2_Bits   { return PCR2_Bits(rm.UM32.Load()) }
func (rm PCR2_Mask) Store(b PCR2_Bits) { rm.UM32.Store(uint32(b)) }

type SR2_Bits uint32

type SR2 struct{ mmio.U32 }

func (r *SR2) Bits(mask SR2_Bits) SR2_Bits { return SR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR2) StoreBits(mask, b SR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR2) SetBits(mask SR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SR2) ClearBits(mask SR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SR2) Load() SR2_Bits              { return SR2_Bits(r.U32.Load()) }
func (r *SR2) Store(b SR2_Bits)            { r.U32.Store(uint32(b)) }

type SR2_Mask struct{ mmio.UM32 }

func (rm SR2_Mask) Load() SR2_Bits   { return SR2_Bits(rm.UM32.Load()) }
func (rm SR2_Mask) Store(b SR2_Bits) { rm.UM32.Store(uint32(b)) }

type PMEM2_Bits uint32

type PMEM2 struct{ mmio.U32 }

func (r *PMEM2) Bits(mask PMEM2_Bits) PMEM2_Bits { return PMEM2_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMEM2) StoreBits(mask, b PMEM2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMEM2) SetBits(mask PMEM2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PMEM2) ClearBits(mask PMEM2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PMEM2) Load() PMEM2_Bits                { return PMEM2_Bits(r.U32.Load()) }
func (r *PMEM2) Store(b PMEM2_Bits)              { r.U32.Store(uint32(b)) }

type PMEM2_Mask struct{ mmio.UM32 }

func (rm PMEM2_Mask) Load() PMEM2_Bits   { return PMEM2_Bits(rm.UM32.Load()) }
func (rm PMEM2_Mask) Store(b PMEM2_Bits) { rm.UM32.Store(uint32(b)) }

type PATT2_Bits uint32

type PATT2 struct{ mmio.U32 }

func (r *PATT2) Bits(mask PATT2_Bits) PATT2_Bits { return PATT2_Bits(r.U32.Bits(uint32(mask))) }
func (r *PATT2) StoreBits(mask, b PATT2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PATT2) SetBits(mask PATT2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PATT2) ClearBits(mask PATT2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PATT2) Load() PATT2_Bits                { return PATT2_Bits(r.U32.Load()) }
func (r *PATT2) Store(b PATT2_Bits)              { r.U32.Store(uint32(b)) }

type PATT2_Mask struct{ mmio.UM32 }

func (rm PATT2_Mask) Load() PATT2_Bits   { return PATT2_Bits(rm.UM32.Load()) }
func (rm PATT2_Mask) Store(b PATT2_Bits) { rm.UM32.Store(uint32(b)) }

type ECCR2_Bits uint32

type ECCR2 struct{ mmio.U32 }

func (r *ECCR2) Bits(mask ECCR2_Bits) ECCR2_Bits { return ECCR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *ECCR2) StoreBits(mask, b ECCR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ECCR2) SetBits(mask ECCR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ECCR2) ClearBits(mask ECCR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ECCR2) Load() ECCR2_Bits                { return ECCR2_Bits(r.U32.Load()) }
func (r *ECCR2) Store(b ECCR2_Bits)              { r.U32.Store(uint32(b)) }

type ECCR2_Mask struct{ mmio.UM32 }

func (rm ECCR2_Mask) Load() ECCR2_Bits   { return ECCR2_Bits(rm.UM32.Load()) }
func (rm ECCR2_Mask) Store(b ECCR2_Bits) { rm.UM32.Store(uint32(b)) }
