package decode

import (
	num "github.com/AlanRostem/scratchboy/nums"
)

type Opcode num.Byte

type OperandInfo struct {
	Operands [2]num.Byte
	Count    int
}

func (o Opcode) Block() Block {
	return Block(0b11000000 & o)
}

func (o Opcode) block0Decode() {

}

func (o Opcode) Decode() Info {
	info := Info{
		Block: o.Block(),
	}
	switch info.Block {
	case Block0:
		o.block0Decode()
	default:
		panic("decoding all blocks is not implemented")
	}
	return info
}
