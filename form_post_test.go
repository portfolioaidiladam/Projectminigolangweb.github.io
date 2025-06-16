package belajar_golang_web

import (
	"fmt"               // Package untuk operasi I/O formatting
	"io"                // Package untuk operasi I/O dasar
	"net/http"          // Package untuk implementasi HTTP client dan server
	"net/http/httptest" // Package untuk testing HTTP
	"strings"           // Package untuk manipulasi string
	"testing"           // Package untuk unit testing
)

// FormPost menangani request HTTP POST dengan form data.
// Fungsi ini memproses form data yang dikirim melalui POST request dan mengembalikan
// salam yang menggabungkan first_name dan last_name dari form data.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi form data
//
// Panic:
//   - Jika terjadi error saat parsing form data
func FormPost(writer http.ResponseWriter, request *http.Request) {
	// Parse form data dari request body
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	// Mengambil nilai first_name dan last_name dari form data
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// Menampilkan salam dengan format "Hello {firstName} {lastName}"
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

// TestFormPost menguji fungsi FormPost dengan mengirim form data melalui POST request.
// Test ini memverifikasi bahwa endpoint dapat memproses form data dengan benar.
func TestFormPost(t *testing.T) {
	// Membuat request body dengan format form-urlencoded
	requestBody := strings.NewReader("first_name=Aidil&last_name=Adam")
	// Membuat request HTTP POST dengan body yang sudah dibuat
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	// Menambahkan header Content-Type untuk form-urlencoded
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler FormPost
	FormPost(recorder, request)

	// Mengambil response dari recorder
	response := recorder.Result()
	// Membaca body response
	body, _ := io.ReadAll(response.Body)

	// Menampilkan response body
	fmt.Println(string(body))
}
