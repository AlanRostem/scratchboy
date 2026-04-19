package disassembler_test

import (
	"os"
	"testing"

	"github.com/AlanRostem/scratchboy/cpu/decode"
	"github.com/AlanRostem/scratchboy/disassembler"
	"github.com/AlanRostem/scratchboy/nums"
)

func testFile(t *testing.T, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	source, err := disassembler.Disassemble(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n" + source)
}

func TestDisassembleEverything(t *testing.T) {
	withoutImmediate := make([]byte, 0)
	with1Immediate := make([]byte, 0)
	with2Immediate := make([]byte, 0)
	everyOpcode := make([]byte, 0)
	for i := range 256 {
		everyOpcode = append(everyOpcode[:], byte(i))
		o, err := decode.TranslateStandardOpcode(nums.Byte(i))
		if err != nil {
			t.Fatal(err)
		}
		info, err := o.Decode()
		if info.IsCBPrefix {
			t.Log("WARNING: CB prefixed not implemented")
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
	source, err := disassembler.Disassemble(withoutImmediate[:])
	if err != nil {
		t.Fatal(err)
	}
	t.Log("No immediate bytes:\n" + source)

	source, err = disassembler.Disassemble(with1Immediate[:])
	if err != nil {
		t.Fatal(err)
	}
	t.Log("One immediate byte:\n" + source)
	source, err = disassembler.Disassemble(with2Immediate[:])
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Two immediate bytes:\n" + source)
}

func TestDisassembleBootRom(t *testing.T) {
	testFile(t, "testdata/bootix_dmg.bin")
}

func TestDisassembleInstrs(t *testing.T) {
	testFile(t, "testdata/cpu_instrs.gb")
}
