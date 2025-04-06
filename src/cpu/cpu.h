#pragma once

#include "register_set.h"
#include "memory_bus.h"

namespace scr
{
    struct CPU
    {
        MemoryBus bus;
        RegisterSet registers;
        VirtualWord programCounter;
    };
}