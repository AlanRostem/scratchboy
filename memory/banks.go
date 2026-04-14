package memory

import "github.com/AlanRostem/scratchboy/nums"

const (
	WorkRAMSize     = 0x2000
	VideoRAMSize    = 0x2000
	ExternalRAMSize = 0x2000
	HighRamSize     = 127
)

type (
	WorkRAM     [WorkRAMSize]nums.Byte
	VideoRAM    [VideoRAMSize]nums.Byte
	ExternalRAM [ExternalRAMSize]nums.Byte
	HighRAM     [HighRamSize]nums.Byte
)
