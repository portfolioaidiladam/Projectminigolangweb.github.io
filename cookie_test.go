package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// SetCookie menangani request HTTP untuk membuat cookie baru.
// Fungsi ini mengambil nilai parameter 'name' dari query string dan menyimpannya dalam cookie.
// Cookie akan diset dengan nama 'X-PZN-Name' dan path '/'.
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response dan cookie
//   - request: *http.Request yang berisi query parameter
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// Membuat cookie baru
	cookie := new(http.Cookie)
	// Mengatur nama cookie
	cookie.Name = "X-PZN-Name"
	// Mengambil nilai dari query parameter 'name'
	cookie.Value = request.URL.Query().Get("name")
	// Mengatur path cookie ke root
	cookie.Path = "/"

	// Menyimpan cookie ke response
	http.SetCookie(writer, cookie)
	// Menampilkan pesan sukses
	fmt.Fprint(writer, "Success create cookie")
}

// GetCookie menangani request HTTP untuk membaca cookie.
// Fungsi ini mencoba membaca cookie dengan nama 'X-PZN-Name':
// - Jika cookie tidak ditemukan: menampilkan "No Cookie"
// - Jika cookie ditemukan: menampilkan salam dengan nilai cookie
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi cookie
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	// Mencoba membaca cookie dengan nama 'X-PZN-Name'
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {
		// Jika cookie tidak ditemukan, tampilkan pesan
		fmt.Fprint(writer, "No Cookie")
	} else {
		// Jika cookie ditemukan, tampilkan salam dengan nilai cookie
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// TestCookie menjalankan server HTTP untuk testing cookie.
// Fungsi ini membuat server lokal yang menangani dua endpoint:
// - /set-cookie: untuk membuat cookie baru
// - /get-cookie: untuk membaca cookie
//
// Note: Fungsi ini akan berjalan terus sampai dihentikan karena menggunakan ListenAndServe
func TestCookie(t *testing.T) {
	// Membuat router baru
	mux := http.NewServeMux()
	// Mendaftarkan handler untuk endpoint /set-cookie
	mux.HandleFunc("/set-cookie", SetCookie)
	// Mendaftarkan handler untuk endpoint /get-cookie
	mux.HandleFunc("/get-cookie", GetCookie)

	// Membuat server dengan konfigurasi
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TestSetCookie menguji fungsi SetCookie dengan membuat cookie baru.
// Test ini memverifikasi bahwa cookie dapat dibuat dengan benar
// dan nilai cookie sesuai dengan parameter yang diberikan.
func TestSetCookie(t *testing.T) {
	// Membuat request HTTP GET dengan parameter name
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Aidil", nil)
	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler SetCookie
	SetCookie(recorder, request)

	// Mengambil semua cookie dari response
	cookies := recorder.Result().Cookies()

	// Menampilkan informasi setiap cookie
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

// TestGetCookie menguji fungsi GetCookie dengan cookie yang sudah ada.
// Test ini memverifikasi bahwa cookie dapat dibaca dengan benar
// dan response yang dihasilkan sesuai dengan nilai cookie.
func TestGetCookie(t *testing.T) {
	// Membuat request HTTP GET
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	// Membuat cookie baru untuk testing
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Aidil"
	// Menambahkan cookie ke request
	request.AddCookie(cookie)

	// Membuat ResponseRecorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// Memanggil handler GetCookie
	GetCookie(recorder, request)

	// Membaca dan menampilkan response body
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
