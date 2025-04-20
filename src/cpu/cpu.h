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

    void CPU_Init(CPU* self);
    void CPU_Step(CPU* self);
    RegisterFile* CPU_GetRegisters(CPU* self);
}