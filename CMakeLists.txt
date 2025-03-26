cmake_minimum_required(VERSION 3.0) 
project(LibScratch)

set(CXX_STANDARD 17)
set(PROJECT_SOURCES src/main.cpp)
include_directories(${PROJECT_NAME} PUBLIC src/main)

add_executable(${PROJECT_NAME} ${PROJECT_SOURCES})

target_compile_options(${PROJECT_NAME} PUBLIC -std=c++17 -Wall -Wfloat-conversion)