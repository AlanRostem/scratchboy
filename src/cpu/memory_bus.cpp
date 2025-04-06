#include "memory_bus.h"

scr::Word scr::MemoryBus_Read(MemoryBus bus, Address addr)
{
    return bus[addr];
}
