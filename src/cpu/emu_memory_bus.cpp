#include "emu_memory_bus.h"

scr::Word scr::EmuMemoryBus_Read(EmuMemoryBus bus, Address addr)
{
    return bus[addr];
}
