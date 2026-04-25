package decode

import num "github.com/AlanRostem/scratchboy/nums"

type Identity num.Byte

type Info struct {
	InstructionId InstructionId
	// EncodedOperands represents the portion of an opcode that contains an operand.
	// I.e., "LD r16, imm16" has one encoded operand "r16", but "imm16" is not encoded.
	EncodedOperands [2]num.Byte
	// EncOpsCount is the number of encoded operands. If it's zero, none should be used.
	EncOpsCount int
	// ImmediateCount represents the number of bytes following the opcode to include
	// in the decoding.
	ImmediateCount int

	IsCBPrefix bool
}
