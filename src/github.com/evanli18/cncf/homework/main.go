package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	// handle
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		for s := range r.Header {
			fmt.Fprintf(w, "%s: %v\n", s, r.Header[s])
		}
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		var version string
		version = os.Getenv("VERSION")
		// default version
		if version == "" {
			version = "0.1"
		}
		fmt.Fprintf(w, "version: %s", version)
	})

	s := &http.Server{
		Addr: ":8080",
	}
	log.Fatal(s.ListenAndServe())
}
