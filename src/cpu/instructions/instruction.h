#pragma once

#include "types/word.h"

namespace scr
{
    struct CPU;
    using Instruction = void(*)(CPU*, Word);

    namespace instructions
    {
        void PREFIX_CB(CPU* cpu, Word opcode);
        void ADD_A_R8(CPU* cpu, Word opcode);
    }
}