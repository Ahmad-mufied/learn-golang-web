package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*

==	XSS (Cross Site Scripting)

-	XSS adalah salah satu security issue yang biasa terjadi ketika membuat web
-	XSS adalah celah keamanan, dimana orang bisa secara sengaja memasukkan parameter yang mengandung JavaScript agar dirender oleh halaman website kita
-	Biasanya tujuan dari XSS adalah mencuri cookie browser pengguna yang sedang mengakses website kita
-	XSS bisa menyebabkan account pengguna kita diambil alih jika tidak ditangani dengan baik

--	Auto Escape

-	Berbeda dengan bahasa pemrograman lain seperti PHP, pada Go-Lang template, masalah XSS sudah diatasi secara otomatis
-	Go-Lang template memiliki fitur Auto Escape, dimana dia bisa mendeteksi data yang perlu ditampilkan di template, jika mengandung tag-tag html atau script, secara otomatis akan di escape
-	Semua function escape bisa diliat disini :
-	https://github.com/golang/go/blob/master/src/html/template/escape.go 
-	https://golang.org/pkg/html/template/#hdr-Contexts 



*/

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini adalah Body<script>alert('Anda di Hack')</script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*

--	Mematikan Auto Escape

-	Jika kita mau, auto escape juga bisa kita matikan
-	Namun, kita perlu memberi tahu template secara eksplisit ketika kita menambahkan template data
-	Kita bisa menggunakan data 
-	template.HTML , jika ini adalah data html
-	template.CSS, jika ini adalah data css
-	template.JS, jika ini adalah data javascript

*/

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<p>Ini adalah Body</p>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
--	Masalah XSS (Cross Site Scripting)

-	Saat kita mematikan fitur auto escape, bisa dipastikan masalah XSS akan mengintai kita
-	Jadi pastikan kita benar-benar percaya terhadap sumber data yang kita matikan auto escape nya


*/

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>aler</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
