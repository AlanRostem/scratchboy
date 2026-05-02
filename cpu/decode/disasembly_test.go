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
		source = fmt.Sprintf("%s\n%s", source, &instruction)
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

func TestDisassembleEverything(t *testing.T) {
	withoutImmediate := make([]byte, 0)
	with1Immediate := make([]byte, 0)
	with2Immediate := make([]byte, 0)
	cbPrefixed := make([]byte, 0)
	for i := range 256 {
		cbPrefixed = append(cbPrefixed, 0xCB)
		cbPrefixed = append(cbPrefixed, byte(i))
	}
	for i := range 256 {
		o, err := decode.TranslateOpcode(nums.Byte(i))
		if err != nil {
			t.Fatal(err)
		}
		info, err := o.DecodePartial()
		if err != nil {
			t.Fatal(err)
		}
		if info.Partial.IsCBPrefix {
			continue
		}
		switch info.ImmediateCount {
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
	source := disassembleProgram(t, withoutImmediate[:])
	t.Log("No immediate bytes:\n" + source)
	source = disassembleProgram(t, with1Immediate[:])
	t.Log("One immediate byte:\n" + source)
	source = disassembleProgram(t, with2Immediate[:])
	t.Log("Two immediate bytes:\n" + source)
	source = disassembleProgram(t, cbPrefixed[:])
	t.Log("CB prefixed:\n" + source)
}

func TestDisassembleBootRom(t *testing.T) {
	testFile(t, "testdata/bootix_dmg.bin")
}

func TestDisassembleInstrs(t *testing.T) {
	testFile(t, "testdata/cpu_instrs.gb")
}
