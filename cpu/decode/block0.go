package decode

import (
	"fmt"

	num "github.com/AlanRostem/scratchboy/nums"
)

type Block0Opcode num.Byte

func (o Block0Opcode) instructionId() InstructionId {
	switch o {
	case b0UniqueNOp:
		return NOp
	case b0UniqueStop:
		return Stop
	}
	// check instructions where the three least significant bits are ZERO
	if 0b0000_0111&o == 0 {
		masked := 0b111_00_000
		switch masked {
		case b0Zeros3BitsJrImm8:
			return JrImm8
		case b0Zeros3BitsJrCondImm8:
			return JrCondImm8
		}
	}
	// check instructions where the three least significant bits are ONE
	if 0b0000_0111&o == 0b0000_0111 {
		masked := 0b00_111_000 & o
		switch masked {
		case b0Ones3BitsRlca:
			return Rlca
		case b0Ones3BitsRrca:
			return Rrca
		case b0Ones3BitsRla:
			return Rla
		case b0Ones3BitsRra:
			return Rra
		case b0Ones3BitsDaa:
			return Daa
		case b0Ones3BitsCpl:
			return Cpl
		case b0Ones3BitsScf:
			return Scf
		case b0Ones3BitsCcf:
			return Ccf
		}
	}
	// check the three least significant bits for instruction id
	switch 0b0000_0111 & o {
	case b0Op3BitsDecR8:
		return DecR8
	case b0Op3BitsIncR8:
		return IncR8
	case b0Op3BitsLdR8Imm8:
		return LdR8Imm8
	}
	// check the least significant nibble for instruction id
	switch 0b0000_1111 & o {
	case b0Op4BitsLdR16Imm16:
		return LdR16Imm16
	case b0Op4BitsLdR16MemA:
		return LdR16memA
	case b0Op4BitsIncR16:
		return IncR16
	case b0Op4BitsAddHlR16:
		return AddHlR16
	case b0Op4BitsLdAR16Mem:
		return LdAR16mem
	case b0Op4BitsDecR16:
		return DecR16
	}
	return InvalidInstruction
}

func (o Block0Opcode) Decode() (Info, error) {
	id := o.instructionId()
	if id == InvalidInstruction {
		return Info{}, fmt.Errorf("opcode could not be decoded: %X", o)
	}
	return Info{
		InstructionId: id,
	}, nil
}

var _ Opcode = (*Block0Opcode)(nil)
