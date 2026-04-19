package disassembler_test

import (
	"os"
	"testing"

	"github.com/AlanRostem/scratchboy/disassembler"
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
	t.Log(source)
}

func TestDisassembleBootRom(t *testing.T) {
	testFile(t, "testdata/bootix_dmg.bin")
}

func TestDisassembleInstrs(t *testing.T) {
	testFile(t, "testdata/cpu_instrs.gb")
}
