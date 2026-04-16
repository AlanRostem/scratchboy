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
	Args          [2]num.Byte
	ArgCount      int
	CycleCount    int
}
