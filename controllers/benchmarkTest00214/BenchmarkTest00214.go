package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BenchmarkTest00214 struct{}

func (bt *BenchmarkTest00214) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
		return
	}
	bt.doPost(w, r)
}

func (bt *BenchmarkTest00214) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for name := range r.Header {
		if name != "User-Agent" && name != "Accept" { // exclude standard headers
			param = name
			break
		}
	}

	bar := param

	key := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		http.Error(w, "Can't generate key", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, "Can't create cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, aes.BlockSize+len([]byte(bar)))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Can't generate IV", http.StatusInternalServerError)
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(bar))

	result := base64.StdEncoding.EncodeToString(ciphertext)
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		http.Error(w, "Can't open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + result + "\n"); err != nil {
		http.Error(w, "Can't write to file", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Sensitive value: '" + htmlEscape(bar) + "' encrypted and stored",
	}
	json.NewEncoder(w).Encode(response)
}

func htmlEscape(text string) string {
	return fmt.Sprintf("%q", text)
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00214", &BenchmarkTest00214{})
	http.ListenAndServe(":8080", nil)
}
