package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

func decodeOctalRow00To07(opcode nums.Byte) (f OpcodeFormat, err error) {
	switch opcode.OctalCol() {
	case 0o00:
		switch opcode.OctalRow() {
		case 0o00:
			f.InstructionEnum = NOp
			return
		case 0o01:
			f.InstructionEnum = LdImm16A
			f.ImmediateCount = 2
			return
		case 0o02:
			f.InstructionEnum = Stop
			return
		}
		switch {
		case isInOctalRowInterval(opcode, 0o03, 0o07):
			f.InstructionEnum = JrCondImm8
			f.ImmediateCount = 1
			const condMask = 0b000_11_000
			cond := (condMask & opcode) >> 3
			f.EncodedOperandCount = 1
			f.EncodedOperands = [2]nums.Byte{
				cond,
			}
			return
		}
	case 0o01:
		const r16Mask = 0b00_11_0000
		r16 := (r16Mask & opcode) >> 4
		f.EncodedOperandCount = 1
		f.EncodedOperands = [2]nums.Byte{
			r16,
		}
		switch {
		case opcode.OctalRow()%2 == 0:
			f.InstructionEnum = LdR16Imm16
			f.ImmediateCount = 2
		case opcode.OctalRow()%2 != 0:
			f.InstructionEnum = AddHlR16
		}
		return
	case 0o02:
		const r16memMask = 0b00_11_0000
		r16mem := (opcode & r16memMask) >> 4
		f.EncodedOperandCount = 1
		f.EncodedOperands = [2]nums.Byte{
			r16mem,
		}
		switch {
		case opcode.OctalRow()%2 == 0:
			f.InstructionEnum = LdR16memA
		case opcode.OctalRow()%2 != 0:
			f.InstructionEnum = LdAR16mem
		}
		return
	case 0o03:
		const r16Mask = 0b00_11_0000
		r16 := (opcode & r16Mask) >> 4
		f.EncodedOperandCount = 1
		f.EncodedOperands = [2]nums.Byte{
			r16,
		}
		switch {
		case opcode.OctalRow()%2 == 0:
			f.InstructionEnum = IncR16
		case opcode.OctalRow()%2 != 0:
			f.InstructionEnum = DecR16
		}
		return
	}
	switch {
	case isInOctalColInterval(opcode, 0o04, 0o05):
		const r8Mask = 0b00_111_000
		r8 := (r8Mask & opcode) >> 3
		f.EncodedOperandCount = 1
		f.EncodedOperands = [2]nums.Byte{
			r8,
		}
		switch opcode.OctalCol() {
		case 0o04:
			f.InstructionEnum = IncR8
		case 0o05:
			f.InstructionEnum = DecR8
		}
		return
	}
	return OpcodeFormat{}, fmt.Errorf("could not decode opcode in octal rows (00-07): 0x%02X", opcode)
}
