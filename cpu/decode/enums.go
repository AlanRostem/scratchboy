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
	LdR8R8
	Halt

	// InstructionIdCount is not an instruction, only used to count the max number of instructions
	InstructionIdCount
	InvalidInstruction = InstructionIdCount
)

func (id InstructionId) String() string {
	switch id {
	case NOp:
		return "NOP"
	case LdR16Imm16:
		return "LD r16, imm16"
	case LdR16memA:
		return "LD r16mem, A"
	case LdAR16mem:
		return "LD A, r16mem"
	case LdImm16Sp:
		return "LD imm16, SP"
	case IncR16:
		return "INC r16"
	case DecR16:
		return "DEC r16"
	case AddHlR16:
		return "ADD [HL], r16"
	case IncR8:
		return "INC r8"
	case DecR8:
		return "DEC r8"
	case LdR8Imm8:
		return "LD r8, imm8"
	case Rlca:
		return "RLCA"
	case Rrca:
		return "RRCA"
	case Rla:
		return "RLA"
	case Rra:
		return "RRA"
	case Daa:
		return "DAA"
	case Cpl:
		return "CPL"
	case Scf:
		return "SCF"
	case Ccf:
		return "CCF"
	case JrImm8:
		return "JR imm8"
	case JrCondImm8:
		return "JR cond, imm8"
	case Stop:
		return "STOP"
	case LdR8R8:
		return "LD r8, r8"
	case Halt:
		return "HALT"
	default:
		return "UNKNOWN"
	}
}
