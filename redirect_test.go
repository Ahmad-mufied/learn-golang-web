package main

import (
	"fmt"
	"net/http"
	"testing"
)

/*
==	Redirect

-	Saat kita membuat website, kadang kita butuh melakukan redirect
-	Misal setelah selesai login, kita lakukan redirect ke halaman dashboard
-	Redirect sendiri sebenarnya sudah standard di HTTP https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections
-	Kita hanya perlu membuat response code 3xx dan menambah header Location
-	Namun untungnya di Go-Lang, ada function yang bisa kita gunakan untuk mempermudah ini

*/

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	// logic
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(w http.ResponseWriter, r *http.Request) {
	// logic
	http.Redirect(w, r, "https://www.google.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
