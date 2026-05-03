package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

func decodeOctalRow20To27(opcode nums.Byte) (f OpcodeFormat, err error) {
	const r8Mask = 0b00000_111
	r8 := r8Mask & opcode
	f.EncodedOperandCount = 1
	f.EncodedOperands = [2]nums.Byte{
		r8,
	}
	switch opcode.OctalRow() {
	case 0o020:
		f.InstructionEnum = AddAR8
		return
	case 0o021:
		f.InstructionEnum = AdcAR8
		return
	case 0o022:
		f.InstructionEnum = SubAR8
		return
	case 0o023:
		f.InstructionEnum = SbcAR8
		return
	case 0o024:
		f.InstructionEnum = AndAR8
		return
	case 0o025:
		f.InstructionEnum = XorAR8
		return
	case 0o026:
		f.InstructionEnum = OrAR8
		return
	case 0o027:
		f.InstructionEnum = CpAR8
		return
	}
	return OpcodeFormat{}, fmt.Errorf("could not decode opcode in octal rows (20-27): 0x%02X", opcode)
}
