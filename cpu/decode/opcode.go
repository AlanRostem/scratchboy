package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

const (
	octalRows10x = 0o10
)

type Opcode interface {
	Decode() (Info, error)
}

func TranslateStandardOpcode(byteRepresentation nums.Byte) (Opcode, error) {
	if checkIllegal(byteRepresentation) {
		return IllegalOpcode(byteRepresentation), nil
	}
	const octalMaskCol = 0b00000111
	const octalMaskRow = 0b11111000
	octalCol := octalMaskCol & byteRepresentation
	octalRow := octalMaskRow & byteRepresentation
	if octalMaskRow < octalRows10x {

	}
	return nil, fmt.Errorf("remaining opcode blocks not implemented")
}

func TranslateCBPrefixedOpcode(byteRepresentation nums.Byte) (Opcode, error) {
	return nil, fmt.Errorf("not implemented")
}
