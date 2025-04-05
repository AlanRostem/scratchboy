#include <stdlib.h>
#include <stdio.h>
#include <assert.h>

#include "heap.h"

#ifdef _DEBUG
namespace scr
{
    namespace debug
    {
        static struct {
            int heapAllocationCount;
        } s_Self;
        
        void* Alloc(size_t size)
        {   
            s_Self.heapAllocationCount++;
            return malloc(size);
        }
        
        void Free(void* ptr)
        {
            s_Self.heapAllocationCount--;
            free(ptr);
        }
        
        void CheckMemoryLeaks()
        {
            if (s_Self.heapAllocationCount == 0)
            {
                return;
            }
            printf("Memory leak detected. There are %d unfreed pointers.", s_Self.heapAllocationCount);
            assert(false);
        }
    }
}

#define malloc scr::debug::Alloc
#define free scr::debug::Free

#endif

const scr::VTableAllocator* scr::system_heap::Get()
{
    static VTableAllocator s_heapAllocator = VTableAllocator{
        .Alloc = malloc,
        .Free = free,
        .ReAlloc = realloc,
    };
    return &s_heapAllocator;
}
