#pragma once

#include "register_file.h"
#include "memory_bus.h"

namespace scr
{
    struct CPU
    {
        MemoryBus bus;
        RegisterFile registers;
        DoubleWord programCounter;
        DoubleWord stackPointer;
    };

    void CPU_Init(CPU* cpu);
    RegisterFile* CPU_GetRegisters(CPU* self);
}