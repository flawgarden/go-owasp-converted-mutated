package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"text/template"
)

type BenchmarkTest00609 struct{}

func (bt BenchmarkTest00609) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00609")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	block, err := des.NewCipher(generateKey())
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Error creating IV", http.StatusInternalServerError)
		return
	}

	ciphertext := encrypt(block, iv, []byte(bar))
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	writeToFile(encoded)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Sensitive value: '" + htmlEscape(bar) + "' encrypted and stored<br/>"))
}

func generateKey() []byte {
	key := make([]byte, 8) // DES key size
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil
	}
	return key
}

func encrypt(block cipher.Block, iv []byte, plaintext []byte) []byte {
	mode := cipher.NewCBCEncrypter(block, iv)
	paddedText := pad(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(paddedText))
	mode.CryptBlocks(ciphertext, paddedText)
	return ciphertext
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

func doSomething(param string) string {
	// Mock implementation of the external logic
	return "processed_" + param
}

func writeToFile(value string) {
	// Implement write to the file logic
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00609", BenchmarkTest00609{})
	http.ListenAndServe(":8080", nil)
}
