#pragma once

#include "types.h"

namespace scr
{
    struct Instruction
    {
        Operand Op0;
        Operand Op1;
        Mnemonic Mnem;
    };

    Instruction Opcode_Decode(Opcode code);
}