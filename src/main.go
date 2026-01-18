package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

/*
type fileRequest struct {
	
}
*/

type Response struct {
	success string `json:"result"`
	Error string `json:"err,omitempty"`
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json") // Tell the client we're sending JSON

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var req fileRequest

	json.NewEncoder(w).Encode(Response{success: "Success!"})
}

func main() {
	http.HandleFunc("/upload", handleUpload)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
