package decode

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

const (
	cbBitIndexedMask = 0b11_000000
	cbStandardMask   = 0b11111_000
)

const (
	cbStandardRlc  = 0b00_000_000
	cbStandardRrc  = 0b00_001_000
	cbStandardRl   = 0b00_010_000
	cbStandardRr   = 0b00_011_000
	cbStandardSla  = 0b00_100_000
	cbStandardSra  = 0b00_101_000
	cbStandardSwap = 0b00_110_000
	cbStandardSrl  = 0b00_111_000
)

const (
	cbBitIndexedBitB3R8 = 0b01_000_000
	cbBitIndexedResB3R8 = 0b10_000_000
	cbBitIndexedSetB3R8 = 0b11_000_000
)

type CBPrefixed nums.Byte

func (o CBPrefixed) DecodePartial() (InstructionFormat, error) {
	id := InvalidInstruction
	switch o & cbBitIndexedMask {
	case cbBitIndexedBitB3R8:
		id = BitB3R8
	case cbBitIndexedResB3R8:
		id = ResB3R8
	case cbBitIndexedSetB3R8:
		id = SetB3R8
	}
	if id != InvalidInstruction {
		b3 := nums.Byte(o&0b00_111_000) >> 3
		r8 := nums.Byte(o & 0b00_000_111)
		return InstructionFormat{
			Partial: PartialFormat{
				InstructionId: id,
				EncOpsCount:   2,
				EncodedOperands: [2]nums.Byte{
					b3, r8,
				},
			},
		}, nil
	}
	switch o & cbStandardMask {
	case cbStandardRlc:
		id = Rlc
	case cbStandardRrc:
		id = Rrc
	case cbStandardRl:
		id = Rl
	case cbStandardRr:
		id = Rr
	case cbStandardSla:
		id = Sla
	case cbStandardSra:
		id = Sra
	case cbStandardSwap:
		id = Swap
	case cbStandardSrl:
		id = Srl
	}
	if id != InvalidInstruction {
		r8 := nums.Byte(o & 0b00_000_111)
		return InstructionFormat{
			Partial: PartialFormat{
				InstructionId: id,
				EncOpsCount:   1,
				EncodedOperands: [2]nums.Byte{
					r8,
				},
			},
		}, nil
	}
	return InstructionFormat{}, fmt.Errorf("could not decode: 0x%02X", o)
}
