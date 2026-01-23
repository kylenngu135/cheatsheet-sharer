package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"mime/multipart"
	// "io"
)

type UploadRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
	Description string
	Title string
}

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

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err :=  r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadReq := UploadRequest{
		File: handler,
		Description: r.FormValue("dsecription"),
		Title: r.FormValue("title"),
	}

	fmt.Println(uploadReq)

	/*
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	*/

	// fmt.Println(fileBytes)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{success: "Success!"})
}

func main() {
	http.HandleFunc("/upload", handleUpload)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
