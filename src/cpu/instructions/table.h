#pragma once

#include "instruction.h"
#include "add.h"

namespace scr
{
    namespace instructions
    {
        const Instruction Table[256] = {
            [0x80] = CPU_ExecADD_A,
            [0x81] = CPU_ExecADD_A,
            [0x82] = CPU_ExecADD_A,
            [0x83] = CPU_ExecADD_A,
            [0x84] = CPU_ExecADD_A,
            [0x85] = CPU_ExecADD_A,
            [0x86] = CPU_ExecADD_A,
            [0x87] = CPU_ExecADD_A,
        };
    }
}