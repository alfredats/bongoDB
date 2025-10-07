#include "main/bongo.hpp"
#include "api/api.hpp"

#include <uuid/uuid.h>

BongoDB::BongoDB() : collection() {
    uuid_t id;
    uuid_generate_random(id);
    char uid_str[36];
    uuid_unparse(id, uid_str);
    this->instanceID = std::string(uid_str, 36);
};

BongoDB& BongoDB::getInstance() {
    static BongoDB instance; 
    return instance;
};

std::string BongoDB::getInstanceID() {
    return this->instanceID;
}

int BongoDB::kv_create(const std::string& key, const std::string& value) {
    auto exists = this->collection.find(key);
    if (exists != this->collection.end()) {
        return BONGO_DB_ALREADY_EXISTS;
    }
    this->collection.emplace(std::make_pair(key, value));
    return BONGO_DB_SUCCESS;
};

int BongoDB::kv_read(const std::string& key, std::string& out_val) {
    auto exists = this->collection.find(key);
    if (exists == this->collection.end()) {
        return BONGO_DB_NOT_FOUND;
    }
    out_val = exists->second;
    return BONGO_DB_SUCCESS;
};

int BongoDB::kv_update(const std::string& key, const std::string& updateValue, std::string& oldValue) {
    auto exists = this->collection.find(key);
    if (exists == this->collection.end()) {
        return BONGO_DB_NOT_FOUND;
    }
    oldValue = exists->second;
    this->collection[key] = updateValue;
    return BONGO_DB_SUCCESS;
};

int BongoDB::kv_delete(const std::string& key) {
    auto exists = this->collection.find(key);
    if (exists == this->collection.end()) {
        return BONGO_DB_NOT_FOUND;
    }
    this->collection.erase(key);
    return BONGO_DB_SUCCESS;
};
