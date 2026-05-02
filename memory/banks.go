package memory

import "github.com/AlanRostem/scratchboy/nums"

const (
	WorkRAMOffsetStart     = 0xC000
	VideoRAMOffsetStart    = 0x8000
	ExternalRAMOffsetStart = 0xA000
	HighRAMOffsetStart     = 0xFF80
)

const (
	WorkRAMSize     = 0x2000
	VideoRAMSize    = 0x2000
	ExternalRAMSize = 0x2000
	HighRAMSize     = 127
)

const (
	WorkRAMOffsetEnd     = WorkRAMOffsetStart + WorkRAMSize - 1
	VideoRAMOffsetEnd    = VideoRAMOffsetStart + VideoRAMSize - 1
	ExternalRAMOffsetEnd = ExternalRAMOffsetStart + ExternalRAMSize - 1
	HighRamEndOffset     = HighRAMOffsetStart + HighRAMSize - 1
)

type (
	WorkRAM     [WorkRAMSize]nums.Byte
	VideoRAM    [VideoRAMSize]nums.Byte
	ExternalRAM [ExternalRAMSize]nums.Byte
	HighRAM     [HighRAMSize]nums.Byte
)

type Banks struct {
	WorkRAM
	VideoRAM
	ExternalRAM
	HighRAM
}
