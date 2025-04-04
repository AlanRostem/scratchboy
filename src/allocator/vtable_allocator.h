#pragma once

#include <stdint.h>

namespace scr
{
    struct VTableAllocator
    {
        void* (*Alloc)(size_t);
        void (*Free)(void*);
        void* (*ReAlloc)(void*, size_t);
    };
}