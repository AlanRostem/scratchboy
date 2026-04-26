package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

type block2Opcode nums.Byte

const (
	b2OperandMask  = 0b00000_111
	b2MnemonicMask = 0b11111_000
)

const (
	b2MnemonicAdd = 0b10000_000
	b2MnemonicAdc = 0b10001_000
	b2MnemonicSub = 0b10010_000
	b2MnemonicSbc = 0b10011_000
	b2MnemonicAnd = 0b10100_000
	b2MnemonicXor = 0b10101_000
	b2MnemonicOr  = 0b10110_000
	b2MnemonicCp  = 0b10111_000
)

func (o block2Opcode) DecodePartial() (InstructionFormat, error) {
	var id InstructionId
	switch o & b2MnemonicMask {
	case b2MnemonicAdd:
		id = AddAR8
	case b2MnemonicAdc:
		id = AdcAR8
	case b2MnemonicSub:
		id = SubAR8
	case b2MnemonicSbc:
		id = SbcAR8
	case b2MnemonicAnd:
		id = AndAR8
	case b2MnemonicOr:
		id = OrAR8
	case b2MnemonicXor:
		id = XorAR8
	case b2MnemonicCp:
		id = CpAR8
	default:
		return InstructionFormat{}, fmt.Errorf("could not decode block1 opcode: 0x%02X", o)
	}
	encOps := [2]nums.Byte{
		nums.Byte(o & b2OperandMask),
	}
	return InstructionFormat{
		InstructionId:   id,
		EncodedOperands: encOps,
		EncOpsCount:     1,
		ImmediateCount:  0,
	}, nil
}

var _ Opcode = (*block2Opcode)(nil)
