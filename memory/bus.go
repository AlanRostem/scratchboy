package memory

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/nums"
)

type AddrBus struct {
	wram *WorkRAM
	vram *VideoRAM
	eram *ExternalRAM
	hram *HighRAM
}

func (ab *AddrBus) Init(
	wram *WorkRAM,
	vram *VideoRAM,
	eram *ExternalRAM,
	hram *HighRAM,
) {
	ab.wram = wram
	ab.vram = vram
	ab.eram = eram
	ab.hram = hram
}

func (ab *AddrBus) Read(addr nums.DByte) nums.Byte {
	bank, offset := ab.getBank(addr)
	return bank[addr-offset]
}

func between(x, a, b nums.DByte) bool {
	return x >= a && x <= b
}

func (ab *AddrBus) getBank(addr nums.DByte) ([]nums.Byte, nums.DByte) {
	// TODO find a cleaner way to return the offset
	// TODO implement reading from rom
	if between(addr, VideoRAMOffsetStart, VideoRAMOffsetEnd) {
		return ab.vram[:], VideoRAMOffsetStart
	}
	if between(addr, ExternalRAMOffsetStart, ExternalRAMOffsetEnd) {
		return ab.eram[:], ExternalRAMOffsetStart
	}
	if between(addr, WorkRAMOffsetStart, WorkRAMOffsetEnd) {
		// TODO research the CGB behavior for address D000 to DFFF
		return ab.wram[:], WorkRAMOffsetStart
	}
	// TODO check how the prohibited echo ram works
	if between(addr, ExternalRAMOffsetStart, ExternalRAMOffsetEnd) {
		return ab.eram[:], ExternalRAMOffsetStart
	}
	panic(fmt.Errorf("memory bank implementation does not exist for address: 0x%04X", addr))
}
