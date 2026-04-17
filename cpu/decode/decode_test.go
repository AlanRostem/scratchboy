package decode_test

import (
	"os"
	"testing"

	"github.com/AlanRostem/scratchboy/cpu/decode"
	"github.com/AlanRostem/scratchboy/nums"
)

func TestDecoding(t *testing.T) {
	data, err := os.ReadFile("testdata/cpu_instrs.gb")
	if err != nil {
		t.Fatal(err)
	}
	for _, mc := range data {
		oc, err := decode.TranslateOpcode(nums.Byte(mc))
		if err != nil {
			// skip the blocks we didn't implement
			continue
		}
		info, err := oc.Decode()
		if err != nil {
			t.Log(err)
			continue
		}
		t.Logf("0x%02X == %s", oc, info.InstructionId)
	}
}
