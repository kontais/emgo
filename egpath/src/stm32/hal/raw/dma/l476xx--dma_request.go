// +build l476xx

// Peripheral: DMA_Request_Periph  DMA Controller.
// Instances:
//  DMA1_Request  mmap.DMA1_CSELR_BASE
//  DMA2_Request  mmap.DMA2_CSELR_BASE
// Registers:
//  0x00 32  CSELR DMA channel selection register.
// Import:
//  stm32/o/l476xx/mmap
package dma

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	C1S CSELR_Bits = 0x0F << 0  //+ Channel 1 Selection.
	C2S CSELR_Bits = 0x0F << 4  //+ Channel 2 Selection.
	C3S CSELR_Bits = 0x0F << 8  //+ Channel 3 Selection.
	C4S CSELR_Bits = 0x0F << 12 //+ Channel 4 Selection.
	C5S CSELR_Bits = 0x0F << 16 //+ Channel 5 Selection.
	C6S CSELR_Bits = 0x0F << 20 //+ Channel 6 Selection.
	C7S CSELR_Bits = 0x0F << 24 //+ Channel 7 Selection.
)

const (
	C1Sn = 0
	C2Sn = 4
	C3Sn = 8
	C4Sn = 12
	C5Sn = 16
	C6Sn = 20
	C7Sn = 24
)
