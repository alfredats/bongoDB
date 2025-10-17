package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alfredats/bongoDB/db-core/golang/src"
)

func TestCreateHandler(t *testing.T) {
	testLibPath := "libbongoDB-cpp.so"
	instance := &src.BongoInstance{}
	instance.Start(0, testLibPath) // Initialize the instance and handle

	// Test valid POST request
	msg := src.MessageInput{Key: "testKey", Value: "testValue"}
	body, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
	w := httptest.NewRecorder()

	instance.Create_handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	// Test invalid method
	req = httptest.NewRequest(http.MethodGet, "/create", nil)
	w = httptest.NewRecorder()
	instance.Create_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 Method Not Allowed, got %d", resp.StatusCode)
	}

	// Test invalid JSON
	req = httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader([]byte("invalid json")))
	w = httptest.NewRecorder()
	instance.Create_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for invalid JSON, got %d", resp.StatusCode)
	}

	// Test missing key or value
	msg = src.MessageInput{Key: "", Value: ""}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
	w = httptest.NewRecorder()
	instance.Create_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing key or value, got %d", resp.StatusCode)
	}
}

func TestReadHandler(t *testing.T) {
	testLibPath := "libbongoDB-cpp.so"
	instance := &src.BongoInstance{}
	instance.Start(0, testLibPath)

	// Test valid POST request
	msg := src.MessageInput{Key: "testKey"}
	body, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, "/read", bytes.NewReader(body))
	w := httptest.NewRecorder()

	instance.Read_handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	// Test invalid method
	req = httptest.NewRequest(http.MethodGet, "/read", nil)
	w = httptest.NewRecorder()
	instance.Read_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 Method Not Allowed, got %d", resp.StatusCode)
	}

	// Test invalid JSON
	req = httptest.NewRequest(http.MethodPost, "/read", bytes.NewReader([]byte("invalid json")))
	w = httptest.NewRecorder()
	instance.Read_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for invalid JSON, got %d", resp.StatusCode)
	}

	// Test missing key
	msg = src.MessageInput{Key: ""}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/read", bytes.NewReader(body))
	w = httptest.NewRecorder()
	instance.Read_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing key, got %d", resp.StatusCode)
	}
}

func TestUpdateHandler(t *testing.T) {
	testLibPath := "libbongoDB-cpp.so"
	instance := &src.BongoInstance{}
	instance.Start(0, testLibPath)

	// Test valid POST request
	msg := src.MessageInput{Key: "testKey", Value: "updatedValue"}
	body, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, "/update", bytes.NewReader(body))
	w := httptest.NewRecorder()

	instance.Update_handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	// Test invalid method
	req = httptest.NewRequest(http.MethodGet, "/update", nil)
	w = httptest.NewRecorder()
	instance.Update_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 Method Not Allowed, got %d", resp.StatusCode)
	}

	// Test invalid JSON
	req = httptest.NewRequest(http.MethodPost, "/update", bytes.NewReader([]byte("invalid json")))
	w = httptest.NewRecorder()
	instance.Update_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for invalid JSON, got %d", resp.StatusCode)
	}

	// Test missing key or value
	msg = src.MessageInput{Key: "", Value: ""}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/update", bytes.NewReader(body))
	w = httptest.NewRecorder()
	instance.Update_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing key or value, got %d", resp.StatusCode)
	}

	// Edge case: Empty value string (should be rejected)
	msg = src.MessageInput{Key: "testKey", Value: ""}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/update", bytes.NewReader(body))
	w = httptest.NewRecorder()
	instance.Update_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for empty value, got %d", resp.StatusCode)
	}

	// Edge case: Very large value string
	largeValue := make([]byte, 10*1024*1024) // 10MB of zeros
	for i := range largeValue {
		largeValue[i] = 'a'
	}
	msg = src.MessageInput{Key: "testKey", Value: string(largeValue)}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/update", bytes.NewReader(body))
	w = httptest.NewRecorder()
	instance.Update_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for large value, got %d", resp.StatusCode)
	}
}

func TestDeleteHandler(t *testing.T) {
	testLibPath := "libbongoDB-cpp.so"
	instance := &src.BongoInstance{}
	instance.Start(0, testLibPath)

	// Test valid POST request
	msg := src.MessageInput{Key: "testKey"}
	body, _ := json.Marshal(msg)
	req := httptest.NewRequest(http.MethodPost, "/delete", bytes.NewReader(body))
	w := httptest.NewRecorder()

	instance.Delete_handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	// Test invalid method
	req = httptest.NewRequest(http.MethodGet, "/delete", nil)
	w = httptest.NewRecorder()
	instance.Delete_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 Method Not Allowed, got %d", resp.StatusCode)
	}

	// Test invalid JSON
	req = httptest.NewRequest(http.MethodPost, "/delete", bytes.NewReader([]byte("invalid json")))
	w = httptest.NewRecorder()
	instance.Delete_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for invalid JSON, got %d", resp.StatusCode)
	}

	// Test missing key
	msg = src.MessageInput{Key: ""}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/delete", bytes.NewReader(body))
	w = httptest.NewRecorder()
	instance.Delete_handler(w, req)
	resp = w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request for missing key, got %d", resp.StatusCode)
	}

	// Edge case: Delete non-existent key (assuming Delete returns non-zero for failure)
	msg = src.MessageInput{Key: "nonExistentKey"}
	body, _ = json.Marshal(msg)
	req = httptest.NewRequest(http.MethodPost, "/delete", bytes.NewReader(body))
	w = httptest.NewRecorder()
	instance.Delete_handler(w, req)
	resp = w.Result()
	// We expect either 200 OK or 500 Internal Server Error depending on implementation
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 200 OK or 500 Internal Server Error for non-existent key, got %d", resp.StatusCode)
	}
}
