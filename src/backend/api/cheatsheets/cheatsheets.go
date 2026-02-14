package cheatsheets

import (
	"fmt"
	"net/http"
	"mime/multipart"
	"io"
	"os"
	"path/filepath"
)

// stores upload documents request
type UploadRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
	Description string
	Title string
}

// stores the response
type Response struct {
	success string `json:"result"`
	Error string `json:"err,omitempty"`
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json") // Tell the client we're sending JSON

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// set max file size to 32MB
	r.Body = http.MaxBytesReader(w, r.Body, 32<<20)

	// make a local file if it exceeds 5MB
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.MultipartForm.RemoveAll()

	// get the uploaded file from "media"
	uploaded_file, file_header, err := r.FormFile("media")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer uploaded_file.Close()

	// build filepath to store media
	flagStoragePath := "../../storage/media/"
	path := filepath.Join(flagStoragePath, file_header.Filename)

	// create file to save data to
	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// copy file contents from uploaded to new file
	_, err1 := io.Copy(file, uploaded_file)

	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Write([]byte("uploaded"))
}
