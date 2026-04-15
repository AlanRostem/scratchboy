package decode

import num "github.com/AlanRostem/scratchboy/nums"

type InstructionId num.Byte

const (
	// Block 0 instructions
	NOp = InstructionId(iota)
	LdR16Imm16
	LdR16memA
	LdAR16mem
	LdImm16Sp
	IncR16
	DecR16
	AddHlR16
	IncR8
	DecR8
	LdR8Imm8
	Rlca
	Rrca
	Rla
	Rra
	Daa
	Cpl
	Scf
	Ccf
	JrImm8
	JrCondImm8
	Stop

	// InstructionIdCount is not an instruction, only used to count the max number of instructions
	InstructionIdCount
)
