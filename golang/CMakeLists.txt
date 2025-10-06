cmake_minimum_required(VERSION 3.10)

# Set the project name
project(goBongo)

# Set the Go source directory
set(GO_SRC_DIR ${CMAKE_CURRENT_SOURCE_DIR})

# Add a custom target to build the Go project
add_custom_target(build_golang ALL
  COMMAND go build -o ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/goBongo ${GO_SRC_DIR}
  WORKING_DIRECTORY ${GO_SRC_DIR}
  COMMENT "Building Go project"
  VERBATIM
)

# Add a custom target to clean the Go build output
add_custom_target(clean_golang
  COMMAND rm -f ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/goBongo
  COMMENT "Cleaning Go build output"
)
