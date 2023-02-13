package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
! Stateless
*--	HTTP merupakan stateless antara client dan server, artinya server tidak akan menyimpan data apapun untuk mengingat setiap request dari client
*--	Hal ini bertujuan agar mudah melakukan scalability di sisi server
*--	Lantas bagaimana caranya agar server bisa mengingat sebuah client? Misal ketika kita sudah login di website, server otomatis harus tahu
*	jika client tersebut sudah login, sehingga request selanjutnya, tidak perlu diminta untuk login lagi
*--	Untuk melakukan hal ini, kita bisa memanfaatkan Cookie

! Cookie
*--	Cookie adalah fitur di HTTP dimana server bisa memberi response cookie (key-value) dan client akan menyimpan cookie tersebut di web browser
*--	Request selanjutnya, client akan selalu membawa cookie tersebut secara otomatis
*--	Dan server secara otomatis akan selalu menerima data cookie yang dibawa oleh client setiap kalo client mengirimkan request

*/

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	/*
	! Membuat Cookie
	*--	Cookie merupakan data yang dibuat di server dan sengaja agar disimpan di web browser
	*--	Untuk membuat cookie di server, kita bisa menggunakan function http.SetCookie()
	*/
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Create Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Ahmad", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Ahmad"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)
	
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
