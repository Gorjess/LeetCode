cmake_minimum_required(VERSION 3.22)
project(dijkstras_algorithm)

set(CMAKE_CXX_STANDARD 20)

add_executable(dijkstras_algorithm dijkstras-algorithm.cpp random_engine.h graph.h alias.h)
