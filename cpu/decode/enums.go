package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

type InstructionId nums.Byte

var _ fmt.Stringer = (*InstructionId)(nil)

const (
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
	AddAR8
	AdcAR8
	SubAR8
	SbcAR8
	AndAR8
	OrAR8
	XorAR8
	CpAR8
	AddAImm8
	AdcAImm8
	SubAImm8
	SbcAImm8
	AndAImm8
	XorAImm8
	OrAImm8
	CpAImm8
	Di
	Ei
	LdhCA
	LdhImm8A
	LdImm16A
	LdhAC
	LdhAImm8
	LdAImm16
	AddSpImm8
	LdHlSpImm8
	LdSpHl
	PopR16stk
	PushR16stk
	Ret
	RetI
	JpImm16
	JpHl
	CallImm16
	RetCond
	JpCondImm16
	CallCondImm16
	RstTgt3
	BitB3R8
	ResB3R8
	SetB3R8
	Rlc
	Rrc
	Rl
	Rr
	Sla
	Sra
	Swap
	Srl

	Illegal
	// InstructionIdCount is not an instruction, only used to count the max number of instructions
	InstructionIdCount
	InvalidInstruction = InstructionIdCount
)

// String returns the disassembled and paramtetrized assembly
// representation of the instruciton. This is useful for a
// disassembler implementation.
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
	case AddAR8:
		return "ADD A, r8"
	case AdcAR8:
		return "ADC A, r8"
	case SubAR8:
		return "SUB A, r8"
	case SbcAR8:
		return "SBC A, r8"
	case AndAR8:
		return "AND A, r8"
	case OrAR8:
		return "OR A, r8"
	case XorAR8:
		return "XOR A, r8"
	case CpAR8:
		return "CP A, r8"
	case AddAImm8:
		return "ADD A, imm8"
	case AdcAImm8:
		return "ADC A, imm8"
	case SubAImm8:
		return "SUB A, imm8"
	case SbcAImm8:
		return "SBC A, imm8"
	case AndAImm8:
		return "AND A, imm8"
	case XorAImm8:
		return "XOR A, imm8"
	case OrAImm8:
		return "OR A, imm8"
	case CpAImm8:
		return "CP A, imm8"
	case Di:
		return "DI"
	case Ei:
		return "EI"
	case LdhCA:
		return "LDH [C], A"
	case LdhImm8A:
		return "LDH [imm8], A"
	case LdImm16A:
		return "LD [imm16], A"
	case LdhAC:
		return "LDH A, [C]"
	case LdhAImm8:
		return "LDH A, imm8"
	case LdAImm16:
		return "LD A, imm16"
	case AddSpImm8:
		return "ADD SP, imm8"
	case LdHlSpImm8:
		return "LD HL, SP +imm8"
	case LdSpHl:
		return "LD SP, HL"
	case PopR16stk:
		return "POP r16stk"
	case PushR16stk:
		return "PUSH r16stk"
	case Ret:
		return "RET"
	case RetI:
		return "RETI"
	case JpImm16:
		return "JP imm16"
	case JpHl:
		return "JP [HL]"
	case CallImm16:
		return "CALL imm16"
	case RetCond:
		return "RET cond"
	case JpCondImm16:
		return "JP cond, imm16"
	case CallCondImm16:
		return "CALL cond, imm16"
	case RstTgt3:
		return "RST tgt3"
	case BitB3R8:
		return "BIT b3, r8"
	case ResB3R8:
		return "RES b3, r8"
	case SetB3R8:
		return "SET b3, r8"
	case Rlc:
		return "RLC r8"
	case Rrc:
		return "RRC r8"
	case Rl:
		return "RL r8"
	case Rr:
		return "RR r8"
	case Sla:
		return "SLA r8"
	case Sra:
		return "SRA r8"
	case Swap:
		return "SWAP r8"
	case Srl:
		return "SRL r8"
	case Illegal:
		return "ILLEGAL"
	default:
		return "UNKNOWN"
	}
}
