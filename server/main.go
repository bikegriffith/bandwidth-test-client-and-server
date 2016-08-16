package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(os.Getenv("ASSET_ROOT")))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/upload", handleUpload)

	log.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", mux)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Echo back the total request size (body + all headers)
	size := (int)(r.ContentLength)
	for k, v := range r.Header {
		size += len(k) + len(v)
	}
	log.Println("Processed upload request of size", size)
	response := fmt.Sprintf("size=%d", size)
	w.Write([]byte(response))
}
