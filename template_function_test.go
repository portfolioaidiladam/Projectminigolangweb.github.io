// Package belajar_golang_web berisi implementasi web sederhana menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MyPage merepresentasikan struktur data untuk halaman
type MyPage struct {
	Name string // Nama pengguna
}

// SayHello adalah method yang mengembalikan string salam
// Parameter:
//   - name: Nama yang akan disapa
// Return: String salam yang menggabungkan nama parameter dan nama dari struct
func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

// TemplateFunction menangani request HTTP untuk menampilkan template dengan fungsi method
// Parameter:
//   - writer: ResponseWriter untuk menulis response
//   - request: Request HTTP yang diterima
func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	// Membuat template baru dengan nama "FUNCTION" dan mengisi kontennya
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	
	// Eksekusi template dengan data MyPage
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Aidil",
	})
}

// TestTemplateFunction melakukan testing untuk fungsi TemplateFunction
func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateFunctionGlobal menangani request HTTP untuk menampilkan template dengan fungsi global
// Menggunakan fungsi len() bawaan Go template
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Aidil",
	})
}

// TestTemplateFunctionGlobal melakukan testing untuk fungsi TemplateFunctionGlobal
func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateFunctionCreateGlobal menangani request HTTP untuk menampilkan template dengan fungsi kustom
// Membuat fungsi global baru bernama "upper" untuk mengubah string menjadi huruf besar
func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	// Mendaftarkan fungsi kustom "upper"
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Aidil Adam BaikHati",
	})
}

// TestTemplateFunctionCreateGlobal melakukan testing untuk fungsi TemplateFunctionCreateGlobal
func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TemplateFunctionCreateGlobalPipeline menangani request HTTP untuk menampilkan template dengan pipeline fungsi
// Menggunakan multiple fungsi dalam satu pipeline: sayHello -> upper
func TemplateFunctionCreateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	// Mendaftarkan dua fungsi kustom: sayHello dan upper
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	// Menggunakan pipeline operator (|) untuk menggabungkan fungsi
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Aidil Adam BaikHati",
	})
}

// TestTemplateFunctionCreateGlobalPipeline melakukan testing untuk fungsi TemplateFunctionCreateGlobalPipeline
func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
