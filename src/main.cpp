#include <iostream>

#include "list/array_list.h"
#include "tree/binary_tree.h"
#include "list/linked_list.h"
#include "allocator/heap.h"

void printArrayList(scr::ArrayList<int>* list)
{
for (std::size_t i = 0; i < scr::ArrayList_Size(list); i++)
    {
        std::cout << *scr::ArrayList_Get(list, i) << ",";
    }
    std::cout << std::endl;
}

/*void printTree(std::shared_ptr<scr::BinarySearchTree::Node> root)
{
    std::cout << root->Value << '(';
    if (root->ChildLeft != nullptr)
    {
        std::cout << root->ChildLeft->Value;
    }   
    else
    {
        std::cout << "x";
    }
    std::cout << ", ";
    if (root->ChildRight != nullptr)
    {
        std::cout << root->ChildRight->Value;
    }
    else
    {
        std::cout << "x";
    }

    std::cout << ")\n";
    if (root->ChildLeft != nullptr)
    {
        printTree(root->ChildLeft);
    }
    if (root->ChildRight != nullptr)
    {
        printTree(root->ChildRight);
    }
}*/

int main()
{
    const scr::VTableAllocator* heap = scr::GetSystemHeap();
    
    scr::ArrayList<int> list;
    scr::ArrayList_Ctor(&list, heap);

    scr::ArrayList_Append(&list, 1);
    scr::ArrayList_Append(&list, 2);
    scr::ArrayList_Append(&list, 4);
    scr::ArrayList_Append(&list, 5);
    printArrayList(&list);
    scr::ArrayList_Remove(&list, 1);
    printArrayList(&list);
    scr::ArrayList_Insert(&list, 2, 6);
    printArrayList(&list);
    
    scr::ArrayList_Dtor(&list);
}