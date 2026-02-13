package main

import (
	// standard library
	"fmt"
	"log"
	"net/http"

	// own custom modules
	"cheatsheet/api"
)

func main() {
	fs_ui := http.FileServer(http.Dir("../ui"))
	fs_frontend := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs_ui)
	http.Handle("/frontend/", http.StripPrefix("/frontend/", fs_frontend))

	http.HandleFunc("/upload", api.HandleUpload)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
