package decode_test

import (
	"testing"

	"github.com/AlanRostem/scratchboy/cpu/decode"
	"github.com/AlanRostem/scratchboy/nums"
)

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
		instruction, err := decode.DecodeStandard(nums.Byte(i))
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

func TestDecodeStandard(t *testing.T) {
	var errs []error
	for raw := range 256 {
		opcode := nums.Byte(raw)
		f, err := decode.DecodeStandard(opcode)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		t.Logf("0x%02X: %s", opcode, f.String())
	}
	if len(errs) != 0 {
		t.Fatalf("decode success rate: %d%%", int(float32(256-len(errs))/256.0*100.0))
	}
}

func TestDecodePrefixed(t *testing.T) {
	var errs []error
	for raw := range 256 {
		opcode := nums.Byte(raw)
		f, err := decode.DecodePrefixed(opcode)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		t.Logf("0x%02X: %s", opcode, f.String())
	}
	if len(errs) != 0 {
		t.Fatalf("decode success rate: %d%%", int(float32(256-len(errs))/256.0*100.0))
	}
}
