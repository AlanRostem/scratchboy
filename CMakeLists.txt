cmake_minimum_required(VERSION 3.0) 
project(ScratchBoy)

set(CXX_STANDARD 17)
set(PROJECT_SOURCES 
    src/main.cpp
    src/allocator/heap.cpp
    src/cpu/cpu.cpp
    src/cpu/register_bank.cpp
    src/cpu/memory_bus.cpp
    src/cpu/instructions/add.cpp
)
include_directories(${PROJECT_NAME} PUBLIC src/)

if (SCRATCH_DEBUG)
    add_compile_definitions(_DEBUG)
endif()

add_executable(${PROJECT_NAME} ${PROJECT_SOURCES})

target_compile_options(${PROJECT_NAME} PUBLIC -std=c++17 -Wall -Wfloat-conversion -Werror -g)