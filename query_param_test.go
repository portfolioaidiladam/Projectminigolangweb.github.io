package belajar_golang_web

import (
	"fmt"               // Package untuk operasi I/O formatting
	"io"                // Package untuk operasi I/O dasar
	"net/http"          // Package untuk implementasi HTTP client dan server
	"net/http/httptest" // Package untuk testing HTTP
	"strings"           // Package untuk manipulasi string
	"testing"           // Package untuk unit testing
)

// SayHello menangani request HTTP dan mengembalikan salam dengan nama yang diberikan melalui query parameter.
// Jika parameter 'name' tidak ada, akan mengembalikan "Hello".
// Jika parameter 'name' ada, akan mengembalikan "Hello {name}".
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func SayHello(writer http.ResponseWriter, request *http.Request) {
	// Mengambil nilai parameter 'name' dari URL query
	name := request.URL.Query().Get("name")
	
	// Jika name kosong, tampilkan "Hello"
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		// Jika name ada, tampilkan "Hello {name}"
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// TestQueryParameter menguji fungsi SayHello dengan parameter query 'name'.
// Test ini memverifikasi bahwa endpoint dapat menangani parameter query dengan benar.
func TestQueryParameter(t *testing.T) {
	// Membuat request HTTP GET dengan parameter query name=Eko
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Aidil", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler SayHello
	SayHello(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan response body
	fmt.Println(string(body))
}

// MultipleQueryParameter menangani request HTTP dengan dua parameter query: first_name dan last_name.
// Fungsi ini akan mengembalikan salam yang menggabungkan first_name dan last_name.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	// Mengambil nilai parameter first_name dan last_name dari URL query
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	// Menampilkan salam dengan format "Hello {firstName} {lastName}"
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

// TestMultipleQueryParameter menguji fungsi MultipleQueryParameter dengan dua parameter query.
// Test ini memverifikasi bahwa endpoint dapat menangani multiple parameter query dengan benar.
func TestMultipleQueryParameter(t *testing.T) {
	// Membuat request HTTP GET dengan parameter query first_name dan last_name
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Aidil&last_name=Adam", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler MultipleQueryParameter
	MultipleQueryParameter(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan response body
	fmt.Println(string(body))
}

// MultipleParameterValues menangani request HTTP dengan multiple values untuk parameter 'name'.
// Fungsi ini akan menggabungkan semua nilai parameter 'name' yang diberikan dengan spasi sebagai pemisah.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	// Mengambil semua nilai parameter 'name' dari URL query
	query := request.URL.Query()
	names := query["name"]
	// Menggabungkan semua nilai dengan spasi sebagai pemisah
	fmt.Fprint(writer, strings.Join(names, " "))
}

// TestMultipleParameterValues menguji fungsi MultipleParameterValues dengan multiple values untuk parameter 'name'.
// Test ini memverifikasi bahwa endpoint dapat menangani multiple values untuk satu parameter query.
func TestMultipleParameterValues(t *testing.T) {
	// Membuat request HTTP GET dengan multiple values untuk parameter name
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Aidil&name=Adam&name=BaikHati", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler MultipleParameterValues
	MultipleParameterValues(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan response body
	fmt.Println(string(body))
}

