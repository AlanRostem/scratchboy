package decode

import (
	"github.com/AlanRostem/scratchboy/nums"
)

type block1Opcode nums.Byte

const (
	b1OpcodeHalt = 0b01110110
)

const (
	b1OperandMaskDest   = 0b00_111_000
	b1OperandMaskSource = 0b00_000_111
)

func (o block1Opcode) DecodePartial() (InstructionFormat, error) {
	if o == b1OpcodeHalt {
		return InstructionFormat{
			Partial: PartialFormat{
				InstructionId: Halt,
			},
		}, nil
	}
	encOps := [2]nums.Byte{
		nums.Byte((o & b1OperandMaskDest) >> 3),
		nums.Byte(o & b1OperandMaskSource),
	}
	return InstructionFormat{
		Partial: PartialFormat{
			InstructionId:   LdR8R8,
			EncodedOperands: encOps,
			EncOpsCount:     2,
		},
		ImmediateCount: 0,
	}, nil
}

var _ Opcode = (*block1Opcode)(nil)
