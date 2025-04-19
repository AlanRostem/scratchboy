#pragma once

#include "types/word.h"

namespace scr
{
    using Address = DoubleWord;

    struct MemoryBus
    {
        Word data[0xFFFF];
    };

    Word MemoryBus_Read(MemoryBus* bus, Address addr);
}
