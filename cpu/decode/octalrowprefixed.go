package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

func decodeOctalRowPrefixed00to07(opcode nums.Byte) (f OpcodeFormat, err error) {
	const r8Mask = 0b00000_111
	r8 := r8Mask & opcode
	f.EncodedOperandCount = 1
	f.EncodedOperands = [2]nums.Byte{
		r8,
	}
	switch opcode.OctalRow() {
	case 0o00:
		f.InstructionEnum = Rlc
		return
	case 0o01:
		f.InstructionEnum = Rrc
		return
	case 0o02:
		f.InstructionEnum = Rl
		return
	case 0o03:
		f.InstructionEnum = Rr
		return
	case 0o04:
		f.InstructionEnum = Sla
		return
	case 0o05:
		f.InstructionEnum = Sra
		return
	case 0o06:
		f.InstructionEnum = Swap
		return
	case 0o07:
		f.InstructionEnum = Srl
		return
	}
	return OpcodeFormat{}, fmt.Errorf("could not decode prefixed opcode in octal rows (00-07): 0x%02X", opcode)
}

func decodeOctalRowPrefixed10to37(opcode nums.Byte) (f OpcodeFormat, err error) {
	const b3Mask = 0b00_111_000
	const r8Mask = 0b00_000_111
	b3 := (opcode & b3Mask) >> 3
	r8 := opcode & r8Mask
	f.EncodedOperandCount = 2
	f.EncodedOperands = [2]nums.Byte{
		b3,
		r8,
	}
	switch {
	case isInOctalRowInterval(opcode, 0o10, 0o17):
		f.InstructionEnum = BitB3R8
	case isInOctalRowInterval(opcode, 0o20, 0o27):
		f.InstructionEnum = ResB3R8
	case isInOctalRowInterval(opcode, 0o30, 0o37):
		f.InstructionEnum = SetB3R8
	}
	return
}
