cmake_minimum_required(VERSION 3.28)

# Set the project name and version
project(bongoDB-cpp VERSION 1.0 LANGUAGES CXX)

# Specify C++17 standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

find_package(PkgConfig REQUIRED)
pkg_search_module(UUID REQUIRED uuid)

# Add executable
add_library(bongoDB-cpp SHARED)
target_include_directories(bongoDB-cpp 
PUBLIC 
    ${CMAKE_CURRENT_SOURCE_DIR}/include
    ${UUID_INCLUDE_DIRS}
)

target_sources(bongoDB-cpp 
PRIVATE 
    ${CMAKE_CURRENT_SOURCE_DIR}/src/api/api.cpp
    ${CMAKE_CURRENT_SOURCE_DIR}/src/main/bongo.cpp
)

target_link_libraries(bongoDB-cpp ${UUID_LIBRARIES})

# Add this line to find Google Test
find_package(GTest REQUIRED)

# Add a test executable
add_executable(UT_bongo) 
target_sources(UT_bongo
PRIVATE 
    src/test/UT_api.cpp
    src/test/UT_bongoMain.cpp
)
target_include_directories(UT_bongo PUBLIC include/main)

# Link Google Test and the bongoDB library
target_link_libraries(UT_bongo GTest::GTest GTest::Main bongoDB-cpp ${UUID_LIBRARIES})
