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

type BenchmarkTest00612 struct{}

func (b *BenchmarkTest00612) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00612")
	if param == "" {
		param = ""
	}

	var bar string
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	// Generate a random IV
	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Failed to generate IV", http.StatusInternalServerError)
		return
	}

	key := make([]byte, 32) // 256-bit key
	if _, err := rand.Read(key); err != nil {
		http.Error(w, "Failed to generate key", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Failed to create GCM", http.StatusInternalServerError)
		return
	}

	ciphertext := aesGcm.Seal(nil, iv, []byte(bar), nil)

	fileTarget := "passwordFile.txt"
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Sensitive value encrypted and stored"}
	json.NewEncoder(w).Encode(response)
}
