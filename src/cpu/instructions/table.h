#pragma once

#include "instruction.h"
#include "add.h"

namespace scr
{
    namespace instructions
    {
        // TODO: make this array 256 large
        const Instruction Table[8] = {
            [0x0] = CPU_ExecADD_A, // TODO: the indices of these were 0x8n, change them back
            [0x1] = CPU_ExecADD_A,
            [0x2] = CPU_ExecADD_A,
            [0x3] = CPU_ExecADD_A,
            [0x4] = CPU_ExecADD_A,
            [0x5] = CPU_ExecADD_A,
            [0x6] = CPU_ExecADD_A,
            [0x7] = CPU_ExecADD_A,
        };
    }
}