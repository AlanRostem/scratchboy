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

func (o IllegalOpcode) DecodePartial() (InstructionFormat, error) {
	return InstructionFormat{
		InstructionId: Illegal,
	}, nil
}
