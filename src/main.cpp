#include <iostream>

#include "list/array_list.h"
#include "tree/binary_tree.h"
#include "list/linked_list.h"
#include "allocator/heap.h"

void printArrayList(scr::ArrayList<int>* list)
{
for (std::size_t i = 0; i < scr::array_list::Size(list); i++)
    {
        std::cout << *scr::array_list::Get(list, i) << ",";
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
    scr::array_list::Ctor(&list, heap);

    scr::array_list::Append(&list, 1);
    scr::array_list::Append(&list, 2);
    scr::array_list::Append(&list, 4);
    scr::array_list::Append(&list, 5);
    printArrayList(&list);
    scr::array_list::Remove(&list, 1);
    printArrayList(&list);
    scr::array_list::Insert(&list, 2, 6);
    printArrayList(&list);
    
    scr::array_list::Dtor(&list);
}