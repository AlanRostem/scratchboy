#pragma once
#include <stdint.h>

#include "list/array_list.h"

namespace scr
{
    template<typename T>
    size_t FNVHash(T any)
    {
        constexpr size_t FNVOffsetBasis = 0xcbf29ce484222325;
        constexpr size_t FNVPrime = 0x100000001b3;

        auto hash = FNVOffsetBasis;
        uint8_t *byte = reinterpret_cast<uint8_t*>(&any);
        size_t count = sizeof(T);
        for (std::size_t i = 0; i < count; i++)
        {
            hash *= FNVPrime;
            hash ^= byte[i];
        }

        return hash;
    }
}