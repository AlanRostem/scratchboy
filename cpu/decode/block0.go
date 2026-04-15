package decode

// Bitmask for block 0 instructions identified by the first nibble
const (
	B0_OpNib_LdR16Imm16 = 0b00_00_0001
	B0_OpNib_LdR16MemA  = 0b00_00_0010
	B0_OpNib_LdAR16Mem  = 0b00_00_1010

	B0_OpNib_IncR16   = 0b00_00_0011
	B0_OpNib_DecR16   = 0b00_00_1011
	B0_OpNib_AddHlR16 = 0b00_00_1001
)

const (
	B0_Op3Bits_IncR8    = 0b00_000_100
	B0_Op3Bits_DecR8    = 0b00_000_101
	B0_Op3Bits_LdR8Imm8 = 0b00_000_110
	B0_Op3Bits_Zeros    = 0b00_000_000
)

const (
	B0_Zeros3Bits_JrImm8     = 0b00_0_00_000
	B0_Zeros3Bits_JrCondImm8 = 0b00_1_00_000
)

// Bitmask for instructions with three ones in the least significant bits
const (
	B0_Ones3Bits_Rlca = 0b00_000_000
	B0_Ones3Bits_Rrca = 0b00_001_000
	B0_Ones3Bits_Rla  = 0b00_010_000
	B0_Ones3Bits_Rra  = 0b00_011_000
	B0_Ones3Bits_Daa  = 0b00_100_000
	B0_Ones3Bits_Cpl  = 0b00_101_000
	B0_Ones3Bits_Scf  = 0b00_110_000
	B0_Ones3Bits_Ccf  = 0b00_111_000
)

const (
	B0_Special_NOp  = 0b00000000
	B0_Special_Stop = 0b00010000
)
