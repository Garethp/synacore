cmake_minimum_required(VERSION 3.2)
project(synacore)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11")

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY "~/synacore/")
set(CMAKE_CXX_COMPILER /usr/bin/g++)

set(SOURCE_FILES main.cpp)

add_executable(synacore main.cpp vm.cpp vm.h memory.cpp memory.h registers.cpp registers.h)