#pragma once

#include "emu_register_set.h"
#include "emu_memory_bus.h"

namespace scr
{
    struct EmuCPU
    {
        EmuMemoryBus bus;
        EmuRegisterSet registers;
        VirtualWord programCounter;
        VirtualWord stackPointer;
    };
}