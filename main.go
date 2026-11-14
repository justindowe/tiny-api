package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var version string

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	version = os.Getenv("VERSION")
	if version == "" {
		version = "dev"
	}

	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/version", versionHandler)

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
