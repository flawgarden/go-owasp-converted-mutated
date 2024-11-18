package controllers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest02103 struct{}

func (b *BenchmarkTest02103) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest02103) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.FormValue("BenchmarkTest02103")
	if param == "" {
		param = ""
	}
	bar := doSomething(param)

	block, err := aes.NewCipher(generateKey())
	if err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCBCEncrypter(block, iv)
	paddedBar := pad([]byte(bar))
	ciphertext := make([]byte, len(paddedBar))
	stream.CryptBlocks(ciphertext, paddedBar)

	fileTarget := "passwordFile.txt"
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Sensitive value: '" + htmlEscape(string(ciphertext)) + "' encrypted and stored<br/>"))
}

func doSomething(param string) string {
	num := 106
	bar := ""
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}

func pad(src []byte) []byte {
	blockSize := aes.BlockSize
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func generateKey() []byte {
	key := "examplekey12345" // 16 bytes for AES-128
	return []byte(key)
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.Handle("/", &BenchmarkTest02103{})
	http.ListenAndServe(":8080", nil)
}
