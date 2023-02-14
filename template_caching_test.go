package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*

==	Template Caching

-	Kode-kode diatas yang sudah kita praktekan sebenarnya tidak efisien
-	Hal ini dikarenakan, setiap Handler dipanggil, kita selalu melakukan parsing ulang template nya
-	Idealnya template hanya melakukan parsing satu kali diawal ketika aplikasinya berjalan
-	Selanjutnya data template akan di caching (disimpan di memory), sehingga kita tidak perlu melakukan parsing lagi
-	Hal ini akan membuat web kita semakin cepat

*/

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
