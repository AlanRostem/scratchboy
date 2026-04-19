package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

type block3Opcode nums.Byte

func (o block3Opcode) Decode() (Info, error) {
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
		if id != InvalidInstruction {
			return Info{
				InstructionId:  id,
				ImmediateCount: 1,
			}, nil
		}
	}
	if o&0b00000_111 == b3InterruptIdBits {
		if o == b3CBPrefix {
			return Info{
				IsCBPrefix: true,
			}, nil
		}
		id := InvalidInstruction
		switch o & 0b00_111_000 {
		case b3InterruptDi:
			id = Di
		case b3InterruptEi:
			id = Ei
		}
		if id != InvalidInstruction {
			return Info{
				InstructionId: id,
			}, nil
		}
	}
	id := InvalidInstruction // TODO
	switch o & b3NoPatternBitMask {
	case b3NoPatternLdhCA:
		id = LdhCA
	case b3NoPatternLdhImm8A:
		id = LdhImm8A
	case b3NoPatternLdhImm16A:
		id = LdImm16A
	case b3NoPatternLdhAC:
		id = LdhAC
	case b3NoPatternLdhAImm8:
		id = LdhAImm8
	case b3NoPatternLdhAImm16:
		id = LdAImm16
	case b3NoPatternAddSpImm8:
		id = AddSpImm8
	case b3NoPatternLdHlSpImm8:
		id = LdHlSpImm8
	case b3NoPatternLdSpHl:
		id = LdSpHl
	}
	if id != InvalidInstruction {
		immCount := b3ImmediateByteCounts[id]
		return Info{
			InstructionId:  id,
			ImmediateCount: immCount,
		}, nil
	}
	return Info{}, fmt.Errorf("could not decode block 3 opcode: 0x%02X", o)
}
