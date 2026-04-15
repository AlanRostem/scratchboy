package decode

import num "github.com/AlanRostem/scratchboy/nums"

type B0BitMaskType num.Byte

const (
	B0BitMaskTypeOpNib = B0BitMaskType(iota)
	B0BitMaskTypeOp3Bits
	B0BitMaskTypeZeros3Bits
	B0BitMaskTypeOnes3Bits
	B0BitMaskTypeUnique
)

// Bitmask for block 0 instructions identified by the first nibble
const (
	B0Op4BitsLdR16Imm16 = 0b00_00_0001
	B0Op4BitsLdR16MemA  = 0b00_00_0010
	B0Op4BitsIncR16     = 0b00_00_0011
	B0Op4BitsAddHlR16   = 0b00_00_1001
	B0Op4BitsLdAR16Mem  = 0b00_00_1010
	B0Op4BitsDecR16     = 0b00_00_1011
)

const (
	B0Op3BitsIncR8    = 0b00_000_100
	B0Op3BitsDecR8    = 0b00_000_101
	B0Op3BitsLdR8Imm8 = 0b00_000_110
)

const (
	B0Zeros3BitsJrImm8     = 0b000_00_000
	B0Zeros3BitsJrCondImm8 = 0b001_00_000
)

// Bitmask for instructions with three ones in the least significant bits
const (
	B0Ones3BitsRlca = 0b00_000_000
	B0Ones3BitsRrca = 0b00_001_000
	B0Ones3BitsRla  = 0b00_010_000
	B0Ones3BitsRra  = 0b00_011_000
	B0Ones3BitsDaa  = 0b00_100_000
	B0Ones3BitsCpl  = 0b00_101_000
	B0Ones3BitsScf  = 0b00_110_000
	B0Ones3BitsCcf  = 0b00_111_000
)

const (
	B0UniqueNOp  = 0b00000000
	B0UniqueStop = 0b00010000
)
