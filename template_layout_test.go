// Package belajar_golang_web berisi implementasi web sederhana menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TemplateLayout menangani request HTTP untuk menampilkan template dengan layout
// Parameter:
//   - writer: ResponseWriter untuk menulis response
//   - request: Request HTTP yang diterima
func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	// Parse semua file template yang dibutuhkan
	// Must akan panic jika terjadi error saat parsing
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",  // Template untuk header
		"./templates/footer.gohtml",  // Template untuk footer
		"./templates/layout.gohtml",  // Template utama yang menggunakan header dan footer
	))

	// Eksekusi template dengan data yang diberikan
	// Data yang dikirim:
	//   - Title: Judul halaman
	//   - Name: Nama pengguna
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name" : "Aidil",
	})
}

// TestTemplateLayout melakukan testing untuk fungsi TemplateLayout
// Test ini akan:
// 1. Membuat request HTTP GET
// 2. Mencatat response
// 3. Menampilkan body response
func TestTemplateLayout(t *testing.T) {
	// Membuat request HTTP GET ke localhost:8080
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Membuat recorder untuk mencatat response
	recorder := httptest.NewRecorder()

	// Memanggil fungsi yang akan di-test
	TemplateLayout(recorder, request)

	// Membaca dan menampilkan body response
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
