package main

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	/*
		! Server

		* Server adalah struct yang terdapat di package net/http yang digunakan sebagai representasi Web Server di Go-Lang
		* Untuk membuat web, kita wajib membuat Server
		* Saat membuat data Server, ada beberapa hal yang perlu kita tentukan, seperti host dan juga port tempat dimana Web kita berjalan
		* Setelah membuat Server, kita bisa menjalankan Server tersebut menggunakan function ListenAndServe()

	*/
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
