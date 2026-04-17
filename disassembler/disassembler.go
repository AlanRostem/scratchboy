package disassembler

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AlanRostem/scratchboy/cpu/decode"
	"github.com/AlanRostem/scratchboy/nums"
)

var r8Names = [8]string{
	0: "B",
	1: "C",
	2: "D",
	3: "E",
	4: "H",
	5: "L",
	6: "[HL]",
	7: "A",
}

var r16Names = [4]string{
	0: "BC",
	1: "DE",
	2: "HL",
	3: "SP",
}

var r16MemNames = [4]string{
	0: "BC",
	1: "DE",
	2: "HL+",
	3: "HL-",
}

var condNames = [4]string{
	0: "NZ",
	1: "Z",
	2: "NC",
	3: "C",
}

func findOperandTokens(instruction string) (string, []string) {
	possibleOps := []string{
		"r8",
		"r16",
		"r16mem",
		"cond",
		"imm8",
		"imm16",
		// TODO implement the rest
	}
	tokens := make([]string, 0)
	splitted := strings.Split(instruction, " ")
	for _, token := range splitted {
		token = strings.ReplaceAll(token, ",", "")
		if slices.Contains(possibleOps, token) {
			tokens = append(tokens, token)
		}
	}
	var mnemonic string
	if len(tokens) == 0 {
		mnemonic = instruction
	} else {
		mnemonic = splitted[0]
	}
	return mnemonic, tokens
}

func Disassemble(data []byte) (string, error) {
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
			source += fmt.Sprintf("ERROR IN DISSASEMBLE: %v\n", err)
			pc++
			continue
		}
		var line string = ""
		line += fmt.Sprintf("$%04X: ", pc)
		instruction := info.InstructionId.String()
		_, opTokens := findOperandTokens(instruction)
		if info.EncOpsCount > 0 {
			for i := range info.EncOpsCount {
				op := info.EncodedOperands[i]
				var operandValue string = ""
				switch opTokens[i] {
				case "r8":
					operandValue = r8Names[op]
				case "r16":
					operandValue = r16Names[op]
				case "cond":
					operandValue = condNames[op]
				case "r16mem":
					operandValue = r16MemNames[op]
				}
				if operandValue != "" {
					instruction = strings.ReplaceAll(instruction, opTokens[i], operandValue)
				}
			}
		}
		pc++
		if info.ImmediateCount > 0 {
			switch info.ImmediateCount {
			case 1:
				instruction = strings.ReplaceAll(instruction, "imm8", fmt.Sprintf("0x%02X", data[pc]))
				pc++
			case 2:
				left := nums.Byte(data[pc])
				right := nums.Byte(data[pc+1])
				whole := left.Concat(right)
				instruction = strings.ReplaceAll(instruction, "imm16", fmt.Sprintf("0x%04X", whole))
				pc += 2
			}
		}
		line += instruction
		source += line + "\n"
	}
	return source, nil
}
