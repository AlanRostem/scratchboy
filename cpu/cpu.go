package cpu

import "github.com/AlanRostem/scratchboy/nums"

type CPU struct {
}

func (cpu *CPU) Run(program []byte) {

}

func (cpu *CPU) MachineCycle(pc int, opcode nums.Byte) (nextPC int) {
	return
}
