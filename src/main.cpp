#include <iostream>

#include "list/array_list.h"

int main()
{
    scr::ArrayList list;
    list.Append(1);
    list.Append(2);
    list.Append(4);
    list.Append(5);
    list.Insert(2, 3);
    for (std::size_t i = 0; i < list.Size(); i++)
    {
        std::cout << list.Get(i) << ",";
    }
    std::cout << std::endl;
}