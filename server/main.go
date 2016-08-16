package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/upload", handleUpload)
	mux.HandleFunc("/download", handleDownload)

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

func handleDownload(w http.ResponseWriter, r *http.Request) {
	// Generate fixed string of given length
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	for i := 0; i < size; i++ {
		w.Write([]byte("."))
	}
	log.Println("Processed download request of size", size)
}
