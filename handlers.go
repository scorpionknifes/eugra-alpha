package main

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

//Handler for handling static wav files
func FileServer(w http.ResponseWriter, r *http.Request) {
	//get the file path the user wants to access
	filename := "." + r.URL.Path

	w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(filename)))
	http.ServeFile(w, r, filename)
	log.Print("File Audio Get: " + filename)
}
