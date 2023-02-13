package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
! Query Parameter
*--	Query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
*--	Query parameter biasanya digunakan untuk mengirim data dari client ke server
*--	Query parameter ditempatkan pada URL
*--	Untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya

*/

func sayHello(writer http.ResponseWriter, request *http.Request) {
	/*
	! url.URL
	*--	Dalam parameter Request, terdapat attribute URL yang berisi data url.URL
	*--	Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method Query() yang akan mengembalikan map

	*/
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParamater(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Ahmad", nil)
	recorder := httptest.NewRecorder()

	sayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*
! Multiple Query Parameter
*--	Dalam spesifikasi URL, kita bisa menambahkan lebih dari satu query parameter
*--	Ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query parameter lainnya
*--	Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya
*/

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Ahmad&last_name=Mufied", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*
! Multiple Value Query Parameter
*--	Sebenarnya URL melakukan parsing query parameter dan menyimpannya dalam map[string][]string
*--	Artinya, dalam satu key query parameter, kita bisa memasukkan beberapa value
*--	Caranya kita bisa menambahkan query parameter dengan nama yang sama, namun value berbeda, misal :
*	name=Eko&name=Kurniawan

*/

func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Ahmad&name=Mufied&name=Nugroho", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
