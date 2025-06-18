// Package belajar_golang_web berisi implementasi pembelajaran web menggunakan Go
package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// templateFiles adalah variabel yang menggunakan embed.FS untuk menyimpan template HTML
// Direktif //go:embed digunakan untuk menyertakan semua file .gohtml dari folder templates
//go:embed templates/*.gohtml
var templateFiles embed.FS

// myTemplates adalah variabel global yang menyimpan template yang sudah di-parse
// template.Must digunakan untuk memastikan tidak ada error saat parsing template
// ParseFS digunakan untuk membaca template dari filesystem yang di-embed
var myTemplates = template.Must(template.ParseFS(templateFiles, "templates/*.gohtml"))

// TemplateCaching adalah handler HTTP yang mendemonstrasikan penggunaan template caching
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini mengeksekusi template "simple.gohtml" dengan data "Hello Template Caching"
func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
}

// TestTemplateCaching adalah unit test untuk fungsi TemplateCaching
// Test ini memverifikasi bahwa template caching berfungsi dengan benar
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat request HTTP GET baru
// 2. Membuat recorder untuk menangkap response
// 3. Memanggil TemplateCaching dengan recorder dan request
// 4. Membaca dan mencetak body response
func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}