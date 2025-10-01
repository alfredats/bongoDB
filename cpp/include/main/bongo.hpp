#pragma once

#include <iostream>
#include <unordered_map>

class BongoDB {
private:
    std::string instanceID;
    std::unordered_map<std::string, std::string> collection;

    BongoDB();
    ~BongoDB() = default;
    BongoDB(const BongoDB&) = delete;
    BongoDB& operator=(const BongoDB&) = delete;

public:
    static BongoDB& getInstance();
    std::string getInstanceID();

    int kv_create(const std::string& key, const std::string& value);
    int kv_read(const std::string& key, std::string& out_val);
    int kv_update(const std::string& key, const std::string& updateValue, std::string& oldValue);
    int kv_delete(const std::string& key);
};