package cpu

import "github.com/AlanRostem/scratchboy/types"

const SizeMemoryBus = 0xFFFF

type MemoryBus struct {
	data []types.Word
}

func NewMemoryBus() *MemoryBus {
	return &MemoryBus{
		data: make([]types.Word, SizeMemoryBus),
	}
}

func (mb *MemoryBus) Write(addr types.DWord, value types.Word) {
	mb.data[addr] = value
}

func (mb *MemoryBus) Read(addr types.DWord) types.Word {
	return mb.data[addr]
}
