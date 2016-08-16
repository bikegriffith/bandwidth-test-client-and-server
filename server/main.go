package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", handleUpload)
	http.ListenAndServe(":3000", nil)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Echo back the total request size (body + all headers)
	size := (int)(r.ContentLength)
	for k, v := range r.Header {
		size += len(k) + len(v)
	}
	response := fmt.Sprintf("size=%d", size)
	w.Write([]byte(response))
}
