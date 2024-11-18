package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

type BenchmarkTest00687 struct{}

func (bt *BenchmarkTest00687) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00687")
	if param == "" {
		param = ""
	}

	num := 106
	bar := ""
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	block, err := aes.NewCipher(generateKey())
	if err != nil {
		http.Error(w, "Encryption error", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Encryption error", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		http.Error(w, "Encryption error", http.StatusInternalServerError)
		return
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(bar), nil)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	if _, err = file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"sensitive_value": encoded,
	}
	json.NewEncoder(w).Encode(response)
}

func generateKey() []byte {
	return []byte("a very secret key") // Replace with a secure key generation
}
