// Package belajar_golang_web berisi implementasi dasar web server menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// TestHandler mendemonstrasikan implementasi handler HTTP paling sederhana
// Fungsi ini membuat server HTTP yang merespon dengan "Hello World" untuk setiap request
func TestHandler(t *testing.T) {
	// Mendefinisikan handler function yang akan menangani request
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	// Konfigurasi server HTTP
	server := http.Server{
		Addr:    "localhost:8080", // Menentukan alamat server
		Handler: handler,          // Menentukan handler yang akan digunakan
	}

	// Menjalankan server dan menangani error jika terjadi
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TestServeMux mendemonstrasikan penggunaan ServeMux untuk routing
// ServeMux memungkinkan penanganan multiple endpoint dengan pattern matching
func TestServeMux(t *testing.T) {
	// Membuat instance baru dari ServeMux
	mux := http.NewServeMux()

	// Mendaftarkan handler untuk root path "/"
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	// Mendaftarkan handler untuk path "/hi"
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	// Mendaftarkan handler untuk path yang dimulai dengan "/images/"
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Image")
	})

	// Mendaftarkan handler untuk path yang dimulai dengan "/images/thumbnails/"
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Thumbnail")
	})

	// Konfigurasi dan menjalankan server dengan ServeMux
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TestRequest mendemonstrasikan cara mengakses informasi request
// Fungsi ini menampilkan method HTTP dan URI dari setiap request yang masuk
func TestRequest(t *testing.T) {
	// Mendefinisikan handler yang akan menampilkan informasi request
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)    // Menampilkan HTTP method (GET, POST, dll)
		fmt.Fprintln(writer, request.RequestURI) // Menampilkan URI yang diminta
	}

	// Konfigurasi dan menjalankan server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
