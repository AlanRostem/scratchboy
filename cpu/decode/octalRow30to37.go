package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

// decodeOctalRow30To37 has a few illegal opcodes, but those are not going to be considered
// since we check them before this function is called.
func decodeOctalRow30To37(opcode nums.Byte) (f OpcodeFormat, err error) {
	switch opcode.OctalCol() {
	case 0o00:
		switch {
		case isInOctalRowInterval(opcode, 0o30, 0o33):
			f.InstructionEnum = RetCond
			const condMask = 0b000_11_000
			cond := (condMask & opcode) >> 3
			f.EncodedOperandCount = 1
			f.EncodedOperands = [2]nums.Byte{
				cond,
			}
			return
		default:
			f.ImmediateCount = 1
			switch opcode.OctalRow() {
			case 0o34:
				f.InstructionEnum = LdhImm8A
				return
			case 0o35:
				f.InstructionEnum = AddSpImm8
				return
			case 0o36:
				f.InstructionEnum = LdhAImm8
				return
			case 0o37:
				f.InstructionEnum = LdHlSpImm8
				return
			}

		}
	case 0o01:
		switch opcode.OctalRow() % 2 {
		case 0:
			f.InstructionEnum = PopR16stk
			const r16StkMask = 0b00_11_0000
			r16Stk := (opcode & r16StkMask) >> 4
			f.EncodedOperandCount = 1
			f.EncodedOperands = [2]nums.Byte{
				r16Stk,
			}
			return
		default:
			switch opcode.OctalRow() {
			case 0o31:
				f.InstructionEnum = Ret
				return
			case 0o33:
				f.InstructionEnum = RetI
				return
			case 0o35:
				f.InstructionEnum = JpHl
				return
			case 0o37:
				f.InstructionEnum = LdSpHl
				return
			}
		}

	case 0o02:
		switch {
		case isInOctalRowInterval(opcode, 0o30, 0o33):
			f.InstructionEnum = JpCondImm16
			f.ImmediateCount = 2
			const condMask = 0b000_11_000
			cond := (condMask & opcode) >> 3
			f.EncodedOperandCount = 1
			f.EncodedOperands = [2]nums.Byte{
				cond,
			}
			return
		case isInOctalRowInterval(opcode, 0o34, 0o37):
			switch opcode.OctalRow() {
			case 0o34:
				f.InstructionEnum = LdhCA
				return
			case 0o35:
				f.InstructionEnum = LdImm16A
				f.ImmediateCount = 2
				return
			case 0o36:
				f.InstructionEnum = LdhAC
				return
			case 0o37:
				f.InstructionEnum = LdAImm16
				f.ImmediateCount = 2
				return
			}
		}
	case 0o03:
		switch opcode.OctalRow() {
		case 0o30:
			f.InstructionEnum = JpImm16
			f.ImmediateCount = 2
			return
		case 0o31:
			f.IsCBPrefix = true
			return
		case 0o36:
			f.InstructionEnum = Di
			return
		case 0o37:
			f.InstructionEnum = Ei
			return
		}
	case 0o04:
		if isInOctalRowInterval(opcode, 0o30, 0o33) {
			const condMask = 0b000_11_000
			cond := (opcode & condMask) >> 3
			f.ImmediateCount = 2
			f.InstructionEnum = CallCondImm16
			f.EncodedOperandCount = 1
			f.EncodedOperands = [2]nums.Byte{
				cond,
			}
			return // else: it's illegal and we return the error
		}
	case 0o05:
		switch {
		case opcode.OctalRow()%2 == 0:
			f.InstructionEnum = PushR16stk
			const r16StkMask = 0b00_11_0000
			r16Stk := (r16StkMask & opcode) >> 4
			f.EncodedOperandCount = 1
			f.EncodedOperands = [2]nums.Byte{
				r16Stk,
			}
			return
		case opcode.OctalRow() == 0o31:
			f.InstructionEnum = CallImm16
			f.ImmediateCount = 2
			return
		}
	case 0o06:
		f.ImmediateCount = 1
		switch opcode.OctalRow() {
		case 0o30:
			f.InstructionEnum = AddAImm8
			return
		case 0o31:
			f.InstructionEnum = AdcAImm8
			return
		case 0o32:
			f.InstructionEnum = SubAImm8
			return
		case 0o33:
			f.InstructionEnum = SbcAImm8
			return
		case 0o34:
			f.InstructionEnum = AndAImm8
			return
		case 0o35:
			f.InstructionEnum = XorAImm8
			return
		case 0o36:
			f.InstructionEnum = OrAImm8
			return
		case 0o37:
			f.InstructionEnum = CpAImm8
			return
		}
	case 0o07:
		const tgt3Mask = 0b00_111_000
		tgt3 := (tgt3Mask & opcode) >> 3
		f.EncodedOperandCount = 1
		f.EncodedOperands = [2]nums.Byte{
			tgt3,
		}
		f.InstructionEnum = RstTgt3
		return
	}
	return OpcodeFormat{}, fmt.Errorf("could not decode opcode in octal rows (30-37): 0x%02X", opcode)
}
