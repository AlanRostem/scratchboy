package disassembler

type readMode int

const (
	readModeOpcode = readMode(iota)
	readModeImm8
	readModeImm16
	readModeCB
)

type Disassembler struct {
	tables         *OpcodeTables
	readMode       readMode
	currentLine    string
	programCounter int
}

func New(opcodesPath string) (*Disassembler, error) {
	tables, err := ParseJsonOpcodeTable(opcodesPath)
	if err != nil {
		return nil, err
	}
	return &Disassembler{
		tables:         tables,
		readMode:       readModeOpcode,
		currentLine:    "",
		programCounter: 0,
	}, nil
}

func (d *Disassembler) Disassemble(program []byte) (string, error) {
	end := len(program)
	sourceCode := ""
	for d.programCounter < end {
		table := d.tables.Unprefixed
		machineCode := program[d.programCounter]
		info := table[machineCode]
		if len(info.Operands) > 0 {

		}
		d.programCounter++
	}
	return sourceCode, nil
}
