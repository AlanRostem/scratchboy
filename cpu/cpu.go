package cpu

import (
	"github.com/AlanRostem/scratchboy/cpu/register"
	"github.com/AlanRostem/scratchboy/memory"
	"github.com/AlanRostem/scratchboy/nums"
)

type CPU struct {
	bus       *AddrBus
	registers *register.GeneralPurposeFile
}

func New(banks *memory.Banks) *CPU {
	return &CPU{
		bus:       NewBus(banks),
		registers: register.NewGeneralPurposeFile(),
	}
}

func (cpu *CPU) Run(program []byte) {

}

func (cpu *CPU) MachineCycle(pc int, opcode nums.Byte) (nextPC int) {
	return
}
