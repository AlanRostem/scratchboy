package decode

import "github.com/AlanRostem/scratchboy/nums"

const (
	illegalD3 = 0xD3
	illegalDB = 0xDB
	illegalDD = 0xDD
	illegalE3 = 0xE3
	illegalE4 = 0xE4
	illegalEB = 0xEB
	illegalEC = 0xEC
	illegalED = 0xED
	illegalF4 = 0xF4
	illegalFC = 0xFC
	illegalFD = 0xFD
)

type IllegalOpcode nums.Byte

func (o IllegalOpcode) DecodePartial() (OpcodeFormat, error) {
	return OpcodeFormat{
		InstructionEnum: Illegal,
	}, nil
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
