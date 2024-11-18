package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/crypto-01/BenchmarkTest01322", BenchmarkTest01322)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest01322(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest01322")
	bar := doSomething(param)

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher([]byte("12345678")) // DES key must be 8 bytes
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}
	mode := cipher.NewCBCEncrypter(block, iv)

	paddedBar := pad([]byte(bar), des.BlockSize)
	ciphertext := make([]byte, len(paddedBar))
	mode.CryptBlocks(ciphertext, paddedBar)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", htmlEncode(param))
}

func doSomething(param string) string {
	return param
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func htmlEncode(s string) string {
	return template.HTMLEscapeString(s)
}
