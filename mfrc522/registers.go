// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package commands contains the command that a MFRC522 supports.
package mfrc522

// Register constants
const (
	CommandReg    = 0x01
	CommIEnReg    = 0x02
	DivlEnReg     = 0x03
	CommIrqReg    = 0x04
	DivIrqReg     = 0x05
	ErrorReg      = 0x06
	Status1Reg    = 0x07
	Status2Reg    = 0x08
	FIFODataReg   = 0x09
	FIFOLevelReg  = 0x0A
	WaterLevelReg = 0x0B
	ControlReg    = 0x0C
	BitFramingReg = 0x0D
	CollReg       = 0x0E

	ModeReg        = 0x11
	TxModeReg      = 0x12
	RxModeReg      = 0x13
	TxControlReg   = 0x14
	TxAutoReg      = 0x15
	TxSelReg       = 0x16
	RxSelReg       = 0x17
	RxThresholdReg = 0x18
	DemodReg       = 0x19
	MifareReg      = 0x1C
	SerialSpeedReg = 0x1F

	CRCResultRegM     = 0x21
	CRCResultRegL     = 0x22
	ModWidthReg       = 0x24
	RFCfgReg          = 0x26
	GsNReg            = 0x27
	CWGsPReg          = 0x28
	ModGsPReg         = 0x29
	TModeReg          = 0x2A
	TPrescalerReg     = 0x2B
	TReloadRegH       = 0x2C
	TReloadRegL       = 0x2D
	TCounterValueRegH = 0x2E
	TCounterValueRegL = 0x2F

	TestSel1Reg     = 0x31
	TestSel2Reg     = 0x32
	TestPinEnReg    = 0x33
	TestPinValueReg = 0x34
	TestBusReg      = 0x35
	AutoTestReg     = 0x36
	VersionReg      = 0x37
	AnalogTestReg   = 0x38
	TestDAC1Reg     = 0x39
	TestDAC2Reg     = 0x3A
	TestADCReg      = 0x3B
)

// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Operation constants
const (
	PCD_IDLE       = 0x00
	PCD_AUTHENT    = 0x0E
	PCD_RECEIVE    = 0x08
	PCD_TRANSMIT   = 0x04
	PCD_TRANSCEIVE = 0x0C
	PCD_RESETPHASE = 0x0F
	PCD_CALCCRC    = 0x03

	PICC_REQIDL    = 0x26
	PICC_REQALL    = 0x52
	PICC_ANTICOLL  = 0x93
	PICC_ANTICOLL2 = 0x95
	PICC_SElECTTAG = 0x93
	PICC_AUTHENT1A = 0x60
	PICC_AUTHENT1B = 0x61
	PICC_READ      = 0x30
	PICC_WRITE     = 0xA0
	PICC_DECREMENT = 0xC0
	PICC_INCREMENT = 0xC1
	PICC_RESTORE   = 0xC2
	PICC_TRANSFER  = 0xB0
	PICC_HALT      = 0x50
)
