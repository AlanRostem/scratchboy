#pragma once

#include "types/word.h"

namespace scr
{
    using Address = VirtualWord;
    using MemoryBus = Word[0xFFFF];

    Word MemoryBus_Read(MemoryBus bus, Address addr);
}
