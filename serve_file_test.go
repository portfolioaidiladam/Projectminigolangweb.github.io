package belajar_golang_web

import (
	_ "embed"  // Package untuk embed file statis ke dalam binary
	"fmt"      // Package untuk operasi I/O formatting
	"net/http" // Package untuk implementasi HTTP client dan server
	"testing"  // Package untuk unit testing
)

// ServeFile menangani request HTTP untuk melayani file HTML berdasarkan parameter query.
// Fungsi ini memeriksa parameter 'name' dari query string:
// - Jika 'name' ada: melayani file 'ok.html'
// - Jika 'name' kosong: melayani file 'notfound.html'
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi query parameter
func ServeFile(writer http.ResponseWriter, request *http.Request) {
	// Memeriksa keberadaan parameter 'name'
	if request.URL.Query().Get("name") != "" {
		// Jika name ada, tampilkan ok.html
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		// Jika name kosong, tampilkan notfound.html
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

// TestServeFileServer menguji implementasi ServeFile dengan server HTTP.
// Fungsi ini membuat server lokal yang melayani file HTML berdasarkan parameter query.
//
// Note: Fungsi ini akan berjalan terus sampai dihentikan karena menggunakan ListenAndServe
func TestServeFileServer(t *testing.T) {
	// Membuat server dengan handler ServeFile
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
// resourceOk menyimpan konten file ok.html yang di-embed ke dalam binary
var resourceOk string

//go:embed resources/notfound.html
// resourceNotFound menyimpan konten file notfound.html yang di-embed ke dalam binary
var resourceNotFound string

// ServeFileEmbed menangani request HTTP untuk menampilkan konten HTML yang di-embed.
// Fungsi ini memeriksa parameter 'name' dari query string:
// - Jika 'name' ada: menampilkan konten ok.html
// - Jika 'name' kosong: menampilkan konten notfound.html
//
// Parameters:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi query parameter
func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	// Memeriksa keberadaan parameter 'name'
	if request.URL.Query().Get("name") != "" {
		// Jika name ada, tampilkan konten ok.html
		fmt.Fprint(writer, resourceOk)
	} else {
		// Jika name kosong, tampilkan konten notfound.html
		fmt.Fprint(writer, resourceNotFound)
	}
}

// TestServeFileServerEmbed menguji implementasi ServeFileEmbed dengan server HTTP.
// Fungsi ini membuat server lokal yang menampilkan konten HTML yang di-embed berdasarkan parameter query.
//
// Note: Fungsi ini akan berjalan terus sampai dihentikan karena menggunakan ListenAndServe
func TestServeFileServerEmbed(t *testing.T) {
	// Membuat server dengan handler ServeFileEmbed
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}