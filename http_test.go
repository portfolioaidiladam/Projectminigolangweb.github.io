// Package belajar_golang_web berisi implementasi dasar web server menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HelloHandler adalah handler HTTP sederhana yang merespon dengan "Hello World"
// Handler ini digunakan sebagai contoh implementasi dasar HTTP handler
func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

// TestHttp mendemonstrasikan cara melakukan testing terhadap HTTP handler
// menggunakan httptest package dari Go
func TestHttp(t *testing.T) {
	// Membuat request HTTP GET ke endpoint /hello
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	
	// Membuat ResponseRecorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()

	// Memanggil handler dengan request dan recorder
	HelloHandler(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	
	// Membaca body response
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	// Menampilkan response body
	fmt.Println(bodyString)
}

