package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TemplateDataMap menangani request HTTP untuk menampilkan template dengan data map.
// Fungsi ini menggunakan map[string]interface{} untuk menyimpan data template.
// Data yang dikirim ke template meliputi:
// - Title: judul halaman
// - Name: nama pengguna
// - Address: map berisi informasi alamat
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	// Membaca dan parse template dari file
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	// Mengeksekusi template dengan data map
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Aidil",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada Lagi",
		},
	})
}

// TestTemplateDataMap menguji fungsi TemplateDataMap dengan data map.
// Test ini memverifikasi bahwa template dapat di-render dengan benar menggunakan data map.
func TestTemplateDataMap(t *testing.T) {
	// Membuat request HTTP GET
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateDataMap
	TemplateDataMap(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Address merepresentasikan struktur data alamat.
// Struct ini digunakan untuk menyimpan informasi alamat pengguna.
type Address struct {
	Street string // Nama jalan
}

// Page merepresentasikan struktur data halaman.
// Struct ini digunakan untuk menyimpan semua data yang diperlukan template.
type Page struct {
	Title   string  // Judul halaman
	Name    string  // Nama pengguna
	Address Address // Informasi alamat pengguna
}

// TemplateDataStruct menangani request HTTP untuk menampilkan template dengan data struct.
// Fungsi ini menggunakan struct Page untuk menyimpan data template.
// Data yang dikirim ke template meliputi:
// - Title: judul halaman
// - Name: nama pengguna
// - Address: struct berisi informasi alamat
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	// Membaca dan parse template dari file
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	// Mengeksekusi template dengan data struct
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Aidil",
		Address: Address{
			Street: "Jalan Belum Ada",
		},
	})
}

// TestTemplateDataStruct menguji fungsi TemplateDataStruct dengan data struct.
// Test ini memverifikasi bahwa template dapat di-render dengan benar menggunakan data struct.
func TestTemplateDataStruct(t *testing.T) {
	// Membuat request HTTP GET
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler TemplateDataStruct
	TemplateDataStruct(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
