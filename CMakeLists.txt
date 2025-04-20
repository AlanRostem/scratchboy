cmake_minimum_required(VERSION 3.20) 
project(ScratchBoy)

set(CXX_STANDARD 17)
set(API_SOURCES 
    src/allocator/heap.cpp
    src/cpu/cpu.cpp
    src/cpu/register_file.cpp
    src/cpu/memory_bus.cpp
    src/cpu/instructions/add.cpp
)
set(BIN_NAME ${PROJECT_NAME})
set(API_NAME ${PROJECT_NAME}API)

include_directories(${API_NAME} PUBLIC src/)
include_directories(${BIN_NAME} PUBLIC src/)

add_library(${API_NAME} ${API_SOURCES})

target_compile_options(${API_NAME} PUBLIC -std=c++17 -Wall -Wfloat-conversion -Werror)

add_executable(${BIN_NAME} src/main.cpp)

target_link_libraries(${BIN_NAME} ${API_NAME})
target_compile_options(${BIN_NAME} PUBLIC -std=c++17 -Wall -Wfloat-conversion -Werror)

if (SCRATCH_DEBUG)
    target_compile_definitions(${BIN_NAME} PUBLIC SCR_DEBUG)
    target_compile_definitions(${API_NAME} PUBLIC SCR_DEBUG)
    target_compile_options(${BIN_NAME} PUBLIC -g)
    target_compile_options(${API_NAME} PUBLIC -g)
endif()