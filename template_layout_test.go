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

== Template Layout

-	Saat kita membuat halaman website, kadang ada beberapa bagian yang selalu sama, misal header dan footer
-	Best practice nya jika terdapat bagian yang selalu sama, disarankan untuk disimpan pada template yang terpisah, agar bisa digunakan di template lain
-	Go-Lang template mendukung import dari template lain


--	Import Template

Untuk melakukan import, kita bisa menggunakan perintah berikut :
-	{{template “nama”}}, artinya kita akan meng-import template “nama” tanpa memberikan data apapun
-	{{template “nama” .Value}}, artinya kita akan meng-import template “nama” dengan memberikann data value


--	Template Name

-	Saat kita membuat template dari file, secara otomatis nama file nya akan menjadi nama template
-	Namun jika kita ingin mengubah nama template nya, kita juga bisa melakukan mengguakan perintah {{define “nama”}} TEMPLATE {{end}}, artinya kita membuat template dengan nama “nama”



*/

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))
	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Ahmad",
	})
}

func TestTemplateTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
