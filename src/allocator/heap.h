#pragma once

#include "vtable_allocator.h"

namespace scr
{
    namespace system_heap
    {
        const VTableAllocator* Get();
    }

#ifdef SCR_DEBUG
    namespace debug
    {
        void CheckMemoryLeaks();
    }
#endif
}