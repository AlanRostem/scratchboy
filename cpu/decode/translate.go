package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

func TranslateOpcode(opcode nums.Byte) (OpcodeFormat, error) {
	if checkIllegal(opcode) {
		return OpcodeFormat{IsIllegal: true}, nil
	}
	if opcode == 0xCB {
		return OpcodeFormat{IsCBPrefix: true}, nil
	}
	switch {
	case isInOctalRowInterval(opcode, 0o00, 0o07):
		return decodeOctalRow00To07(opcode)
	case isInOctalRowInterval(opcode, 0o10, 0o17):
	case isInOctalRowInterval(opcode, 0o20, 0o27):
	case isInOctalRowInterval(opcode, 0o30, 0o37):
	}
	return OpcodeFormat{}, fmt.Errorf("opcode not implemented")
}
