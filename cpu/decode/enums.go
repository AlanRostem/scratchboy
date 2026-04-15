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
	InvalidInstruction = InstructionIdCount
)

func (id InstructionId) String() string {
	switch id {
	case NOp:
		return "NOp"
	case LdR16Imm16:
		return "LdR16Imm16"
	case LdR16memA:
		return "LdR16memA"
	case LdAR16mem:
		return "LdAR16mem"
	case LdImm16Sp:
		return "LdImm16Sp"
	case IncR16:
		return "IncR16"
	case DecR16:
		return "DecR16"
	case AddHlR16:
		return "AddHlR16"
	case IncR8:
		return "IncR8"
	case DecR8:
		return "DecR8"
	case LdR8Imm8:
		return "LdR8Imm8"
	case Rlca:
		return "Rlca"
	case Rrca:
		return "Rrca"
	case Rla:
		return "Rla"
	case Rra:
		return "Rra"
	case Daa:
		return "Daa"
	case Cpl:
		return "Cpl"
	case Scf:
		return "Scf"
	case Ccf:
		return "Ccf"
	case JrImm8:
		return "JrImm8"
	case JrCondImm8:
		return "JrCondImm8"
	case Stop:
		return "Stop"
	default:
		return "UnknownInstruction"
	}
}
