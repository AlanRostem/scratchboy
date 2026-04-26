package decode

import (
	"fmt"
	"slices"
	"strings"

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

var r16stkNames = []string{
	0: "BC",
	1: "DE",
	2: "HL",
	3: "AF",
}

var condNames = [4]string{
	0: "NZ",
	1: "Z",
	2: "NC",
	3: "C",
}
var tgt3Names = [8]string{
	0b000: "00h",
	0b001: "08h",
	0b010: "10h",
	0b011: "18h",
	0b100: "20h",
	0b101: "28h",
	0b110: "30h",
	0b111: "38h",
}

var possibleOps = [9]string{
	"r8",
	"r16",
	"r16mem",
	"r16stk",
	"cond",
	"imm8",
	"imm16",
	"tgt3",
	"b3",
}

func findOperandTokens(instruction string) (string, []string) {

	tokens := make([]string, 0)
	splitted := strings.Split(instruction, " ")
	for _, token := range splitted {
		token = strings.ReplaceAll(token, ",", "")
		token = strings.ReplaceAll(token, "+", "")
		if slices.Contains(possibleOps[:], token) {
			tokens = append(tokens, token)
		}
	}
	mnemonic := splitted[0]
	return mnemonic, tokens
}

type Identity nums.Byte

// InstructionFormat is the emulator's format of the DMG CPU instructions
type InstructionFormat struct {
	InstructionId InstructionId
	// EncodedOperands represents the portion of an opcode that contains an operand.
	// I.e., "LD r16, imm16" has one encoded operand "r16", but "imm16" is not encoded.
	EncodedOperands [2]nums.Byte
	// EncOpsCount is the number of encoded operands. If it's zero, none should be used.
	EncOpsCount int
	// ImmmediateOperandBytes is each following byte of the instruction.
	ImmmediateOperandBytes [2]nums.Byte
	// ImmediateCount represents the number of bytes following the opcode to include
	// in the decoding.
	ImmediateCount int
	// Indicates if the instruction is the 0xCB opcode.
	IsCBPrefix bool
}

func (info *InstructionFormat) String() string {
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
			case "r16stk":
				operandValue = r16stkNames[op]
			case "tgt3":
				operandValue = tgt3Names[op]
			case "b3":
				operandValue = fmt.Sprintf("%d", op)
			}
			if operandValue != "" {
				// limiting the replace by 1 since we will replace by "appearence first"
				instruction = strings.Replace(instruction, opTokens[i], operandValue, 1)
			}
		}
	}
	if info.ImmediateCount > 0 {
		switch info.ImmediateCount {
		case 1:
			instruction = strings.ReplaceAll(
				instruction,
				"imm8",
				fmt.Sprintf(
					"0x%02X",
					info.ImmmediateOperandBytes[0]))
		case 2:
			right := nums.Byte(info.ImmmediateOperandBytes[0])
			left := nums.Byte(info.ImmmediateOperandBytes[1])
			whole := left.Concat(right)
			instruction = strings.ReplaceAll(
				instruction,
				"imm16",
				fmt.Sprintf(
					"0x%04X",
					whole))
		}
	}
	return instruction
}
