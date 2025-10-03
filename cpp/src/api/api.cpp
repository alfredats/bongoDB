#include "api/api.hpp"
#include "main/bongo.hpp"
#include <iostream>
#include <cstring>


// init & shutdown
int bongoDB_init() {
    BongoDB::getInstance();
    return BONGO_DB_SUCCESS;
};

// crud
int bongoDB_create(const char* key, const char* value) {
    BongoDB& db = BongoDB::getInstance();
    return db.kv_create(std::string(key, strlen(key)), std::string(value, strlen(value)));
};

int bongoDB_read(const char* search_fmt, const char* value) {
    BongoDB& db = BongoDB::getInstance();
    std::string outval = "";
    int result = db.kv_read(std::string(search_fmt, strlen(search_fmt)), outval);
    if (result == BONGO_DB_SUCCESS) { value = outval.c_str(); }
    return result;
};

int bongoDB_update(const char* key, const char* value) {
    BongoDB& db = BongoDB::getInstance();
    std::string outval = "";
    return db.kv_update(std::string(key, strlen(key)), std::string(value, strlen(value)), outval);
};

int bongoDB_delete(const char* key) {
    BongoDB& db = BongoDB::getInstance();
    return db.kv_delete(std::string(key, strlen(key)));
};