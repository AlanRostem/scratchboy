#pragma once

#include "cpu/cpu.h"

namespace scr
{
    void CPU_ExecADD_A(CPU* cpu, Register target);
    void CPU_ExecADD_A_HL(CPU* cpu);
}