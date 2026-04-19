package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

type block3Opcode nums.Byte

func (o block3Opcode) Decode() (Info, error) {
	// TODO handle CB
	if o&0b00_000_111 == b3ArithIdBits {
		id := InvalidInstruction
		switch o & 0b00_111_000 {
		case b3ArithImmAdd:
			id = AddAImm8
		case b3ArithImmAdc:
			id = AdcAImm8
		case b3ArithImmSub:
			id = SubAImm8
		case b3ArithImmSbc:
			id = SbcAImm8
		case b3ArithImmAnd:
			id = AndAImm8
		case b3ArithImmXor:
			id = XorAImm8
		case b3ArithImmOr:
			id = OrAImm8
		case b3ArithImmCp:
			id = CpAImm8
		}
		if id == InvalidInstruction {
			return Info{}, fmt.Errorf("block 3 arithmetic opcode not found: 0x%02X", o)
		}
		return Info{
			InstructionId:  id,
			ImmediateCount: 1,
			EncOpsCount:    0,
			CBPrefixed:     false,
		}, nil

	}
	return Info{}, fmt.Errorf("could not decode opcode: 0x%02X", o)
}
