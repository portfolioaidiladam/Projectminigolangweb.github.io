package belajar_golang_web

import (
	"fmt"               // Package untuk operasi I/O formatting
	"io"                // Package untuk operasi I/O dasar
	"net/http"          // Package untuk implementasi HTTP client dan server
	"net/http/httptest" // Package untuk testing HTTP
	"testing"           // Package untuk unit testing
)

// RequestHeader menangani request HTTP dan mengembalikan nilai dari header 'content-type'.
// Fungsi ini mengambil nilai header 'content-type' dari request dan menampilkannya sebagai response.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request termasuk headers
func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	// Mengambil nilai header 'content-type' dari request
	contentType := request.Header.Get("content-type")
	// Menampilkan nilai content-type sebagai response
	fmt.Fprint(writer, contentType)
}

// TestRequestHeader menguji fungsi RequestHeader dengan menambahkan header 'Content-Type'.
// Test ini memverifikasi bahwa endpoint dapat membaca dan mengembalikan nilai header dengan benar.
func TestRequestHeader(t *testing.T) {
	// Membuat request HTTP POST
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	// Menambahkan header Content-Type dengan nilai application/json
	request.Header.Add("Content-Type", "application/json")

	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler RequestHeader
	RequestHeader(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan response body (nilai content-type)
	fmt.Println(string(body))
}

// ResponseHeader menangani request HTTP dan menambahkan custom header ke response.
// Fungsi ini menambahkan header 'X-Powered-By' ke response dan mengembalikan "OK".
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response dan menambahkan headers
//   - request: *http.Request yang berisi informasi request
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	// Menambahkan custom header 'X-Powered-By' ke response
	writer.Header().Add("X-Powered-By", "Aidil Adam BaikHati")
	// Menampilkan "OK" sebagai response body
	fmt.Fprint(writer, "OK")
}

// TestResponseHeader menguji fungsi ResponseHeader dengan memverifikasi custom header.
// Test ini memastikan bahwa endpoint dapat menambahkan custom header ke response dengan benar.
func TestResponseHeader(t *testing.T) {
	// Membuat request HTTP POST
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	// Menambahkan header Content-Type dengan nilai application/json
	request.Header.Add("Content-Type", "application/json")

	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler ResponseHeader
	ResponseHeader(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan response body
	fmt.Println(string(body))

	// Menampilkan nilai header 'x-powered-by' dari response
	fmt.Println(response.Header.Get("x-powered-by"))
}
