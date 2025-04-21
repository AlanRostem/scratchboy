#pragma once

#include "types/word.h"

namespace scr::opcodes
{
    constexpr Word PREFIX_CB = 0xCB;
    constexpr Word ADD_A_B = 0x80;
    constexpr Word ADD_A_C = 0x81;
    constexpr Word ADD_A_D = 0x82;
    constexpr Word ADD_A_E = 0x83;
    constexpr Word ADD_A_H = 0x84;
    constexpr Word ADD_A_L = 0x85;
    constexpr Word ADD_A_HL = 0x86;
    constexpr Word ADD_A_A = 0x87;

    constexpr Word RL_A = 0x17;
}