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

type ArgInfo struct {
	Args  [4]num.Byte
	Count int
}

type Info struct {
	Block Block
	IsNOp bool
}

func getBlock(opcode num.Byte) Block {
	if 0b11000000&opcode == 0 {
		return Block0
	}
	return BlockInvalid
}
