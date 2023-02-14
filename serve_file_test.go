package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

/*

ServeFile

--	Kadang ada kasus misal kita hanya ingin menggunakan static file sesuai dengan yang kita inginkan
--	Hal ini bisa dilakukan menggunakan function http.ServeFile()
--	Dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke http response


*/

func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*

Go-Lang Embed

--	Parameter function http.ServeFile hanya berisi string file name, sehingga tidak bisa menggunakan Go-Lang Embed
--	Namun bukan berarti kita tidak bisa menggunakan Go-Lang embed, karena jika untuk melakukan load file, kita hanya
	butuh menggunakan package fmt dan ResponseWriter saja

*/

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
