package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*

==	Template Function

-	Selain mengakses field, dalam template, function juga bisa diakses
-	Cara mengakses function sama seperti mengakses field, namun jika function tersebut memiliki parameter, kita bisa
	tambahkan parameter ketika memanggil function di template nya
-	{{.FunctionName}}, memanggil field FunctionName atau function FunctionName()
-	{{.FunctionName “eko”, “kurniawan”}}, memanggil function FunctionName(“eko”, “kurniawan”)


*/

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Ahmad",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Ahmad",
	})
}

/*

-- Global Function

-	Go-Lang template memiliki beberapa global function
-	Global function adalah function yang bisa digunakan secara langsung, tanpa menggunakan template data
-	Berikut adalah beberapa global function di Go-Lang template
-	https://github.com/golang/go/blob/master/src/text/template/funcs.go


*/

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*

--	Menambah Global Function

-	Kita juga bisa menambah global function
-	Untuk menambah global function, kita bisa menggunakan method Funcs pada template
-	Perlu diingat, bahwa menambahkan global function harus dilakukan sebelum melakukan parsing template

*/

func TemplateCreateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Ahmad Mufied Nugroho",
	})
}

func TestTemplateCreateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCreateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/*

--	Function Pipelines

-	Go-Lang template mendukung function pipelines, artinya hasil dari function bisa dikirim ke function berikutnya
-	Untuk menggunakan function pipelines, kita bisa menggunakan tanda | , misal :
-	{{ sayHello .Name | upper }}, artinya akan memanggil global function sayHello(Name) hasil dari sayHello(Name) akan dikirim ke function upper(hasil)
-	Kita bisa menambahkan function pipelines lebih dari satu


*/

func TemplateCreateFunctionGlobalPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Ahmad Mufied Nugroho",
	})
}

func TestTemplateCreateFunctionGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCreateFunctionGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
