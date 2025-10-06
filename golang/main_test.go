package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Import the package where the server code is defined
// Assuming the package is main, so no import needed

// To fix undefined errors, import golang files as package or define the test in the same package main

func TestCreateHandler(t *testing.T) {
	globalBongo = &BongoHandle{}         // Mock or initialize as needed
	testLibPath := "libbongoDB-cpp.so"   // Adjust the path as necessary for your test environment
	err := globalBongo.Init(testLibPath) // Use a mock or test library path
	if err != nil {
		t.Fatalf("Failed to initialize globalBongo: %v", err)
	}

	// Test valid POST request
	msg := MessageInput{Key: "testKey", Value: "testValue"}
	body, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
	w := httptest.NewRecorder()

	create_handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	// Test invalid method
	req = httptest.NewRequest(http.MethodGet, "/create", nil)
	w = httptest.NewRecorder()
	create_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 Method Not Allowed, got %d", resp.StatusCode)
	}

	// Test invalid JSON
	req = httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader([]byte("invalid json")))
	w = httptest.NewRecorder()
	create_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for invalid JSON, got %d", resp.StatusCode)
	}

	// Test missing key or value
	msg = MessageInput{Key: "", Value: ""}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
	w = httptest.NewRecorder()
	create_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing key or value, got %d", resp.StatusCode)
	}
}

func TestReadHandler(t *testing.T) {
	globalBongo = &BongoHandle{}         // Mock or initialize as needed
	testLibPath := "libbongoDB-cpp.so"   // Adjust the path as necessary for your test environment
	err := globalBongo.Init(testLibPath) // Use a mock or test library path
	if err != nil {
		t.Fatalf("Failed to initialize globalBongo: %v", err)
	}

	// Test valid POST request
	msg := MessageInput{Key: "testKey"}
	body, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, "/read", bytes.NewReader(body))
	w := httptest.NewRecorder()

	read_handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	// Test invalid method
	req = httptest.NewRequest(http.MethodGet, "/read", nil)
	w = httptest.NewRecorder()
	read_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 Method Not Allowed, got %d", resp.StatusCode)
	}

	// Test invalid JSON
	req = httptest.NewRequest(http.MethodPost, "/read", bytes.NewReader([]byte("invalid json")))
	w = httptest.NewRecorder()
	read_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for invalid JSON, got %d", resp.StatusCode)
	}

	// Test missing key
	msg = MessageInput{Key: ""}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/read", bytes.NewReader(body))
	w = httptest.NewRecorder()
	read_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing key, got %d", resp.StatusCode)
	}
}
