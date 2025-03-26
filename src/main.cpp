#include <iostream>

#include "list/array_list.h"

void printArrayList(const scr::ArrayList& list)
{
    for (std::size_t i = 0; i < list.Size(); i++)
    {
        std::cout << list.Get(i) << ",";
    }
    std::cout << std::endl;
}

int main()
{
    scr::ArrayList list;
    list.Append(1);
    list.Append(2);
    list.Append(4);
    list.Append(5);

    printArrayList(list);

    list.Remove(1);

    printArrayList(list);

    list.Insert(2, 6);

    printArrayList(list);
}