#include "instruction.h"

namespace scr
{
    enum class OperandCond
    {
        NZ = 0,
        Z,
        NC,
        C,
    };
}

void scr::instructions::JP_COND_IMM16(CPU* cpu, Word opcode)
{
    
}
