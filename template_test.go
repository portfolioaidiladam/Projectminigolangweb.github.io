package belajar_golang_web

import (
	"embed"
	"fmt"               // Package untuk operasi I/O formatting
	"html/template"     // Package untuk parsing dan eksekusi template HTML
	"io"                // Package untuk operasi I/O dasar
	"net/http"          // Package untuk implementasi HTTP client dan server
	"net/http/httptest" // Package untuk testing HTTP
	"testing"           // Package untuk unit testing
)

//go:embed templates
var templates embed.FS

// SimpleHTML menangani request HTTP untuk menampilkan template HTML sederhana.
// Fungsi ini membuat template HTML inline dan mengeksekusinya dengan data "Hello HTML Template".
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	// Mendefinisikan template HTML sederhana dengan placeholder {{.}}
	templateText := `<html><body>{{.}}</body></html>`
	
	// Membuat dan parse template baru dengan nama "SIMPLE"
	// Menggunakan template.Must untuk menangani error parsing
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	// Mengeksekusi template dengan data "Hello HTML Template"
	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

// TestSimpleHTML menguji fungsi SimpleHTML dengan template HTML sederhana.
// Test ini memverifikasi bahwa template dapat di-render dengan benar.
func TestSimpleHTML(t *testing.T) {
	// Membuat request HTTP GET
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler SimpleHTML
	SimpleHTML(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// SimpleHTMLFile menangani request HTTP untuk menampilkan template HTML dari file.
// Fungsi ini membaca template dari file 'simple.gohtml' dan mengeksekusinya.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	// Membaca dan parse template dari file
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	// Mengeksekusi template dengan data "Hello HTML Template"
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

// TestSimpleHTMLFile menguji fungsi SimpleHTMLFile dengan template dari file.
// Test ini memverifikasi bahwa template dari file dapat di-render dengan benar.
func TestSimpleHTMLFile(t *testing.T) {
	// Membuat request HTTP GET
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler SimpleHTMLFile
	SimpleHTMLFile(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateDirectory menangani request HTTP untuk menampilkan template HTML dari direktori.
// Fungsi ini membaca semua file template dengan ekstensi .gohtml dari direktori templates.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	// Membaca dan parse semua file template dari direktori
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	// Mengeksekusi template simple.gohtml dengan data "Hello HTML Template"
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

// TestTemplateDirectory menguji fungsi TemplateDirectory dengan template dari direktori.
// Test ini memverifikasi bahwa template dari direktori dapat di-render dengan benar.
func TestTemplateDirectory(t *testing.T) {
	// Membuat request HTTP GET
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateDirectory
	TemplateDirectory(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateEmbed menangani request HTTP untuk menampilkan template HTML yang di-embed.
// Fungsi ini membaca template dari filesystem yang di-embed ke dalam binary.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	// Membaca dan parse template dari embedded filesystem
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	// Mengeksekusi template simple.gohtml dengan data "Hello HTML Template"
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

// TestTemplateEmbed menguji fungsi TemplateEmbed dengan template yang di-embed.
// Test ini memverifikasi bahwa template yang di-embed dapat di-render dengan benar.
func TestTemplateEmbed(t *testing.T) {
	// Membuat request HTTP GET
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateEmbed
	TemplateEmbed(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}