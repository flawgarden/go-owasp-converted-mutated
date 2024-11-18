package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
)

type BenchmarkTest02195 struct{}

func (b *BenchmarkTest02195) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest02195) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02195")
	bar := doSomething(r, param)

	block, err := des.NewCipher([]byte("12345678"))
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	src := []byte(bar)
	src = pad(src, block.BlockSize())
	dst := make([]byte, len(src))
	mode.CryptBlocks(dst, src)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encoded := base64.StdEncoding.EncodeToString(dst)
	if _, err := file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(string(src)))
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	if (500/42)+196 > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}
