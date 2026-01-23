
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	http.HandleFunc("/api/files", listPDF)
	http.HandleFunc("/api/files/", showPDF)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", enableCORS(http.DefaultServeMux))
}

func listPDF(w http.ResponseWriter, r *http.Request) {
	files, _ := os.ReadDir("./files")
	var pdfList []File
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".pdf") {
			pdfList = append(pdfList, File{Name: f.Name(), Type: "pdf"})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pdfList)
}

func showPDF(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/api/files")
	filepath := filepath.Join("./files", filename)

	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filepath)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
