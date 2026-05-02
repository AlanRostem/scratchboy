package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

type Opcode interface {
	// DecodePartial decodes the instruction into the correct
	// format, excluding the immediate bytes. The decoder in
	// the residing package is responsible for the immediate
	// byte decoding.
	DecodePartial() (InstructionFormat, error)
}

func TranslateStandardOpcode(byteRepresentation nums.Byte) (Opcode, error) {
	if checkIllegal(byteRepresentation) {
		return IllegalOpcode(byteRepresentation), nil
	}
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

func checkIllegal(byteRepresentation nums.Byte) bool {
	switch byteRepresentation {
	case illegalD3:
		fallthrough
	case illegalDB:
		fallthrough
	case illegalDD:
		fallthrough
	case illegalE3:
		fallthrough
	case illegalE4:
		fallthrough
	case illegalEB:
		fallthrough
	case illegalEC:
		fallthrough
	case illegalED:
		fallthrough
	case illegalF4:
		fallthrough
	case illegalFC:
		fallthrough
	case illegalFD:
		return true
	}
	return false
}

func TranslateCBPrefixedOpcode(byteRepresentation nums.Byte) (Opcode, error) {
	return CBPrefixed(byteRepresentation), nil
}
