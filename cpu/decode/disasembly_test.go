package decode_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/AlanRostem/scratchboy/cpu/decode"
	"github.com/AlanRostem/scratchboy/nums"
)

func disassembleProgram(t *testing.T, program []byte) string {
	t.Helper()
	decoded, err := decode.DecodeProgram(program)
	if err != nil {
		t.Fatal(err)
	}
	source := ""
	for _, instruction := range decoded {
		source = fmt.Sprintf("%s%s\n", source, &instruction)
	}
	return source
}

func testFile(t *testing.T, path string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	source := disassembleProgram(t, data)
	t.Log("\n" + source)
}

type AllInstructionsTestData struct {
	WithoutImmediate []byte
	With1Immediate   []byte
	With2Immediate   []byte
	CbPrefixed       []byte
}

func CreateAllInstructionsTestData() *AllInstructionsTestData {
	withoutImmediate := make([]byte, 0)
	with1Immediate := make([]byte, 0)
	with2Immediate := make([]byte, 0)
	cbPrefixed := make([]byte, 0)
	for i := range 256 {
		cbPrefixed = append(cbPrefixed, 0xCB)
		cbPrefixed = append(cbPrefixed, byte(i))
	}
	for i := range 256 {
		instruction, err := decode.TranslateOpcode(nums.Byte(i))
		if err != nil {
			continue
		}
		if instruction.IsCBPrefix {
			continue
		}
		switch instruction.ImmediateCount {
		case 0:
			withoutImmediate = append(withoutImmediate, byte(i))
		case 1:
			with1Immediate = append(with1Immediate, byte(i))
			with1Immediate = append(with1Immediate, 0)
		case 2:
			with2Immediate = append(with2Immediate, byte(i))
			with2Immediate = append(with2Immediate, 0)
			with2Immediate = append(with2Immediate, 0)
		}
	}
	return &AllInstructionsTestData{
		WithoutImmediate: withoutImmediate,
		With1Immediate:   with1Immediate,
		With2Immediate:   with2Immediate,
		CbPrefixed:       cbPrefixed,
	}
}

func TestTranslateAllBytes(t *testing.T) {
	for raw := range 256 {
		opcode := nums.Byte(raw)
		f, err := decode.TranslateOpcode(opcode)
		if err != nil {
			t.Fail()
			continue
		}
		t.Logf("0x%02X: %s", opcode, f.String())
	}
}

func TestDisassembleEverything(t *testing.T) {
	testData := CreateAllInstructionsTestData()
	source := disassembleProgram(t, testData.WithoutImmediate)
	t.Log("No immediate bytes:\n" + source)
	source = disassembleProgram(t, testData.With1Immediate)
	t.Log("One immediate byte:\n" + source)
	source = disassembleProgram(t, testData.With2Immediate)
	t.Log("Two immediate bytes:\n" + source)
	source = disassembleProgram(t, testData.CbPrefixed)
	t.Log("CB prefixed:\n" + source)
}

func TestDisassembleBootRom(t *testing.T) {
	testFile(t, "testdata/bootix_dmg.bin")
}

func TestDisassembleInstrs(t *testing.T) {
	testFile(t, "testdata/cpu_instrs.gb")
}
