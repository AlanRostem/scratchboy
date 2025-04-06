#include "instruction.h"

namespace scr
{
    inline constexpr Instruction newInstruction(Mnemonic mnemonic, Operand operand0, Operand operand1)
    {
        return Instruction{
            .Op0 = operand0,
            .Op1 = operand1,
            .Mnem = mnemonic,
        };
    }
}

scr::Instruction scr::Opcode_Decode(Opcode code)
{
    static constexpr Instruction instructionTable[] = {
        [0x00] = newInstruction(Mnemonic::NOP, Operand::None, Operand::None),
    };
    return instructionTable[code];
}