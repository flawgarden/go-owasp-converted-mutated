package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest00686 struct{}

func (b *BenchmarkTest00686) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00686")
	bar := ""
	if param != "" {
		barBytes, err := base64.StdEncoding.DecodeString(param)
		if err == nil {
			bar = string(barBytes)
		}
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	key := make([]byte, 32) // AES-256 key
	if _, err := rand.Read(key); err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	input := []byte{42} // Input byte
	if bar != "" {
		input = []byte(bar)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		http.Error(w, "Error generating nonce", http.StatusInternalServerError)
		return
	}

	ciphertext := aesgcm.Seal(nonce, nonce, input, nil)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(string(input)))
	w.Write([]byte(response))
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00686", &BenchmarkTest00686{})
	http.ListenAndServe(":8080", nil)
}
