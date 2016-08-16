package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	// TODO:
	// * add max upload size limit (to prevent dos)
	// * add HTTP headers to prevent caching and compression (client Accept-Encoding: identity,
	//   server Cache-Control: no-cache)
	// * add simple authentication (but no SSL/TLS) to prevent abuse
	mux.HandleFunc("/upload", addDefaultHeaders(handleUpload))
	mux.HandleFunc("/download", addDefaultHeaders(handleDownload))

	mux.Handle("/", http.FileServer(http.Dir(os.Getenv("STATIC_ROOT"))))

	log.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", mux)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Echo back the total request size (body + all headers)
	size := (int)(r.ContentLength)
	log.Println("Processed upload request of size", size)
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
