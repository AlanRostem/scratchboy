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
	id := InvalidInstruction
	switch o & b3NoPatternFlowBitMask {
	case b3FlowNoPatternRet:
		id = Ret
	case b3FlowNoPatternRetI:
		id = RetI
	case b3FlowNoPatternJpImm16:
		id = JpImm16
	case b3FlowNoPatternJpHl:
		id = JpHl
	case b3FlowNoPatternCallImm16:
		id = CallImm16
	}
	immCount := 0
	foundImmCount, ok := b3ImmediateByteCounts[id]
	if ok {
		immCount = foundImmCount
	}
	if id != InvalidInstruction {
		return Info{
			InstructionId:  id,
			ImmediateCount: immCount,
		}, nil
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
	switch o & 0b00_00_1111 {
	case b3StackPop:
		id = PopR16stk
	case b3StackPush:
		id = PushR16stk
	}
	if id != InvalidInstruction {
		return Info{
			InstructionId: id,
			EncodedOperands: [2]nums.Byte{
				0: nums.Byte(o&0b00_11_0000) >> 4,
			},
			EncOpsCount: 1,
		}, nil
	}
	switch o & b3EncodedOperandsBitMask {
	case b3EncodedOperandsRetCond:
		id = RetCond
	case b3EncodedOperandsJpCondImm16:
		id = JpCondImm16
	case b3EncodedOperandsCallCondImm16:
		id = CallCondImm16
	case b3EncodedOperandsRstTgt3:
		id = RstTgt3
	}
	immCount = 0
	foundImmCount, ok = b3ImmediateByteCounts[id]
	if ok {
		immCount = foundImmCount
	}
	if id != InvalidInstruction {
		return Info{
			InstructionId:  id,
			ImmediateCount: immCount,
			EncOpsCount:    1,
			EncodedOperands: [2]nums.Byte{
				// for tgt3, I take the 3rd bit here as well since for
				// the other instrucitions it's always zero so it wont
				// matter for this mask
				nums.Byte(0b001_11_000&o) >> 3,
			},
		}, nil
	}
	return Info{}, fmt.Errorf("could not decode block 3 opcode: 0x%02X", o)
}
