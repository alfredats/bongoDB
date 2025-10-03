#pragma once

#include <iostream>

#define BONGO_DB_SUCCESS 0
#define BONGO_DB_ERROR -1
#define BONGO_DB_NOT_FOUND -2
#define BONGO_DB_ALREADY_EXISTS -3
#define BONGO_DB_INVALID_ARGUMENT -4
#define BONGO_DB_OPERATION_FAILED -5

extern "C" {

// init & shutdown
int bongoDB_init();

// crud
int bongoDB_create(const char* key, const char* value);
int bongoDB_read(const char* search_fmt, const char* value);
int bongoDB_update(const char* key, const char* value);
int bongoDB_delete(const char* key);
}
