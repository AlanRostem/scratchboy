package decode

const (
	Nib4_LdR16Imm16 = 0b0000_0001
	Nib4_LdR16MemA  = 0b0000_0010
	Nib4_LdAR16Mem  = 0b0000_1010

	Nib4_IncR16   = 0b0000_0011
	Nib4_DecR16   = 0b0000_1011
	Nib4_AddHlR16 = 0b0000_1001
)

// const (
// 	Nib3
// )

const (
	Nib3AllOnes_Rlca = 0b00000_111
	Nib3AllOnes_Rrca = 0b00001_111
	Nib3AllOnes_Rla  = 0b00010_111
	Nib3AllOnes_Rra  = 0b00011_111
	Nib3AllOnes_Daa  = 0b00100_111
	Nib3AllOnes_Cpl  = 0b00101_111
	Nib3AllOnes_Scf  = 0b00110_111
	Nib3AllOnes_Ccf  = 0b00111_111
)

const (
	Nib4_LdImm16Sp = 0b0000_1000
	Full_Stop      = 0b00010000
)
