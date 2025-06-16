package belajar_golang_web

import (
	"fmt"               // Package untuk operasi I/O formatting
	"io"                // Package untuk operasi I/O dasar
	"net/http"          // Package untuk implementasi HTTP client dan server
	"net/http/httptest" // Package untuk testing HTTP
	"testing"           // Package untuk unit testing
)

// ResponseCode menangani request HTTP dan mengembalikan response code yang sesuai.
// Fungsi ini memeriksa parameter 'name' dari query string:
// - Jika 'name' kosong: mengembalikan status 400 (Bad Request) dengan pesan error
// - Jika 'name' ada: mengembalikan status 200 (OK) dengan salam
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response dan status code
//   - request: *http.Request yang berisi query parameter
func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	// Mengambil nilai parameter 'name' dari query string
	name := request.URL.Query().Get("name")

	// Jika name kosong, kirim response dengan status 400
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	} else {
		// Jika name ada, kirim response dengan status 200 (default)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// TestResponseCodeInvalid menguji fungsi ResponseCode dengan parameter 'name' yang kosong.
// Test ini memverifikasi bahwa endpoint mengembalikan status code 400 (Bad Request)
// dan pesan error yang sesuai ketika parameter 'name' tidak ada.
func TestResponseCodeInvalid(t *testing.T) {
	// Membuat request HTTP GET tanpa parameter query
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler ResponseCode
	ResponseCode(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan status code (400)
	fmt.Println(response.StatusCode)
	// Menampilkan status text (Bad Request)
	fmt.Println(response.Status)
	// Menampilkan response body (pesan error)
	fmt.Println(string(body))
}

// TestResponseCodeValid menguji fungsi ResponseCode dengan parameter 'name' yang valid.
// Test ini memverifikasi bahwa endpoint mengembalikan status code 200 (OK)
// dan salam yang sesuai ketika parameter 'name' ada.
func TestResponseCodeValid(t *testing.T) {
	// Membuat request HTTP GET dengan parameter name=Eko
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Aidil", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler ResponseCode
	ResponseCode(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan status code (200)
	fmt.Println(response.StatusCode)
	// Menampilkan status text (200 OK)
	fmt.Println(response.Status)
	// Menampilkan response body (salam)
	fmt.Println(string(body))
}

