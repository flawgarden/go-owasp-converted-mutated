package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/crypto/chacha20poly1305"
)

type BenchmarkTest00255 struct {
}

func (bt *BenchmarkTest00255) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
		return
	}
	bt.doPost(w, r)
}

func (bt *BenchmarkTest00255) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest00255"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}

	key := make([]byte, chacha20poly1305.KeySize)
	random := rand.Reader
	if _, err := io.ReadFull(random, key); err != nil {
		http.Error(w, "Error generating random key", http.StatusInternalServerError)
		return
	}

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(random, nonce); err != nil {
		http.Error(w, "Error generating random nonce", http.StatusInternalServerError)
		return
	}

	input := []byte(bar)
	ciphertext := aead.Seal(nil, nonce, input, nil)

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

	if _, err := w.Write([]byte("Sensitive value: '" + html.EscapeString(string(input)) + "' encrypted and stored<br/>")); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
