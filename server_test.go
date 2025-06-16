// Package belajar_golang_web berisi implementasi dasar web server menggunakan Go
package belajar_golang_web

import (
	"net/http"
	"testing"
)

// TestServer mendemonstrasikan cara membuat dan menjalankan server HTTP dasar
// Fungsi ini membuat server yang berjalan di localhost:8080 tanpa handler khusus
func TestServer(t *testing.T) {
	// Membuat instance server HTTP dengan konfigurasi dasar
	server := http.Server{
		Addr: "localhost:8080", // Menentukan alamat server
	}

	// Menjalankan server dan menangani error jika terjadi
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
