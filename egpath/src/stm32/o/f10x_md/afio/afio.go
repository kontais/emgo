// Peripheral: AFIO_Periph  Alternate Function I/O.
// Instances:
//  AFIO  mmap.AFIO_BASE
// Registers:
//  0x00 32  EVCR
//  0x04 32  MAPR
//  0x08 32  EXTICR[4]
//  0x1C 32  MAPR2
// Import:
//  stm32/o/f10x_md/mmap
package afio

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PIN      EVCR_Bits = 0x0F << 0 //+ PIN[3:0] bits (Pin selection).
	PIN_0    EVCR_Bits = 0x01 << 0 //  Bit 0.
	PIN_1    EVCR_Bits = 0x02 << 0 //  Bit 1.
	PIN_2    EVCR_Bits = 0x04 << 0 //  Bit 2.
	PIN_3    EVCR_Bits = 0x08 << 0 //  Bit 3.
	PIN_PX0  EVCR_Bits = 0x00 << 0 //  Pin 0 selected.
	PIN_PX1  EVCR_Bits = 0x01 << 0 //  Pin 1 selected.
	PIN_PX2  EVCR_Bits = 0x02 << 0 //  Pin 2 selected.
	PIN_PX3  EVCR_Bits = 0x03 << 0 //  Pin 3 selected.
	PIN_PX4  EVCR_Bits = 0x04 << 0 //  Pin 4 selected.
	PIN_PX5  EVCR_Bits = 0x05 << 0 //  Pin 5 selected.
	PIN_PX6  EVCR_Bits = 0x06 << 0 //  Pin 6 selected.
	PIN_PX7  EVCR_Bits = 0x07 << 0 //  Pin 7 selected.
	PIN_PX8  EVCR_Bits = 0x08 << 0 //  Pin 8 selected.
	PIN_PX9  EVCR_Bits = 0x09 << 0 //  Pin 9 selected.
	PIN_PX10 EVCR_Bits = 0x0A << 0 //  Pin 10 selected.
	PIN_PX11 EVCR_Bits = 0x0B << 0 //  Pin 11 selected.
	PIN_PX12 EVCR_Bits = 0x0C << 0 //  Pin 12 selected.
	PIN_PX13 EVCR_Bits = 0x0D << 0 //  Pin 13 selected.
	PIN_PX14 EVCR_Bits = 0x0E << 0 //  Pin 14 selected.
	PIN_PX15 EVCR_Bits = 0x0F << 0 //  Pin 15 selected.
	PORT     EVCR_Bits = 0x07 << 4 //+ PORT[2:0] bits (Port selection).
	PORT_0   EVCR_Bits = 0x01 << 4 //  Bit 0.
	PORT_1   EVCR_Bits = 0x02 << 4 //  Bit 1.
	PORT_2   EVCR_Bits = 0x04 << 4 //  Bit 2.
	PORT_PA  EVCR_Bits = 0x00 << 4 //  Port A selected.
	PORT_PB  EVCR_Bits = 0x01 << 4 //  Port B selected.
	PORT_PC  EVCR_Bits = 0x02 << 4 //  Port C selected.
	PORT_PD  EVCR_Bits = 0x03 << 4 //  Port D selected.
	PORT_PE  EVCR_Bits = 0x04 << 4 //  Port E selected.
	EVOE     EVCR_Bits = 0x01 << 7 //+ Event Output Enable.
)

const (
	PINn  = 0
	PORTn = 4
	EVOEn = 7
)

const (
	SPI1_REMAP                MAPR_Bits = 0x01 << 0  //+ SPI1 remapping.
	I2C1_REMAP                MAPR_Bits = 0x01 << 1  //+ I2C1 remapping.
	USART1_REMAP              MAPR_Bits = 0x01 << 2  //+ USART1 remapping.
	USART2_REMAP              MAPR_Bits = 0x01 << 3  //+ USART2 remapping.
	USART3_REMAP              MAPR_Bits = 0x03 << 4  //+ USART3_REMAP[1:0] bits (USART3 remapping).
	USART3_REMAP_0            MAPR_Bits = 0x01 << 4  //  Bit 0.
	USART3_REMAP_1            MAPR_Bits = 0x02 << 4  //  Bit 1.
	USART3_REMAP_NOREMAP      MAPR_Bits = 0x00 << 4  //  No remap (TX/PB10, RX/PB11, CK/PB12, CTS/PB13, RTS/PB14).
	USART3_REMAP_PARTIALREMAP MAPR_Bits = 0x01 << 4  //  Partial remap (TX/PC10, RX/PC11, CK/PC12, CTS/PB13, RTS/PB14).
	USART3_REMAP_FULLREMAP    MAPR_Bits = 0x03 << 4  //  Full remap (TX/PD8, RX/PD9, CK/PD10, CTS/PD11, RTS/PD12).
	TIM1_REMAP                MAPR_Bits = 0x03 << 6  //+ TIM1_REMAP[1:0] bits (TIM1 remapping).
	TIM1_REMAP_0              MAPR_Bits = 0x01 << 6  //  Bit 0.
	TIM1_REMAP_1              MAPR_Bits = 0x02 << 6  //  Bit 1.
	TIM1_REMAP_NOREMAP        MAPR_Bits = 0x00 << 6  //  No remap (ETR/PA12, CH1/PA8, CH2/PA9, CH3/PA10, CH4/PA11, BKIN/PB12, CH1N/PB13, CH2N/PB14, CH3N/PB15).
	TIM1_REMAP_PARTIALREMAP   MAPR_Bits = 0x01 << 6  //  Partial remap (ETR/PA12, CH1/PA8, CH2/PA9, CH3/PA10, CH4/PA11, BKIN/PA6, CH1N/PA7, CH2N/PB0, CH3N/PB1).
	TIM1_REMAP_FULLREMAP      MAPR_Bits = 0x03 << 6  //  Full remap (ETR/PE7, CH1/PE9, CH2/PE11, CH3/PE13, CH4/PE14, BKIN/PE15, CH1N/PE8, CH2N/PE10, CH3N/PE12).
	TIM2_REMAP                MAPR_Bits = 0x03 << 8  //+ TIM2_REMAP[1:0] bits (TIM2 remapping).
	TIM2_REMAP_0              MAPR_Bits = 0x01 << 8  //  Bit 0.
	TIM2_REMAP_1              MAPR_Bits = 0x02 << 8  //  Bit 1.
	TIM2_REMAP_NOREMAP        MAPR_Bits = 0x00 << 8  //  No remap (CH1/ETR/PA0, CH2/PA1, CH3/PA2, CH4/PA3).
	TIM2_REMAP_PARTIALREMAP1  MAPR_Bits = 0x01 << 8  //  Partial remap (CH1/ETR/PA15, CH2/PB3, CH3/PA2, CH4/PA3).
	TIM2_REMAP_PARTIALREMAP2  MAPR_Bits = 0x02 << 8  //  Partial remap (CH1/ETR/PA0, CH2/PA1, CH3/PB10, CH4/PB11).
	TIM2_REMAP_FULLREMAP      MAPR_Bits = 0x03 << 8  //  Full remap (CH1/ETR/PA15, CH2/PB3, CH3/PB10, CH4/PB11).
	TIM3_REMAP                MAPR_Bits = 0x03 << 10 //+ TIM3_REMAP[1:0] bits (TIM3 remapping).
	TIM3_REMAP_0              MAPR_Bits = 0x01 << 10 //  Bit 0.
	TIM3_REMAP_1              MAPR_Bits = 0x02 << 10 //  Bit 1.
	TIM3_REMAP_NOREMAP        MAPR_Bits = 0x00 << 10 //  No remap (CH1/PA6, CH2/PA7, CH3/PB0, CH4/PB1).
	TIM3_REMAP_PARTIALREMAP   MAPR_Bits = 0x02 << 10 //  Partial remap (CH1/PB4, CH2/PB5, CH3/PB0, CH4/PB1).
	TIM3_REMAP_FULLREMAP      MAPR_Bits = 0x03 << 10 //  Full remap (CH1/PC6, CH2/PC7, CH3/PC8, CH4/PC9).
	TIM4_REMAP                MAPR_Bits = 0x01 << 12 //+ TIM4_REMAP bit (TIM4 remapping).
	CAN_REMAP                 MAPR_Bits = 0x03 << 13 //+ CAN_REMAP[1:0] bits (CAN Alternate function remapping).
	CAN_REMAP_0               MAPR_Bits = 0x01 << 13 //  Bit 0.
	CAN_REMAP_1               MAPR_Bits = 0x02 << 13 //  Bit 1.
	CAN_REMAP_REMAP1          MAPR_Bits = 0x00 << 13 //  CANRX mapped to PA11, CANTX mapped to PA12.
	CAN_REMAP_REMAP2          MAPR_Bits = 0x02 << 13 //  CANRX mapped to PB8, CANTX mapped to PB9.
	CAN_REMAP_REMAP3          MAPR_Bits = 0x03 << 13 //  CANRX mapped to PD0, CANTX mapped to PD1.
	PD01_REMAP                MAPR_Bits = 0x01 << 15 //+ Port D0/Port D1 mapping on OSC_IN/OSC_OUT.
	TIM5CH4_IREMAP            MAPR_Bits = 0x01 << 16 //+ TIM5 Channel4 Internal Remap.
	ADC1_ETRGINJ_REMAP        MAPR_Bits = 0x01 << 17 //+ ADC 1 External Trigger Injected Conversion remapping.
	ADC1_ETRGREG_REMAP        MAPR_Bits = 0x01 << 18 //+ ADC 1 External Trigger Regular Conversion remapping.
	ADC2_ETRGINJ_REMAP        MAPR_Bits = 0x01 << 19 //+ ADC 2 External Trigger Injected Conversion remapping.
	ADC2_ETRGREG_REMAP        MAPR_Bits = 0x01 << 20 //+ ADC 2 External Trigger Regular Conversion remapping.
	SWJ_CFG                   MAPR_Bits = 0x07 << 24 //+ SWJ_CFG[2:0] bits (Serial Wire JTAG configuration).
	SWJ_CFG_0                 MAPR_Bits = 0x01 << 24 //  Bit 0.
	SWJ_CFG_1                 MAPR_Bits = 0x02 << 24 //  Bit 1.
	SWJ_CFG_2                 MAPR_Bits = 0x04 << 24 //  Bit 2.
	SWJ_CFG_RESET             MAPR_Bits = 0x00 << 24 //  Full SWJ (JTAG-DP + SW-DP) : Reset State.
	SWJ_CFG_NOJNTRST          MAPR_Bits = 0x01 << 24 //  Full SWJ (JTAG-DP + SW-DP) but without JNTRST.
	SWJ_CFG_JTAGDISABLE       MAPR_Bits = 0x02 << 24 //  JTAG-DP Disabled and SW-DP Enabled.
	SWJ_CFG_DISABLE           MAPR_Bits = 0x04 << 24 //  JTAG-DP Disabled and SW-DP Disabled.
)

const (
	SPI1_REMAPn         = 0
	I2C1_REMAPn         = 1
	USART1_REMAPn       = 2
	USART2_REMAPn       = 3
	USART3_REMAPn       = 4
	TIM1_REMAPn         = 6
	TIM2_REMAPn         = 8
	TIM3_REMAPn         = 10
	TIM4_REMAPn         = 12
	CAN_REMAPn          = 13
	PD01_REMAPn         = 15
	TIM5CH4_IREMAPn     = 16
	ADC1_ETRGINJ_REMAPn = 17
	ADC1_ETRGREG_REMAPn = 18
	ADC2_ETRGINJ_REMAPn = 19
	ADC2_ETRGREG_REMAPn = 20
	SWJ_CFGn            = 24
)

const (
	EXTI0    EXTICR_Bits = 0x0F << 0  //+ EXTI 0 configuration.
	EXTI1    EXTICR_Bits = 0x0F << 4  //+ EXTI 1 configuration.
	EXTI2    EXTICR_Bits = 0x0F << 8  //+ EXTI 2 configuration.
	EXTI3    EXTICR_Bits = 0x0F << 12 //+ EXTI 3 configuration.
	EXTI0_PA EXTICR_Bits = 0x00 << 12 //  PA[0] pin.
	EXTI0_PB EXTICR_Bits = 0x01 << 0  //  PB[0] pin.
	EXTI0_PC EXTICR_Bits = 0x02 << 0  //  PC[0] pin.
	EXTI0_PD EXTICR_Bits = 0x03 << 0  //  PD[0] pin.
	EXTI0_PE EXTICR_Bits = 0x04 << 0  //  PE[0] pin.
	EXTI0_PF EXTICR_Bits = 0x05 << 0  //  PF[0] pin.
	EXTI0_PG EXTICR_Bits = 0x06 << 0  //  PG[0] pin.
	EXTI1_PA EXTICR_Bits = 0x00 << 12 //  PA[1] pin.
	EXTI1_PB EXTICR_Bits = 0x01 << 4  //  PB[1] pin.
	EXTI1_PC EXTICR_Bits = 0x02 << 4  //  PC[1] pin.
	EXTI1_PD EXTICR_Bits = 0x03 << 4  //  PD[1] pin.
	EXTI1_PE EXTICR_Bits = 0x04 << 4  //  PE[1] pin.
	EXTI1_PF EXTICR_Bits = 0x05 << 4  //  PF[1] pin.
	EXTI1_PG EXTICR_Bits = 0x06 << 4  //  PG[1] pin.
	EXTI2_PA EXTICR_Bits = 0x00 << 12 //  PA[2] pin.
	EXTI2_PB EXTICR_Bits = 0x01 << 8  //  PB[2] pin.
	EXTI2_PC EXTICR_Bits = 0x02 << 8  //  PC[2] pin.
	EXTI2_PD EXTICR_Bits = 0x03 << 8  //  PD[2] pin.
	EXTI2_PE EXTICR_Bits = 0x04 << 8  //  PE[2] pin.
	EXTI2_PF EXTICR_Bits = 0x05 << 8  //  PF[2] pin.
	EXTI2_PG EXTICR_Bits = 0x06 << 8  //  PG[2] pin.
	EXTI3_PA EXTICR_Bits = 0x00 << 12 //  PA[3] pin.
	EXTI3_PB EXTICR_Bits = 0x01 << 12 //  PB[3] pin.
	EXTI3_PC EXTICR_Bits = 0x02 << 12 //  PC[3] pin.
	EXTI3_PD EXTICR_Bits = 0x03 << 12 //  PD[3] pin.
	EXTI3_PE EXTICR_Bits = 0x04 << 12 //  PE[3] pin.
	EXTI3_PF EXTICR_Bits = 0x05 << 12 //  PF[3] pin.
	EXTI3_PG EXTICR_Bits = 0x06 << 12 //  PG[3] pin.
)

const (
	EXTI0n = 0
	EXTI1n = 4
	EXTI2n = 8
	EXTI3n = 12
)
