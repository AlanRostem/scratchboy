package decode

import (
	"github.com/AlanRostem/scratchboy/nums"
)

func decodeOctalRow10To17(opcode nums.Byte) (f OpcodeFormat, err error) {
	const r16MaskDst = 0b00_111_000
	const r16MaskSrc = 0b00_000_111
	r8Dst := (opcode & r16MaskDst) >> 3
	r8Src := (opcode & r16MaskSrc)
	if r8Dst == R16HL && r8Src == R16HL {
		f.InstructionEnum = Halt
		return
	}
	f.InstructionEnum = LdR8R8
	f.EncodedOperandCount = 2
	f.EncodedOperands = [2]nums.Byte{
		r8Dst,
		r8Src,
	}
	return
}
