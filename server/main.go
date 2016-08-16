package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/upload", addDefaultHeaders(handleUpload))
	mux.HandleFunc("/download", addDefaultHeaders(handleDownload))

	staticRoot := os.Getenv("STATIC_ROOT")
	mux.Handle("/", http.FileServer(http.Dir(staticRoot)))

	log.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", mux)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Read in the HTTP body (up to limit)
	maxUploadSizeMB, _ := strconv.Atoi(os.Getenv("MAX_UPLOAD_MB"))
	r.Body = http.MaxBytesReader(w, r.Body, (int64)(maxUploadSizeMB*1024*1024))
	size, _ := io.Copy(ioutil.Discard, r.Body)
	expected := r.ContentLength
	log.Println("Processed upload request of size", size, "and expected", expected)
	response := fmt.Sprintf("OK")
	w.Write([]byte(response))
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	// Generate random bytes of given length
	w.Header().Set("Content-Type", "application/octet-stream")
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	b := make([]byte, 1)
	for i := 0; i < size; i++ {
		b[0] = (byte)(rand.Int())
		w.Write(b)
	}
	log.Println("Processed download request of size", size)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add CORS support for browser-based clients
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Length, Accept-Encoding, Authorization")
		// Ensure browser and proxies do not try to cache this; we need a full test.
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		fn(w, r)
	}
}
