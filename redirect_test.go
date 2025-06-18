// Package belajar_golang_web berisi implementasi pembelajaran web menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// RedirectTo adalah handler HTTP yang menampilkan pesan "Hello Redirect"
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini digunakan sebagai endpoint tujuan redirect
func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}

// RedirectFrom adalah handler HTTP yang melakukan redirect ke endpoint "/redirect-to"
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini mendemonstrasikan cara melakukan redirect internal menggunakan http.StatusTemporaryRedirect (307)
func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

// RedirectOut adalah handler HTTP yang melakukan redirect ke URL eksternal
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini mendemonstrasikan cara melakukan redirect ke website eksternal
// menggunakan http.StatusTemporaryRedirect (307)
func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "https://www.aidiladam.com/", http.StatusTemporaryRedirect)
}

// TestRedirect adalah test yang menjalankan server HTTP dengan beberapa endpoint redirect
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat mux baru untuk menangani routing
// 2. Mendaftarkan tiga handler:
//    - /redirect-from: untuk redirect internal
//    - /redirect-to: endpoint tujuan redirect internal
//    - /redirect-out: untuk redirect ke website eksternal
// 3. Membuat dan menjalankan server di localhost:8080
//
// Server ini dapat diakses melalui browser untuk menguji fungsionalitas redirect
func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
