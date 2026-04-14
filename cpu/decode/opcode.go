package decode

import num "github.com/AlanRostem/scratchboy/nums"

type Identity num.Byte

type Opcode num.Byte

type OperandInfo struct {
	Operands [2]num.Byte
	Count    int
}

func (o Opcode) Block() Block {
	return Block(0b11000000 & o)
}

func (o Opcode) Operands() OperandInfo {
	panic("not implemented")
}

func (o Opcode) Identity() Identity {
	panic("not implemented")
}
