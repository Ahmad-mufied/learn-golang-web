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

==	Template

--	Web Dinamis

-	Sampai saat ini kita hanya membahas tentang membuat response menggunakan String dan juga static file
-	Pada kenyataannya, saat kita membuat web, kita pasti akan membuat halaman yang dinamis, bisa berubah-ubah sesuai dengan data yang diakses oleh user
-	Di Go-Lang terdapat fitur HTML Template, yaitu fitur template yang bisa kita gunakan untuk membuat HTML yang dinamis

--	HTML Template

-	Fitur HTML template terdapat di package html/template
-	Sebelum menggunakan HTML template, kita perlu terlebih dahulu membuat template nya
-	Template bisa berubah file atau string
-	Bagian dinamis pada HTML Template, adalah bagian yang menggunakan tanda {{  }}

--	Membuat Template

-	Saat membuat template dengan string, kira perlu memberi tahu nama template nya
-	Dan untuk membuat text template, cukup buat text html, dan untuk konten yang dinamis, kita bisa gunakan tanda {{.}}, contoh :
-	<html><body>{{.}}</body></html>



*/

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//t, err := template.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*

--	Template Dari File

-	Selain membuat template dari string, kita juga bisa membuat template langsung dari file
-	Hal ini mempermudah kita, karena bisa langsung membuat file html
-	Saat membuat template menggunakan file, secara otomatis nama file akan menjadi nama template nya,
	misal jika kita punya file simple.html, maka nama template nya adalah simple.html
*/

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*

--	Template Directory

-	Kadang biasanya kita jarang sekali menyebutkan file template satu persatu
-	Alangkah baiknya untuk template kita simpan di satu directory
-	Go-Lang template mendukung proses load template dari directory
-	Hal ini memudahkan kita, sehingga tidak perlu menyebutkan nama file nya satu per satu

*/

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*

--	Template dari Go-Lang Embed

-	Sejak Go-Lang 1.16, karena sudah ada Go-Lang Embed, jadi direkomendasikan menggunakan Go-Lang embed untuk menyimpan data template
-	Menggunakan Go-Lang embed menjadi kita tidak perlu ikut meng-copy template file lagi, karena sudah otomatis di embed di dalam distribution file

*/

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
