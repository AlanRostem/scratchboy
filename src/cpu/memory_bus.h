#pragma once

#include "types.h"

namespace scr
{
    using Address = VirtualWord;
    using MemoryBus = Word[0xFFFF];

    Word MemoryBus_Read(MemoryBus bus, Address addr);
}
