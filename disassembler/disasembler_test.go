package disassembler_test

import (
	"testing"

	"github.com/AlanRostem/scratchboy/disassembler"
)

func TestParseOpcodesJson(t *testing.T) {
	obj, err := disassembler.ParseOpcodesJson("opcodes.json")
	if err != nil {
		t.Fatal(err)
	}
	_, ok := obj.Unprefixed[0x00]
	if !ok {
		t.Fail()
	}
	for op, value := range obj.Unprefixed {
		t.Logf("0x%02X/%s: %dB", op, value.Mnemonic, value.Bytes)
	}
	_, ok = obj.CBprefixed[0x00]
	if !ok {
		t.Fail()
	}
}
