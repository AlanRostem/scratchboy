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

func NewBus(
	wramRef *WorkRAM,
	vramRef *VideoRAM,
	eramRef *ExternalRAM,
	hramRef *HighRAM,
) *AddrBus {
	return &AddrBus{
		wram: wramRef,
		vram: vramRef,
		eram: eramRef,
		hram: hramRef,
	}
}

func (ab *AddrBus) WriteRange(beginAddr nums.DByte, values []nums.Byte) {
	bank, offset := ab.getBank(beginAddr)
	start := beginAddr - offset
	end := start + nums.DByte(len(values))
	copy(bank[start:end], values[:])
}

func (ab *AddrBus) Write(addr nums.DByte, value nums.Byte) {
	bank, offset := ab.getBank(addr)
	bank[addr-offset] = value
}

func (ab *AddrBus) Read(addr nums.DByte) nums.Byte {
	bank, offset := ab.getBank(addr)
	return bank[addr-offset]
}

func isBetweenInclusive(x, a, b nums.DByte) bool {
	return x >= a && x <= b
}

func (ab *AddrBus) getBank(addr nums.DByte) (bank []nums.Byte, bankOffset nums.DByte) {
	// TODO find a cleaner way to return the offset
	// TODO implement reading from rom
	if isBetweenInclusive(addr, VideoRAMOffsetStart, VideoRAMOffsetEnd) {
		return ab.vram[:], VideoRAMOffsetStart
	}
	if isBetweenInclusive(addr, ExternalRAMOffsetStart, ExternalRAMOffsetEnd) {
		return ab.eram[:], ExternalRAMOffsetStart
	}
	if isBetweenInclusive(addr, WorkRAMOffsetStart, WorkRAMOffsetEnd) {
		// TODO research the CGB behavior for address D000 to DFFF
		return ab.wram[:], WorkRAMOffsetStart
	}
	// TODO check how the prohibited echo ram works
	if isBetweenInclusive(addr, ExternalRAMOffsetStart, ExternalRAMOffsetEnd) {
		return ab.eram[:], ExternalRAMOffsetStart
	}
	panic(fmt.Errorf("memory bank implementation does not exist for address: 0x%04X", addr))
}
