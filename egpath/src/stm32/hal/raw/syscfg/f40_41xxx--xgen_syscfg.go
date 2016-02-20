// +build f40_41xxx

package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type SYSCFG_Periph struct {
	MEMRMP MEMRMP
	PMC    PMC
	EXTICR [4]EXTICR
	_      [2]uint32
	CMPCR  CMPCR
}

func (p *SYSCFG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var SYSCFG = (*SYSCFG_Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE)))

type MEMRMP_Bits uint32

type MEMRMP struct{ mmio.U32 }

func (r *MEMRMP) Bits(mask MEMRMP_Bits) MEMRMP_Bits { return MEMRMP_Bits(r.U32.Bits(uint32(mask))) }
func (r *MEMRMP) StoreBits(mask, b MEMRMP_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MEMRMP) SetBits(mask MEMRMP_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *MEMRMP) ClearBits(mask MEMRMP_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *MEMRMP) Load() MEMRMP_Bits                 { return MEMRMP_Bits(r.U32.Load()) }
func (r *MEMRMP) Store(b MEMRMP_Bits)               { r.U32.Store(uint32(b)) }

type MEMRMP_Mask struct{ mmio.UM32 }

func (rm MEMRMP_Mask) Load() MEMRMP_Bits   { return MEMRMP_Bits(rm.UM32.Load()) }
func (rm MEMRMP_Mask) Store(b MEMRMP_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) MEM_MODE() MEMRMP_Mask {
	return MEMRMP_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MEM_MODE)}}
}

func (p *SYSCFG_Periph) FB_MODE() MEMRMP_Mask {
	return MEMRMP_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(FB_MODE)}}
}

func (p *SYSCFG_Periph) SWP_FMC() MEMRMP_Mask {
	return MEMRMP_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(SWP_FMC)}}
}

type PMC_Bits uint32

type PMC struct{ mmio.U32 }

func (r *PMC) Bits(mask PMC_Bits) PMC_Bits { return PMC_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMC) StoreBits(mask, b PMC_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMC) SetBits(mask PMC_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *PMC) ClearBits(mask PMC_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *PMC) Load() PMC_Bits              { return PMC_Bits(r.U32.Load()) }
func (r *PMC) Store(b PMC_Bits)            { r.U32.Store(uint32(b)) }

type PMC_Mask struct{ mmio.UM32 }

func (rm PMC_Mask) Load() PMC_Bits   { return PMC_Bits(rm.UM32.Load()) }
func (rm PMC_Mask) Store(b PMC_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) ADCxDC2() PMC_Mask {
	return PMC_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(ADCxDC2)}}
}

func (p *SYSCFG_Periph) MII_RMII_SEL() PMC_Mask {
	return PMC_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(MII_RMII_SEL)}}
}

type EXTICR_Bits uint32

type EXTICR struct{ mmio.U32 }

func (r *EXTICR) Bits(mask EXTICR_Bits) EXTICR_Bits { return EXTICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *EXTICR) StoreBits(mask, b EXTICR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EXTICR) SetBits(mask EXTICR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *EXTICR) ClearBits(mask EXTICR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *EXTICR) Load() EXTICR_Bits                 { return EXTICR_Bits(r.U32.Load()) }
func (r *EXTICR) Store(b EXTICR_Bits)               { r.U32.Store(uint32(b)) }

type EXTICR_Mask struct{ mmio.UM32 }

func (rm EXTICR_Mask) Load() EXTICR_Bits   { return EXTICR_Bits(rm.UM32.Load()) }
func (rm EXTICR_Mask) Store(b EXTICR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) EXTI0(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&(*[4]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8))[n], uint32(EXTI0)}}
}

func (p *SYSCFG_Periph) EXTI1(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&(*[4]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8))[n], uint32(EXTI1)}}
}

func (p *SYSCFG_Periph) EXTI2(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&(*[4]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8))[n], uint32(EXTI2)}}
}

func (p *SYSCFG_Periph) EXTI3(n int) EXTICR_Mask {
	return EXTICR_Mask{mmio.UM32{&(*[4]mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8))[n], uint32(EXTI3)}}
}

type CMPCR_Bits uint32

type CMPCR struct{ mmio.U32 }

func (r *CMPCR) Bits(mask CMPCR_Bits) CMPCR_Bits { return CMPCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMPCR) StoreBits(mask, b CMPCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMPCR) SetBits(mask CMPCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CMPCR) ClearBits(mask CMPCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CMPCR) Load() CMPCR_Bits                { return CMPCR_Bits(r.U32.Load()) }
func (r *CMPCR) Store(b CMPCR_Bits)              { r.U32.Store(uint32(b)) }

type CMPCR_Mask struct{ mmio.UM32 }

func (rm CMPCR_Mask) Load() CMPCR_Bits   { return CMPCR_Bits(rm.UM32.Load()) }
func (rm CMPCR_Mask) Store(b CMPCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) CMP_PD() CMPCR_Mask {
	return CMPCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(CMP_PD)}}
}

func (p *SYSCFG_Periph) READY() CMPCR_Mask {
	return CMPCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(READY)}}
}
