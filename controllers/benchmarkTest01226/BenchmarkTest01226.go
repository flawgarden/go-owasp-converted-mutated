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

type BenchmarkTest01226 struct{}

func (b *BenchmarkTest01226) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.doPost(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func (b *BenchmarkTest01226) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.FormValue("BenchmarkTest01226")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	key := make([]byte, 32) // 32 bytes for AES-256
	if _, err := rand.Read(key); err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, "Error initializing cipher", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		http.Error(w, "Error initializing GCM", http.StatusInternalServerError)
		return
	}

	encrypted := gcm.Seal(nil, iv, []byte(bar), nil)

	fileTarget := "passwordFile.txt"
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("secret_value=%s\n", base64.StdEncoding.EncodeToString(encrypted))); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(bar)))); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/crypto-01/BenchmarkTest01226", &BenchmarkTest01226{})
	http.ListenAndServe(":8080", nil)
}
