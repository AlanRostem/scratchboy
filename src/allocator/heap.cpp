#include <stdlib.h>

#include "heap.h"

const scr::VTableAllocator* scr::system_heap::Get()
{
    static VTableAllocator s_heapAllocator = VTableAllocator{
        .Alloc = malloc,
        .Free = free,
        .ReAlloc = realloc,
    };
    return &s_heapAllocator;
}
