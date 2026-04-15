package decode

import (
	num "github.com/AlanRostem/scratchboy/nums"
)

type Opcode interface {
	Decode() (Info, error)
}

func TranslateOpcode(byteRepresentation num.Byte) Opcode {
	block := Block(0b11000000 & byteRepresentation)
	switch block {
	case Block0:
		return Block0Opcode(byteRepresentation)
	default:
		panic("remaining opcode blocks not implemented")
	}
}
