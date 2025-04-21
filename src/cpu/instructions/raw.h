#pragma once

#include "types/word.h"

namespace scr
{
    namespace opcodes
    {
        constexpr Word PrefixCB = 0xCB;
        constexpr Word AddAB = 0x80;
        constexpr Word AddAC = 0x81;
        constexpr Word AddAD = 0x82;
        constexpr Word AddAE = 0x83;
        constexpr Word AddAH = 0x84;
        constexpr Word AddAL = 0x85;
        constexpr Word AddAHL = 0x86;
        constexpr Word AddAA = 0x87;
    }
}