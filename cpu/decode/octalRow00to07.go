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
	case 0o02:
	case 0o03:
	}
	switch {
	case isInOctalColInterval(opcode, 0o04, 0o05):
	}
	return OpcodeFormat{}, fmt.Errorf("could not decode opcode in octal rows (00-07): 0x%02X", opcode)
}
