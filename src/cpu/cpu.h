#pragma once

#include "register_bank.h"
#include "memory_bus.h"

namespace scr
{
    struct CPU
    {
        MemoryBus bus;
        RegisterBank registers;
        VirtualWord programCounter;
        VirtualWord stackPointer;
    };

    void CPU_Init(CPU* cpu);
    RegisterBank* CPU_GetRegisters(CPU* self);
}