#include <gtest/gtest.h>
#include "api/api.hpp"
#include "main/bongo.hpp"

class BongoDBTest : public ::testing::Test {
protected:
    BongoDB* instance;
    
    void SetUp() override {
        this->instance = &BongoDB::getInstance();
    }

    void TearDown() override {
    }
};

TEST_F(BongoDBTest, CreateTest) {
    std::string key = "testKey";
    std::string value = "testValue";
    EXPECT_EQ(instance->kv_create(key, value), BONGO_DB_SUCCESS);
    EXPECT_EQ(instance->kv_create(key, value), BONGO_DB_ALREADY_EXISTS);
}

TEST_F(BongoDBTest, ReadTest) {
    std::string key = "testKey";
    std::string value;
    EXPECT_EQ(instance->kv_read(key, value), BONGO_DB_SUCCESS);
    EXPECT_EQ(value, "testValue");
    EXPECT_EQ(instance->kv_read("nonExistentKey", value), BONGO_DB_NOT_FOUND);
}

TEST_F(BongoDBTest, UpdateTest) {
    std::string key = "testKey";
    std::string newValue = "newValue";
    std::string oldValue;
    EXPECT_EQ(instance->kv_update(key, newValue, oldValue), BONGO_DB_SUCCESS);
    EXPECT_EQ(oldValue, "testValue");
    EXPECT_EQ(instance->kv_update("nonExistentKey", newValue, oldValue), BONGO_DB_NOT_FOUND);
}

TEST_F(BongoDBTest, DeleteTest) {
    std::string key = "testKey";
    EXPECT_EQ(instance->kv_delete(key), BONGO_DB_SUCCESS);
    EXPECT_EQ(instance->kv_delete("nonExistentKey"), BONGO_DB_NOT_FOUND);
}
