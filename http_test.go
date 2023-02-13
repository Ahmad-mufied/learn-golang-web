package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
! HTTP Test
*--	Go-Lang sudah menyediakan package khusus untuk membuat unit test terhadap fitur Web yang kita buat
*--	Semuanya ada di dalam package net/http/httptest https://golang.org/pkg/net/http/httptest/
*--	Dengan menggunakan package ini, kita bisa melakukan testing handler web di Go-Lang tanpa harus menjalankan aplikasi web nya
*--	Kita bisa langsung fokus terhadap handler function yang ingin kita test

*/

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	/*
	! httptest.NewRequest()
	*--	NewRequest(method, url, body) merupakan function yang digunakan untuk membuat http.Request
	*--	Kita bisa menentukan method, url dan body yang akan kita kirim sebagai simulasi unit test
	*--	Selain itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim, seperti header, cookie, dan lain-lain
	*/
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)

	/*
	! httptest.NewRecorder()
	*--	NewRequest(method, url, body) merupakan function yang digunakan untuk membuat http.Request
	*--	Kita bisa menentukan method, url dan body yang akan kita kirim sebagai simulasi unit test
	*--	Selain itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim, seperti header, cookie, dan lain-lain
	*/
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
