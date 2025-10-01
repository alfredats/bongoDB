#include "api/api.hpp"
#include "main/bongo.hpp"
#include <iostream>


// init & shutdown
int bongoDB_init() {
    std::cout << "here be the init call" << std::endl;
    return BONGO_DB_ERROR;
};
int bongoDB_shutdown(){
    std::cout << "here be the shutdown call" << std::endl;
    return BONGO_DB_ERROR;
};

// crud
int bongoDB_create(const char* key, const char* value) {
    std::cout << "here be the create call" << std::endl;
    return BONGO_DB_ERROR;
};

int bongoDB_read(const char* search_fmt, const char* value) {
    std::cout << "here be the read call" << std::endl;
    return BONGO_DB_ERROR;
};

int bongoDB_update(const char* key, const char* value) {
    std::cout << "here be the update call" << std::endl;
    return BONGO_DB_ERROR;
};

int bongoDB_delete(const char* key) {
    std::cout << "here be the delete call" << std::endl;
    return BONGO_DB_ERROR;
};


// heartbeat
int bongoDB_heartbeat() {
    std::cout << "here be the heartbeat call" << std::endl;
    return BONGO_DB_ERROR;
};