package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// TestFileServer menguji implementasi file server menggunakan http.Dir.
// Fungsi ini membuat server lokal yang melayani file statis dari direktori './resources'.
// File dapat diakses melalui URL dengan prefix '/static/'.
//
// Note: Fungsi ini akan berjalan terus sampai dihentikan karena menggunakan ListenAndServe
func TestFileServer(t *testing.T) {
	// Membuat file server dari direktori './resources'
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	// Membuat router baru
	mux := http.NewServeMux()
	// Mendaftarkan handler untuk path '/static/' dengan menghapus prefix '/static'
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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

//go:embed resources
// resources adalah variabel yang menyimpan file statis yang di-embed ke dalam binary.
// Direktori 'resources' akan di-embed ke dalam binary saat kompilasi.
var resources embed.FS

// TestFileServerGolangEmbed menguji implementasi file server menggunakan Go embed.
// Fungsi ini membuat server lokal yang melayani file statis yang di-embed ke dalam binary.
// File dapat diakses melalui URL dengan prefix '/static/'.
//
// Note: Fungsi ini akan berjalan terus sampai dihentikan karena menggunakan ListenAndServe
func TestFileServerGolangEmbed(t *testing.T) {
	// Mengambil sub-filesystem 'resources' dari embedded filesystem
	directory, _ := fs.Sub(resources, "resources")
	// Membuat file server dari embedded filesystem
	fileServer := http.FileServer(http.FS(directory))

	// Membuat router baru
	mux := http.NewServeMux()
	// Mendaftarkan handler untuk path '/static/' dengan menghapus prefix '/static'
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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
