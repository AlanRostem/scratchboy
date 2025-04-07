#pragma once

#include "types.h"

namespace scr
{
    using Address = VirtualWord;
    using EmuMemoryBus = Word[0xFFFF];

    Word EmuMemoryBus_Read(EmuMemoryBus bus, Address addr);
}
