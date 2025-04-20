#include <string.h>

#include "memory_bus.h"

void scr::MemoryBus_Init(MemoryBus* self)
{
    memset(self->data, 0, MemoryBusSize);
}

void scr::MemoryBus_Write(MemoryBus* self, Address addr, Word value)
{
    self->data[addr] = value;
}

scr::Word scr::MemoryBus_Read(MemoryBus* self, Address addr)
{
    return self->data[addr];
}
