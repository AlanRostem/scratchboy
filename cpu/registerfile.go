package cpu

import "github.com/AlanRostem/scratchboy/types"

type RegisterFile struct {
	registers []types.Word
}

func NewRegisterFile() *RegisterFile {
	return &RegisterFile{
		registers: make([]types.Word, RegCount),
	}
}

func (rf *RegisterFile) Set(reg Register, value types.Word) {
	rf.registers[reg] = value
}

func (rf *RegisterFile) Get(reg Register) types.Word {
	return rf.registers[reg]
}

func (rf *RegisterFile) VSet(vreg VirtualRegister, value types.DWord) {
	// TODO: implement
}

func (rf *RegisterFile) VGet(vreg VirtualRegister) types.DWord {
	return 255 // TODO: implement
}

func (rf *RegisterFile) SetArithFlag(flag ArithmeticFlag, state bool) {

}

func (rf *RegisterFile) getCombinedRegValue(reg0 Register, reg1 Register) types.DWord {
	// val0 := rf.registers[reg0]
	// val1 := rf.registers[reg1]
	// combinedValue :=
	return types.DWord(RegCount)
}
