#include <iostream>

#include "allocator/heap.h"
#include "cpu/cpu.h"

int main()
{
    const scr::VTableAllocator* heap = scr::system_heap::Get();
    void* cpuVoid = heap->Alloc(sizeof(scr::CPU));
    scr::CPU* cpu = reinterpret_cast<scr::CPU*>(cpuVoid);
    scr::CPU_Init(cpu);
    scr::CPU_Step(cpu);
    heap->Free(cpu);
    scr::debug::CheckMemoryLeaks();
}