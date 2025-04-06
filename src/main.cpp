#include <iostream>

#include "list/array_list.h"
#include "tree/binary_tree.h"
#include "allocator/heap.h"

void printArrayList(scr::ArrayList<int> *list)
{
    auto iter = scr::ArrayList_Iterate(list);
    while (scr::ArrayListIterator_Next(&iter))
    {
        std::cout << *scr::ArrayListIterator_Get(&iter) << ",";
    }
    std::cout << std::endl;
}

int main()
{
    const scr::VTableAllocator *heap = scr::system_heap::Get();
    scr::ArrayList<int> *list = scr::NewArrayList<int>(heap);

    scr::ArrayList_Append(list, 1);
    scr::ArrayList_Append(list, 2);
    scr::ArrayList_Append(list, 4);
    scr::ArrayList_Append(list, 5);
    printArrayList(list);
    scr::ArrayList_Remove(list, 1);
    printArrayList(list);
    scr::ArrayList_Insert(list, 2, 6);
    printArrayList(list);

    scr::ArrayList_Insert(list, 2, 7);
    scr::ArrayList_Insert(list, 2, 7);
    scr::ArrayList_Insert(list, 2, 7);
    scr::ArrayList_Insert(list, 2, 7);
    scr::ArrayList_Insert(list, 2, 7);
    scr::ArrayList_Insert(list, 2, 7);
    scr::ArrayList_Insert(list, 2, 7);
    printArrayList(list);

    scr::ArrayList_DeInit(list);
    heap->Free(list);
    
    scr::debug::CheckMemoryLeaks();
}