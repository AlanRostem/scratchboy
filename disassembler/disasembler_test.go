package disassembler_test

import (
	"os"
	"testing"

	"github.com/AlanRostem/scratchboy/disassembler"
)

func TestDisassemble(t *testing.T) {
	data, err := os.ReadFile("testdata/cpu_instrs.gb")
	if err != nil {
		t.Fatal(err)
	}
	source, err := disassembler.Disassemble(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(source)
}
