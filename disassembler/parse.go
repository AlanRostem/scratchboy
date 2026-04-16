package disassembler

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operand struct {
	Name      string
	Bytes     int
	Immediate bool
}

type flags struct {
	Z, N, H, C string
}

type OpcodeInfo struct {
	Mnemonic  string
	Bytes     int
	Cycles    []int
	Operands  []operand
	Immediate bool
	Flags     flags
}

type opcodesJson map[string]map[string]OpcodeInfo

type OpcodeTables struct {
	Unprefixed map[byte]OpcodeInfo
	CBprefixed map[byte]OpcodeInfo
}

func ParseJsonOpcodeTable(path string) (*OpcodeTables, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var obj opcodesJson
	err = json.Unmarshal([]byte(content), &obj)
	tables := &OpcodeTables{
		Unprefixed: make(map[byte]OpcodeInfo),
		CBprefixed: make(map[byte]OpcodeInfo),
	}
	fillTable := func(tableName string, out map[byte]OpcodeInfo) error {
		table, ok := obj[tableName]
		if !ok {
			return fmt.Errorf("table %s not found in json", tableName)
		}
		for key, value := range table {
			opcode, err := strconv.ParseInt(strings.ReplaceAll(key, "0x", ""), 16, 64)
			if err != nil {
				return err
			}
			out[byte(opcode)] = value
		}
		return nil
	}
	err = fillTable("unprefixed", tables.Unprefixed)
	if err != nil {
		return nil, err
	}
	err = fillTable("cbprefixed", tables.CBprefixed)
	if err != nil {
		return nil, err
	}
	return tables, err
}
