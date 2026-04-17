package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

type Block0Opcode nums.Byte

func (o Block0Opcode) idInfo() (id InstructionId, operandMask nums.Byte) {
	switch o {
	case b0UniqueNOp:
		return NOp, b0OperandMaskNone
	case b0UniqueStop:
		return Stop, b0OperandMaskNone
	}
	// check instructions where the three least significant bits are ZERO
	if 0b0000_0111&o == 0 {
		masked := 0b111_00_000 & o
		switch masked {
		case b0Zeros3BitsJrImm8:
			return JrImm8, b0OperandMaskZeros3Bits
		case b0Zeros3BitsJrCondImm8:
			return JrCondImm8, b0OperandMaskZeros3Bits
		}
	}
	// check instructions where the three least significant bits are ONE
	if 0b0000_0111&o == 0b0000_0111 {
		masked := 0b00_111_000 & o
		switch masked {
		case b0Ones3BitsRlca:
			return Rlca, b0OperandMaskNone
		case b0Ones3BitsRrca:
			return Rrca, b0OperandMaskNone
		case b0Ones3BitsRla:
			return Rla, b0OperandMaskNone
		case b0Ones3BitsRra:
			return Rra, b0OperandMaskNone
		case b0Ones3BitsDaa:
			return Daa, b0OperandMaskNone
		case b0Ones3BitsCpl:
			return Cpl, b0OperandMaskNone
		case b0Ones3BitsScf:
			return Scf, b0OperandMaskNone
		case b0Ones3BitsCcf:
			return Ccf, b0OperandMaskNone
		}
	}
	// check the three least significant bits for instruction id
	switch 0b0000_0111 & o {
	case b0Op3BitsDecR8:
		return DecR8, b0OperandMaskOp3Bits
	case b0Op3BitsIncR8:
		return IncR8, b0OperandMaskOp3Bits
	case b0Op3BitsLdR8Imm8:
		return LdR8Imm8, b0OperandMaskOp3Bits
	}
	// check the least significant nibble for instruction id
	switch 0b0000_1111 & o {
	case b0Op4BitsLdR16Imm16:
		return LdR16Imm16, b0OperandMaskOp4Bits
	case b0Op4BitsLdR16MemA:
		return LdR16memA, b0OperandMaskOp4Bits
	case b0Op4BitsIncR16:
		return IncR16, b0OperandMaskOp4Bits
	case b0Op4BitsAddHlR16:
		return AddHlR16, b0OperandMaskOp4Bits
	case b0Op4BitsLdAR16Mem:
		return LdAR16mem, b0OperandMaskOp4Bits
	case b0Op4BitsDecR16:
		return DecR16, b0OperandMaskOp4Bits
	}
	return InvalidInstruction, b0OperandMaskNone
}

func (o Block0Opcode) Decode() (Info, error) {
	id, mask := o.idInfo()
	if id == InvalidInstruction {
		return Info{}, fmt.Errorf("opcode could not be decoded: 0x%02X / 0b%08b", o, o)
	}
	immediateCount := 0
	count, ok := block0ImmediateByteCounts[id]
	if ok {
		immediateCount = count
	}
	encOpsCount := 0
	encOps := [2]nums.Byte{}
	if mask != b0OperandMaskNone {
		encOpsCount = 1
		encOps[0] = nums.Byte(o) & mask
		for mask&0b0000_0001 != 0b0000_0001 {
			mask >>= 1
			encOps[0] >>= 1
		}
	}
	return Info{
		InstructionId:   id,
		ImmediateCount:  immediateCount,
		EncodedOperands: encOps,
		EncOpsCount:     encOpsCount,
	}, nil
}

var _ Opcode = (*Block0Opcode)(nil)
