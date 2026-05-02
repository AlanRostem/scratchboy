package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

const (
	octalRows10x = 0o10
)

type Opcode interface {
	// DecodePartial decodes the instruction into the correct
	// format, excluding the immediate bytes. The decoder in
	// the residing package is responsible for the immediate
	// byte decoding.
	DecodePartial() (InstructionFormat, error)
}

func isInOctalInterval(opcodeByte, octalRowStart, octalRowEnd nums.Byte) bool {
	// const octalMaskCol = 0b00000111
	const octalMaskRow = 0b11111000
	// octalCol := octalMaskCol & opcodeByte
	octalRow := octalMaskRow & opcodeByte
	return octalRow >= octalRowStart && octalMaskRow < octalRowEnd
}

func TranslateOpcode(byteRepresentation nums.Byte) (Opcode, error) {
	if checkIllegal(byteRepresentation) {
		return IllegalOpcode(byteRepresentation), nil
	}
	switch {
	case isInOctalInterval(byteRepresentation, 0, octalRows10x):

	}
	return nil, fmt.Errorf("remaining opcode blocks not implemented")
}
