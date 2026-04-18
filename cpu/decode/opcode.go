package decode

import (
	"fmt"

	num "github.com/AlanRostem/scratchboy/nums"
)

type Opcode interface {
	Decode() (Info, error)
}

func TranslateOpcode(byteRepresentation num.Byte) (Opcode, error) {
	block := Block(0b11000000&byteRepresentation) >> 6
	switch block {
	case Block0:
		return block0Opcode(byteRepresentation), nil
	case Block1:
		return block1Opcode(byteRepresentation), nil
	case Block2:
		return block2Opcode(byteRepresentation), nil
	case Block3:
		return block3Opcode(byteRepresentation), nil
	default:
		return nil, fmt.Errorf("remaining opcode blocks not implemented")
	}
}
