// Package belajar_golang_web berisi implementasi pembelajaran web menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// LogMiddleware adalah middleware untuk logging request
// Middleware ini akan mencetak log sebelum dan sesudah handler dieksekusi
type LogMiddleware struct {
	Handler http.Handler
}

// ServeHTTP mengimplementasikan interface http.Handler
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Middleware ini akan:
// 1. Mencetak log "Before Execute Handler"
// 2. Menjalankan handler yang dibungkus
// 3. Mencetak log "After Execute Handler"
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

// ErrorHandler adalah middleware untuk menangani panic
// Middleware ini akan menangkap panic yang terjadi di handler dan mengembalikan response error
type ErrorHandler struct {
	Handler http.Handler
}

// ServeHTTP mengimplementasikan interface http.Handler
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Middleware ini akan:
// 1. Menggunakan defer untuk menangkap panic
// 2. Jika terjadi panic:
//    - Mencetak log "Terjadi Error"
//    - Mengembalikan status 500 Internal Server Error
//    - Menulis pesan error ke response
// 3. Jika tidak ada panic, menjalankan handler normal
func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

// TestMiddleware adalah test yang menjalankan server HTTP dengan middleware
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat mux baru dengan beberapa endpoint:
//    - /: menampilkan "Hello Middleware"
//    - /foo: menampilkan "Hello Foo"
//    - /panic: sengaja memicu panic untuk testing error handler
// 2. Membuat instance LogMiddleware yang membungkus mux
// 3. Membuat instance ErrorHandler yang membungkus LogMiddleware
// 4. Membuat dan menjalankan server di localhost:8080
//
// Server ini mendemonstrasikan penggunaan middleware untuk:
// - Logging request
// - Penanganan error/panic
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		panic("Ups")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
