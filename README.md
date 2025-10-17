# bongoDB

bongoDB is a project consisting of multiple components primarily written in Go and C++. It includes a core database engine implemented in C++ and Go, as well as a router component written in Go. The project is designed to provide a high-performance, scalable database solution.

## High 

## Project Structure

- `db-core/`: Core database engine implementation
  - `cpp/`: C++ source code and headers for the core engine
  - `golang/`: Go bindings and related code for the core engine
- `db-router/`: Router component implemented in Go
- `build/`: Build artifacts and output
- Root files:
  - `.gitignore`: Git ignore rules
  - `CMakeLists.txt`: Build configuration for C++ components
  - Various Go modules and source files for building and testing

## Features

- High-performance database core implemented in C++
- Go language bindings for easy integration
- Router component for managing database requests
- Unit tests for core components

## High level Overview
```
     ┌───────────────────────────────┐
     │            Router             │
     │                               │
     │  1. Receives query            │
     │  2. Forwards to replicas      │
     │  3. Collects responses        │
     │  4. Quorum consensus          │
     └───────────────────────────────┘
                     │
     ┌───────────────┼────────────────┐
     │               │                │
 ┌────────┐     ┌────────┐      ┌────────┐
 │ Core 1 │     │ Core 2 │      │ Core 3 │
 │  HTTP  │     │  HTTP  │      │  HTTP  │
 └────────┘     └────────┘      └────────┘
```


## Getting Started

### Prerequisites

- C++ compiler supporting C++11 or later
- Go 1.18 or later
- CMake 
- Make

### Building the Project

1. Build the project (C++ & Go components)

```bash
cmake -S. -B./build
cd ./build && cmake --build
```

### Running Tests

Run unit tests for the Go core:

```bash
cd build/bin && BONGO_LIB=../lib ./UT_bongo
```

## Feature Roadmap

1. Replica Instantiation (WIP)
2. Quorum consensus
3. Sharding

as well as performance improvements along the way.