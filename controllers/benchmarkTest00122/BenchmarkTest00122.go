package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest00122 struct{}

func (b *BenchmarkTest00122) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00122")
	param, _ = url.QueryUnescape(param)

	iv := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	key := make([]byte, 32) // AES-256 key
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	ciphertext := aesGCM.Seal(nil, iv, []byte(param), nil)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Sensitive value encrypted and stored"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00122", &BenchmarkTest00122{})
	http.ListenAndServe(":8080", nil)
}
