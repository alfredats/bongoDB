package src

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var MAX_VALUE_SIZE = 1024 // 1KB

type MessageInput struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type MessageResponse struct {
	Value string `json:"value"`
}

type BongoInstance struct {
	ID     int
	Port   int
	Handle BongoHandle
	mux    *http.ServeMux
}

func (bi *BongoInstance) Create_handler(w http.ResponseWriter, r *http.Request) {
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

	if result := bi.Handle.Create(msgIn.Key, msgIn.Value); result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Create with k: '%s', v: '%s' successful\n", msgIn.Key, msgIn.Value)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (bi *BongoInstance) Read_handler(w http.ResponseWriter, r *http.Request) {
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

	readValue, result := bi.Handle.Read(msgIn.Key)
	if result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Read with k: '%s', returned '%s' successful\n", msgIn.Key, readValue)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Value: readValue})
}

func (bi *BongoInstance) Update_handler(w http.ResponseWriter, r *http.Request) {
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

	result := bi.Handle.Update(msgIn.Key, msgIn.Value)
	if result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Update with k: '%s', v: '%s' successful\n", msgIn.Key, msgIn.Value)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (bi *BongoInstance) Delete_handler(w http.ResponseWriter, r *http.Request) {
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

	result := bi.Handle.Delete(msgIn.Key)
	if result != 0 {
		http.Error(w, fmt.Sprintf("Create failed with code %d", result), http.StatusInternalServerError)
		return
	}
	log.Printf("Delete with k: '%s' successful\n", msgIn.Key)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (bi *BongoInstance) Start(port int, libPath string) {
	bi.Handle.Init(libPath)
	bi.mux = http.NewServeMux()

	// Simple HTTP server to demonstrate the API endpoints
	bi.mux.HandleFunc("/create", bi.Create_handler)
	bi.mux.HandleFunc("/read", bi.Read_handler)
	bi.mux.HandleFunc("/update", bi.Update_handler)
	bi.mux.HandleFunc("/delete", bi.Delete_handler)

	str_portnum := strconv.Itoa(port)
	fmt.Printf("Starting server on :%s\n", str_portnum)
	err := http.ListenAndServe(":"+str_portnum, nil)
	if err != nil {
		panic(err)
	}
}
