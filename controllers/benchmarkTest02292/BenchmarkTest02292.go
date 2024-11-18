package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

type BenchmarkTest02292 struct{}

func (b *BenchmarkTest02292) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest02292) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	names := r.URL.Query()
	for name, values := range names {
		for _, value := range values {
			if value == "BenchmarkTest02292" {
				param = name
				break
			}
		}
	}

	bar := doSomething(param)

	block, err := des.NewCipher([]byte("12345678"))
	if err != nil {
		http.Error(w, "Error initializing cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, []byte(bar))

	file, _ := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	file.WriteString("secret_value=" + encoded + "\n")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sensitive value: '" + bar + "' encrypted and stored",
	})
}

func doSomething(param string) string {
	return param // здесь предполагается некоторая обработка параметра
}

func main() {
	http.ListenAndServe(":8080", &BenchmarkTest02292{})
}
