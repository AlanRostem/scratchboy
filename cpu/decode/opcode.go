package decode

import (
	"fmt"

	num "github.com/AlanRostem/scratchboy/nums"
)

type Opcode interface {
	Decode() (Info, error)
}

func TranslateOpcode(byteRepresentation num.Byte) (Opcode, error) {
	block := Block(0b11000000 & byteRepresentation)
	switch block {
	case Block0:
		return Block0Opcode(byteRepresentation), nil
	default:
		return nil, fmt.Errorf("remaining opcode blocks not implemented")
	}
}
