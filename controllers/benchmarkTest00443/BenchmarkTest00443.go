package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type BenchmarkTest00443 struct{}

func (b *BenchmarkTest00443) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00443")

	bar := param
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ciphertext := aead.Seal(nil, iv, []byte(bar), nil)

	output := map[string]string{
		"secret_value": base64.StdEncoding.EncodeToString(ciphertext),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
