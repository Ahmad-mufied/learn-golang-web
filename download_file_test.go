package main

import (
	"fmt"
	"net/http"
	"testing"
)

/*

==	Download File

-	Selain upload file, kadang kita ingin membuat halaman website yang digunakan untuk download sesuatu
-	Sebenarnya di Go-Lang sudah disediakan menggunakan FileServer dan ServeFile
-	Dan jika kita ingin memaksa file di download (tanpa di render oleh browser, kita bisa menggunakan header Content-Disposition)
-	https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition

*/

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(w, r, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
