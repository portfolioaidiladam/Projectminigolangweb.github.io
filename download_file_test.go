// Package belajar_golang_web berisi implementasi pembelajaran web menggunakan Go
package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// DownloadFile adalah handler HTTP yang menangani proses download file
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini:
// 1. Mengambil nama file dari query parameter "file"
// 2. Memvalidasi apakah parameter file ada
// 3. Mengatur header Content-Disposition untuk memaksa browser mendownload file
// 4. Mengirim file dari direktori ./resources/ ke client
//
// Jika parameter file kosong, handler akan mengembalikan status 400 Bad Request
func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, request, "./resources/"+file)
}

// TestDownloadFile adalah test yang menjalankan server HTTP untuk download file
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat server HTTP baru
// 2. Mengatur handler DownloadFile untuk semua request
// 3. Menjalankan server di localhost:8080
//
// Server dapat diakses melalui browser dengan URL:
// http://localhost:8080?file=nama_file.png
func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
