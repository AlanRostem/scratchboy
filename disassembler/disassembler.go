package disassembler

import (
	"fmt"

	"github.com/AlanRostem/scratchboy/cpu/decode"
	"github.com/AlanRostem/scratchboy/nums"
)

type Disassembler struct {
}

func New() *Disassembler {
	return &Disassembler{}
}

func (d *Disassembler) Disassemble(data []byte) (string, error) {
	pc := 0
	source := ""
	for pc < len(data) {
		mc := data[pc]
		oc, err := decode.TranslateOpcode(nums.Byte(mc))
		if err != nil {
			// skip the blocks we didn't implement
			pc++
			continue
		}
		info, err := oc.Decode()
		if err != nil {
			fmt.Printf("ERROR IN DISSASEMBLE: %v", err)
			pc++
			continue
		}
		line := fmt.Sprintf("0x%02X: 0x%02X == %s", pc, oc, info.InstructionId)
		if info.EncOpsCount > 0 {
			line += " "
			for i := range info.EncOpsCount {
				op := info.EncodedOperands[i]
				line += fmt.Sprintf("op_0x%02X", op)
			}
		}
		if info.ImmediateCount > 0 {
			switch info.ImmediateCount {
			case 1:
				pc++
				line += fmt.Sprintf(", 0x%02X", data[pc])
			case 2:
				left := nums.Byte(data[pc])
				right := nums.Byte(data[pc+1])
				whole := left.Concat(right)
				line += fmt.Sprintf(", 0x%02X", whole)
				pc += 2
			}
		}
		source += line + "\n"
		pc++
	}
	return source, nil
}
