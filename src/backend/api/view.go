package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func ListPDF(w http.ResponseWriter, r *http.Request) {
	files, _ := os.ReadDir("../../storage/media")
	var pdfList []File
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".pdf") {
			pdfList = append(pdfList, File{Name: f.Name(), Type: "pdf"})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pdfList)
}

func ShowPDF(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/api/files")
	filePath := filepath.Join("../../storage/media", filename)

	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filePath)
}
