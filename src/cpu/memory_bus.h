#pragma once

#include <uchar.h>

#include "types/word.h"

namespace scr
{
    using Address = DoubleWord;

    constexpr size_t MemoryBusSize = 0xFFFF;
    struct MemoryBus
    {
        Word data[MemoryBusSize];
    };

    void MemoryBus_Init(MemoryBus* self);
    void MemoryBus_Write(MemoryBus* self, Address addr, Word value);
    Word MemoryBus_Read(MemoryBus* self, Address addr);
}
