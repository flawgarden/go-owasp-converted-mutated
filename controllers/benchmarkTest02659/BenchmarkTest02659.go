package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02659 struct{}

func (b *BenchmarkTest02659) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest02659) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.FormValue("BenchmarkTest02659")
	bar := doSomething(param)

	block, err := aes.NewCipher(generateKey())
	if err != nil {
		http.Error(w, "Error initializing cipher", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		http.Error(w, "Error generating nonce", http.StatusInternalServerError)
		return
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(bar), nil)

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

	output := fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(string(ciphertext)))
	w.Write([]byte(output))
}

func generateKey() []byte {
	key := make([]byte, 32) // AES-256
	if _, err := rand.Read(key); err != nil {
		return nil
	}
	return key
}

func doSomething(param string) string {
	return param
}

func htmlEscape(s string) string {
	return s // Implement HTML escaping if necessary
}
