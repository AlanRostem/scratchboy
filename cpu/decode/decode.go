package decode

import (
	"github.com/AlanRostem/scratchboy/nums"
	num "github.com/AlanRostem/scratchboy/nums"
)

type Block num.Byte

const (
	Block0 = Block(iota)
	Block1
	Block2
	Block3
	BlockInvalid
)

func DecodeProgram(program []byte) ([]InstructionFormat, error) {
	pc := 0
	source := make([]InstructionFormat, 0)
	cbMode := false
	for pc < len(program) {
		mc := program[pc]
		var oc Opcode
		var err error
		if cbMode {
			cbMode = false
			oc, err = TranslateCBPrefixedOpcode(nums.Byte(mc))
		} else {
			oc, err = TranslateStandardOpcode(nums.Byte(mc))
		}
		if err != nil {
			return nil, err
		}
		info, err := oc.DecodePartial()
		if err != nil {
			return nil, err
		}
		if info.IsCBPrefix {
			cbMode = true
			pc++
			continue
		}
		switch info.ImmediateCount {
		case 1:
			imm8 := program[pc+1]
			info.ImmmediateBytes[0] = nums.Byte(imm8)
		case 2:
			left := program[pc+1]
			right := program[pc+2]
			info.ImmmediateBytes[0] = num.Byte(left)
			info.ImmmediateBytes[1] = num.Byte(right)
		}
		source = append(source, info)
		pc += info.ImmediateCount + 1
	}
	return source, nil
}
