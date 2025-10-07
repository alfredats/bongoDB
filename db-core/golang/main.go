package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var globalBongo *BongoHandle
var MAX_VALUE_SIZE = 1024 // 1KB

type MessageInput struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type MessageResponse struct {
	Value string `json:"value"`
}

func create_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests, please", http.StatusMethodNotAllowed)
		return
	}

	var msgIn MessageInput
	if err := json.NewDecoder(r.Body).Decode(&msgIn); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if msgIn.Key == "" || msgIn.Value == "" {
		http.Error(w, "Missing key or value", http.StatusBadRequest)
		return
	}

	if msgIn.Value != "" && len(msgIn.Value) > MAX_VALUE_SIZE {
		http.Error(w, fmt.Sprintf("Value exceeds maximum size of %d bytes", MAX_VALUE_SIZE), http.StatusBadRequest)
		return
	}

	if result := globalBongo.Create(msgIn.Key, msgIn.Value); result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Create with k: '%s', v: '%s' successful\n", msgIn.Key, msgIn.Value)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func read_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests, please", http.StatusMethodNotAllowed)
		return
	}

	var msgIn MessageInput
	if err := json.NewDecoder(r.Body).Decode(&msgIn); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if msgIn.Key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	readValue, result := globalBongo.Read(msgIn.Key)
	if result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Read with k: '%s', returned '%s' successful\n", msgIn.Key, readValue)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Value: readValue})
}

func update_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests, please", http.StatusMethodNotAllowed)
		return
	}

	var msgIn MessageInput
	if err := json.NewDecoder(r.Body).Decode(&msgIn); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if msgIn.Key == "" || msgIn.Value == "" {
		http.Error(w, "Missing key or value", http.StatusBadRequest)
		return
	}

	if msgIn.Value != "" && len(msgIn.Value) > MAX_VALUE_SIZE {
		http.Error(w, fmt.Sprintf("Value exceeds maximum size of %d bytes", MAX_VALUE_SIZE), http.StatusBadRequest)
		return
	}

	result := globalBongo.Update(msgIn.Key, msgIn.Value)
	if result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Update with k: '%s', v: '%s' successful\n", msgIn.Key, msgIn.Value)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func delete_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests, please", http.StatusMethodNotAllowed)
		return
	}

	var msgIn MessageInput
	if err := json.NewDecoder(r.Body).Decode(&msgIn); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if msgIn.Key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	result := globalBongo.Delete(msgIn.Key)
	if result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Delete with k: '%s' successful\n", msgIn.Key)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	var portnum *int = flag.Int("port", 8080, "Port to run the server on")
	var libPath *string = flag.String("libpath", "", "Path to the shared library")
	flag.Parse()

	if libPath == nil || *libPath == "" {
		log.Fatal("Please provide the path to the shared library using -libpath")
		return
	}

	globalBongo = &BongoHandle{}
	globalBongo.Init(*libPath)

	// Simple HTTP server to demonstrate the API endpoints
	http.HandleFunc("/create", create_handler)
	http.HandleFunc("/read", read_handler)
	http.HandleFunc("/update", update_handler)
	http.HandleFunc("/delete", delete_handler)

	str_portnum := strconv.Itoa(*portnum)
	fmt.Printf("Starting server on :%s\n", str_portnum)
	err := http.ListenAndServe(":"+str_portnum, nil)
	if err != nil {
		panic(err)
	}
}
