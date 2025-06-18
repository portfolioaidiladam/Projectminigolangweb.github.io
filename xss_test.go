// Package belajar_golang_web berisi implementasi pembelajaran web menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TemplateAutoEscape adalah handler HTTP yang mendemonstrasikan fitur auto-escape template Go
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini mengeksekusi template "post.gohtml" dengan data yang berisi script berbahaya
// untuk mendemonstrasikan bagaimana Go secara otomatis meng-escape konten berbahaya
func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini Adalah Body<script>alert('Anda di Heck')</script></p>",
	})
}

// TestTemplateAutoEscape adalah unit test untuk fungsi TemplateAutoEscape
// Test ini memverifikasi bahwa fitur auto-escape template berfungsi dengan benar
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat request HTTP GET baru
// 2. Membuat recorder untuk menangkap response
// 3. Memanggil TemplateAutoEscape dengan recorder dan request
// 4. Membaca dan mencetak body response untuk verifikasi
func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TestTemplateAutoEscapeServer adalah test yang menjalankan server HTTP untuk TemplateAutoEscape
// Test ini memungkinkan pengujian manual melalui browser
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Server berjalan di localhost:8080 dan akan menangani request untuk TemplateAutoEscape
func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TemplateAutoEscapeDisabled adalah handler HTTP yang mendemonstrasikan cara menonaktifkan auto-escape
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini menggunakan template.HTML untuk menandai konten HTML yang aman
// dan tidak perlu di-escape
func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<h1>Ini Adalah Body</h1>"),
	})
}

// TestTemplateAutoEscapeDisabled adalah unit test untuk fungsi TemplateAutoEscapeDisabled
// Test ini memverifikasi bahwa auto-escape dapat dinonaktifkan dengan benar
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat request HTTP GET baru
// 2. Membuat recorder untuk menangkap response
// 3. Memanggil TemplateAutoEscapeDisabled dengan recorder dan request
// 4. Membaca dan mencetak body response untuk verifikasi
func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TestTemplateAutoEscapeDisabledServer adalah test yang menjalankan server HTTP untuk TemplateAutoEscapeDisabled
// Test ini memungkinkan pengujian manual melalui browser
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Server berjalan di localhost:8080 dan akan menangani request untuk TemplateAutoEscapeDisabled
func TestTemplateAutoEscapeDisabledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TemplateXSS adalah handler HTTP yang mendemonstrasikan potensi kerentanan XSS
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini mengambil konten dari query parameter "body" dan menampilkannya tanpa escape
// PERHATIAN: Handler ini sengaja dibuat untuk demonstrasi dan tidak aman untuk production
func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

// TestTemplateXSS adalah unit test untuk fungsi TemplateXSS
// Test ini mendemonstrasikan bagaimana XSS dapat terjadi jika auto-escape dinonaktifkan
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat request HTTP GET dengan query parameter berisi script
// 2. Membuat recorder untuk menangkap response
// 3. Memanggil TemplateXSS dengan recorder dan request
// 4. Membaca dan mencetak body response untuk verifikasi
func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TestTemplateXSSServer adalah test yang menjalankan server HTTP untuk TemplateXSS
// Test ini memungkinkan pengujian manual melalui browser
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Server berjalan di localhost:8080 dan akan menangani request untuk TemplateXSS
// PERHATIAN: Server ini hanya untuk tujuan pembelajaran dan tidak aman untuk production
func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}