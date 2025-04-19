#pragma once

#include "types/word.h"

namespace scr
{
    struct CPU;
    using Instruction = void(*)(CPU*, Word);
}