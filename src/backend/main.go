package main

import (
	// standard library
	"fmt"
	"log"
	"net/http"

	// own custom modules
	"cheatsheet/api"
	"cheatsheet/middleware"
	
)

func main() {
	fs_ui := http.FileServer(http.Dir("../ui"))
	fs_frontend := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs_ui)
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs_frontend))

	http.HandleFunc("/upload", api.HandleUpload)
	http.HandleFunc("/api/files", api.ListPDF)
	http.HandleFunc("/api/files/", api.ShowPDF)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.EnableCORS(http.DefaultServeMux)))
}
