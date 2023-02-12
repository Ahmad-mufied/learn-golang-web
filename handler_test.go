package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writter http.ResponseWriter, response *http.Request) {
		// Logic Web
		fmt.Fprint(writter, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Image")
	})
	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
