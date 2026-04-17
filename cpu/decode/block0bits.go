package decode

import "github.com/AlanRostem/scratchboy/nums"

const (
	b0OperandMaskNone       = nums.Byte(0b00000000)
	b0OperandMaskOp4Bits    = nums.Byte(0b00110000)
	b0OperandMaskOp3Bits    = nums.Byte(0b00111000)
	b0OperandMaskZeros3Bits = nums.Byte(0b00011000)
)

// Bit identifiers for block 0 instructions identified by the first nibble
const (
	b0Op4BitsLdR16Imm16 = 0b00_00_0001
	b0Op4BitsLdR16MemA  = 0b00_00_0010
	b0Op4BitsIncR16     = 0b00_00_0011
	b0Op4BitsAddHlR16   = 0b00_00_1001
	b0Op4BitsLdAR16Mem  = 0b00_00_1010
	b0Op4BitsDecR16     = 0b00_00_1011
)

const (
	b0Op3BitsIncR8    = 0b00_000_100
	b0Op3BitsDecR8    = 0b00_000_101
	b0Op3BitsLdR8Imm8 = 0b00_000_110
)

const (
	b0Zeros3BitsJrImm8     = 0b000_00_000
	b0Zeros3BitsJrCondImm8 = 0b001_00_000
)

// Bit identifiers for instructions with three ones in the least significant bits
const (
	b0Ones3BitsRlca = 0b00_000_000
	b0Ones3BitsRrca = 0b00_001_000
	b0Ones3BitsRla  = 0b00_010_000
	b0Ones3BitsRra  = 0b00_011_000
	b0Ones3BitsDaa  = 0b00_100_000
	b0Ones3BitsCpl  = 0b00_101_000
	b0Ones3BitsScf  = 0b00_110_000
	b0Ones3BitsCcf  = 0b00_111_000
)

const (
	b0UniqueNOp  = 0b00000000
	b0UniqueStop = 0b00010000
)
