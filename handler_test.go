package main

import (
	"fmt"
	"net/http"
	"testing"
)

/*
! Handler
*-- Server hanya bertugas sebagai Web Server, sedangkan untuk menerima HTTP Request yang masuk ke Server,
*	kita butuh yang namanya Handler
*-- Handler di Go-Lang di representasikan dalam interface, dimana dalam kontraknya terdapat sebuah function bernama ServeHTTP()
*	yang digunakan sebagai function yang akan di eksekusi ketika menerima HTTP Request

!HandlerFunc
*--	Salah satu implementasi dari interface Handler adalah HandlerFunc
*--	Kita bisa menggunakan HandlerFunc untuk membuat function handler HTTP


*/

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

/*
! ServerMux
*--	Saat membuat web, kita biasanya ingin membuat banyak sekali endpoint URL
*--	HandlerFunc sayangnya tidak mendukung itu
*--	Alternative implementasi dari Handler adalah ServeMux
*--	ServeMux adalah implementasi Handler yang bisa mendukung multiple endpoint

! URL Pattern
*--	URL Pattern dalam ServeMux sederhana, kita tinggal menambahkan string yang ingin kita gunakan sebagai endpoint,
*	tanpa perlu memasukkan domain web kita
*--	Jika URL Pattern dalam ServeMux kita tambahkan di akhirnya dengan garis miring, artinya semua url tersebut akan
*	menerima path dengan awalan tersebut, misal /images/ artinya akan dieksekusi jika endpoint nya /images/, /images/contoh, /images/contoh/lagi
*--	Namun jika terdapat URL Pattern yang lebih panjang, maka akan diprioritaskan yang lebih panjang, misal jika terdapat URL /images/
*	dan /images/thumbnails/, maka jika mengakses /images/thumbnails/ akan mengakses /images/thumbnails/, bukan /images

*/

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

/*
! Request
*--	Request adalah struct yang merepresentasikan HTTP Request yang dikirim oleh Web Browser
*--	Semua informasi request yang dikirim bisa kita dapatkan di Request
*--	Seperti, URL, http method, http header, http body, dan lain-lain

*/

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
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
