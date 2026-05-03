package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

func DecodeStandard(opcode nums.Byte) (OpcodeFormat, error) {
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
		return decodeOctalRow10To17(opcode)
	case isInOctalRowInterval(opcode, 0o20, 0o27):
		return decodeOctalRow20To27(opcode)
	case isInOctalRowInterval(opcode, 0o30, 0o37):
		return decodeOctalRow30To37(opcode)
	}
	return OpcodeFormat{}, fmt.Errorf("opcode not implemented")
}

func DecodePrefixed(opcode nums.Byte) (OpcodeFormat, error) {
	switch {
	case isInOctalRowInterval(opcode, 0o00, 0o07):
		return decodeOctalRowPrefixed00to07(opcode)
	case isInOctalRowInterval(opcode, 0o10, 0o37):
		return decodeOctalRowPrefixed10to37(opcode)
	}
	return OpcodeFormat{}, fmt.Errorf("opcode not implemented")
}
