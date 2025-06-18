// Package belajar_golang_web berisi implementasi pembelajaran web menggunakan Go
package belajar_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// UploadForm adalah handler HTTP yang menampilkan form upload file
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini mengeksekusi template "upload.form.gohtml" untuk menampilkan form upload
func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

// Upload adalah handler HTTP yang memproses upload file
// Parameter:
//   - writer: http.ResponseWriter untuk menulis response
//   - request: *http.Request yang berisi informasi request
//
// Handler ini:
// 1. Mengambil file dari form dengan nama field "file"
// 2. Membuat file baru di direktori ./resources/
// 3. Menyalin konten file yang diupload ke file tujuan
// 4. Mengambil nilai field "name" dari form
// 5. Menampilkan template sukses dengan data nama dan path file
func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(32 << 20)
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

// TestUploadForm adalah test yang menjalankan server HTTP untuk form upload
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat mux baru untuk menangani routing
// 2. Mendaftarkan handler:
//    - /: untuk menampilkan form upload
//    - /upload: untuk memproses upload file
//    - /static/: untuk menangani akses file yang sudah diupload
// 3. Membuat dan menjalankan server di localhost:8080
func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// uploadFileTest adalah variabel yang menyimpan konten file test yang akan diupload
// File ini di-embed menggunakan direktif //go:embed
//go:embed resources/PNZ-ICON.png
var uploadFileTest []byte

// TestUploadFile adalah unit test untuk fungsi Upload
// Parameter:
//   - t: *testing.T untuk menangani testing
//
// Test ini:
// 1. Membuat buffer untuk menyimpan body request
// 2. Membuat writer multipart untuk form data
// 3. Menambahkan field "name" ke form
// 4. Menambahkan file test ke form
// 5. Membuat request POST ke /upload
// 6. Mengatur Content-Type header
// 7. Memanggil handler Upload
// 8. Membaca dan mencetak response
func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Aidil Adam BaikHati")
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

