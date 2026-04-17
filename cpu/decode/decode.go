package decode

import num "github.com/AlanRostem/scratchboy/nums"

type Block num.Byte

const (
	Block0 = Block(iota)
	Block1
	Block2
	Block3
	BlockInvalid
)

type Identity num.Byte

type Info struct {
	InstructionId InstructionId
	// EncodedOperands represents the portion of an opcode that contains an operand.
	// I.e., "LD r16, imm16" has one encoded operand "r16", but "imm16" is not encoded.
	EncodedOperands [2]num.Byte
	// EncOpsCount is the number of encoded operands. If it's zero, none will be used.
	EncOpsCount int
	// ImmediateCount represents the number of bytes following the opcode to include
	// in the decoding.
	ImmediateCount int
}
