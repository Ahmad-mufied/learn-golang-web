package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*

==	Template Data
-	Saat kita membuat template, kadang kita ingin menambahkan banyak data dinamis
-	Hal ini bisa kita lakukan dengan cara menggunakan data struct atau map
-	Namun perlu dilakukan perubahan di dalam text template nya, kita perlu memberi tahu Field atau Key mana yang akan kita gunakan untuk mengisi data dinamis di template
-	Kita bisa menyebutkan dengan cara seperti ini {{.NamaField}}


*/

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Ahmad",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum ada",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Mufied",
		Address: Address{
			Street: "Jalan Baru",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
