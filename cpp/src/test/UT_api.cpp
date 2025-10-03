#include <gtest/gtest.h>
#include "api/api.hpp"

// Test for bongoDB_init
TEST(BongoApiTest, Init) {
    EXPECT_EQ(bongoDB_init(), BONGO_DB_SUCCESS);
}

// Test for bongoDB_create
TEST(BongoApiTest, Create) {
    EXPECT_EQ(bongoDB_create("testKey", "testValue"), BONGO_DB_SUCCESS);
}

// Test for bongoDB_read
TEST(BongoApiTest, Read) {
    bongoDB_create("testKey", "testValue"); // Create a key-value pair first
    char value[256]; // Assuming a buffer size of 256
    EXPECT_EQ(bongoDB_read("testKey", value, sizeof(value)), BONGO_DB_SUCCESS);
    EXPECT_STREQ(value, "testValue"); // Check if the value is correct
}

// Test for bongoDB_update
TEST(BongoApiTest, Update) {
    bongoDB_create("testKey", "testValue");
    EXPECT_EQ(bongoDB_update("testKey", "newValue"), BONGO_DB_SUCCESS);
    char value[256]; // Assuming a buffer size of 256
    bongoDB_read("testKey", value, sizeof(value));
    EXPECT_STREQ(value, "newValue"); // Check if the value is updated
}

// Test for bongoDB_delete
TEST(BongoApiTest, Delete) {
    bongoDB_create("testKey", "testValue");
    EXPECT_EQ(bongoDB_delete("testKey"), BONGO_DB_SUCCESS);
    char* value;
    EXPECT_EQ(bongoDB_read("testKey", value, sizeof(value)), BONGO_DB_NOT_FOUND); // Check if the key is not found
}
