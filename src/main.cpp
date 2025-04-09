#include <iostream>

#include "cpu/instructions/add.h"
#include "allocator/heap.h"

int main()
{
    const scr::VTableAllocator* heap = scr::system_heap::Get();
    void* cpuVoid = heap->Alloc(sizeof(scr::CPU));
    scr::CPU* cpu = reinterpret_cast<scr::CPU*>(cpuVoid);
    scr::CPU_Init(cpu);
    scr::CPU_ExecADD_A(cpu, scr::Register::E);
    heap->Free(cpu);
    scr::debug::CheckMemoryLeaks();
}