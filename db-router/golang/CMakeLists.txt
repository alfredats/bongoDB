cmake_minimum_required(VERSION 3.28)

# Set the project name
project(bongosh)

# Set the Go source directory
set(GO_SRC_DIR ${CMAKE_CURRENT_SOURCE_DIR})

# Add a custom target to build the Go project
add_custom_target(bongosh_build ALL
  COMMAND go build -o ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/bongosh ${GO_SRC_DIR}
  WORKING_DIRECTORY ${GO_SRC_DIR}
  COMMENT "Building Go project"
  VERBATIM
)

# Add a custom target to clean the Go build output
add_custom_target(bongosh_clean
  COMMAND rm -f ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/bongosh
  COMMENT "Cleaning Go build output"
)
