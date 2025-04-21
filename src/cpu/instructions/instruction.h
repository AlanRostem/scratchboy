#pragma once

#include "types/word.h"

namespace scr
{
    struct CPU;
    using Instruction = void(*)(CPU*, Word);

    namespace instructions
    {
        void PrefixCB(CPU* cpu, Word opcode);
        void AddAR8(CPU* cpu, Word opcode);
    }
}