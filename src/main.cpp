#include <iostream>

#include "list/array_list.h"
#include "tree/binary_tree.h"

void printArrayList(const scr::ArrayList& list)
{
    for (std::size_t i = 0; i < list.Size(); i++)
    {
        std::cout << list.Get(i) << ",";
    }
    std::cout << std::endl;
}

void printTree(std::shared_ptr<scr::BinaryNode> root)
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
}

int main()
{
    scr::BinarySearchTree tree(10);
    tree.Insert(5);
    tree.Insert(20);
    tree.Insert(3);
    tree.Insert(7);
    tree.Insert(15);

    auto root = tree.Root();
    printTree(root);

    // scr::ArrayList list;
    // list.Append(1);
    // list.Append(2);
    // list.Append(4);
    // list.Append(5);
    // printArrayList(list);
    // list.Remove(1);
    // printArrayList(list);
    // list.Insert(2, 6);
    // printArrayList(list);
}