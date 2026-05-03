package cpu

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/memory"
	"github.com/AlanRostem/scratchboy/nums"
)

type AddrBus struct {
	wram *memory.WorkRAM
	vram *memory.VideoRAM
	eram *memory.ExternalRAM
	hram *memory.HighRAM
}

func NewBus(banks *memory.Banks) *AddrBus {
	return &AddrBus{
		wram: &banks.WorkRAM,
		vram: &banks.VideoRAM,
		eram: &banks.ExternalRAM,
		hram: &banks.HighRAM,
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
	if isBetweenInclusive(addr, memory.VideoRAMOffsetStart, memory.VideoRAMOffsetEnd) {
		return ab.vram[:], memory.VideoRAMOffsetStart
	}
	if isBetweenInclusive(addr, memory.ExternalRAMOffsetStart, memory.ExternalRAMOffsetEnd) {
		return ab.eram[:], memory.ExternalRAMOffsetStart
	}
	if isBetweenInclusive(addr, memory.WorkRAMOffsetStart, memory.WorkRAMOffsetEnd) {
		// TODO research the CGB behavior for address D000 to DFFF
		return ab.wram[:], memory.WorkRAMOffsetStart
	}
	// TODO check how the prohibited echo ram works
	if isBetweenInclusive(addr, memory.ExternalRAMOffsetStart, memory.ExternalRAMOffsetEnd) {
		return ab.eram[:], memory.ExternalRAMOffsetStart
	}
	panic(fmt.Errorf("memory bank implementation does not exist for address: 0x%04X", addr))
}
