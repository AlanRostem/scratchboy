package register

import "github.com/AlanRostem/scratchboy/types"

type RegisterFile struct {
	registers []types.Byte
}

func NewRegisterFile() *RegisterFile {
	return &RegisterFile{
		registers: make([]types.Byte, RegisterCount),
	}
}

func (rf *RegisterFile) Set(reg Register, value types.Byte) {
	rf.registers[reg] = value
}

func (rf *RegisterFile) Get(reg Register) types.Byte {
	return rf.registers[reg]
}

func (rf *RegisterFile) VSet(vreg VirtualRegister, value types.DByte) {
	// TODO: implement
}

func (rf *RegisterFile) VGet(vreg VirtualRegister) types.DByte {
	return 255 // TODO: implement
}

func (rf *RegisterFile) getCombinedRegValue(reg0 Register, reg1 Register) types.DByte {
	// val0 := rf.registers[reg0]
	// val1 := rf.registers[reg1]
	// combinedValue :=
	return types.DByte(RegisterCount)
}
