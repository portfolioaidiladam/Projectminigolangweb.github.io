package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TemplateActionIf menangani request HTTP untuk menampilkan template dengan action if.
// Fungsi ini menggunakan struct Page untuk menyimpan data template.
// Template akan menampilkan konten berbeda berdasarkan kondisi tertentu.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Aidil",
	})
}

// TestTemplateActionIf menguji fungsi TemplateActionIf dengan action if.
// Test ini memverifikasi bahwa template dapat menangani kondisi if dengan benar.
func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateActionOperator menangani request HTTP untuk menampilkan template dengan operator perbandingan.
// Fungsi ini menggunakan map untuk menyimpan data template.
// Template akan menampilkan konten berbeda berdasarkan hasil perbandingan nilai.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 70,
	})
}

// TestTemplateActionOperator menguji fungsi TemplateActionOperator dengan operator perbandingan.
// Test ini memverifikasi bahwa template dapat menangani operator perbandingan dengan benar.
func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateActionRange menangani request HTTP untuk menampilkan template dengan action range.
// Fungsi ini menggunakan map untuk menyimpan data template.
// Template akan melakukan iterasi pada slice hobbies dan menampilkan setiap elemennya.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
}

// TestTemplateActionRange menguji fungsi TemplateActionRange dengan action range.
// Test ini memverifikasi bahwa template dapat melakukan iterasi dengan benar.
func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateActionWith menangani request HTTP untuk menampilkan template dengan action with.
// Fungsi ini menggunakan map untuk menyimpan data template.
// Template akan mengakses nested data address menggunakan action with.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name": "Aidil",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada",
			"City": "Jakarta",
		},
	})
}

// TestTemplateActionWith menguji fungsi TemplateActionWith dengan action with.
// Test ini memverifikasi bahwa template dapat mengakses nested data dengan benar.
func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
