package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest struct{}

func (bt *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This endpoint requires a POST request.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.FormValue("BenchmarkTest00050")

	block, err := des.NewCipher([]byte("12345678")) // DES needs 8-byte key
	if err != nil {
		http.Error(w, "Error with cipher initialization.", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Error generating initialization vector.", http.StatusInternalServerError)
		return
	}

	c := cipher.NewCBCEncrypter(block, iv)
	input := []byte(param)
	plaintext := pad(input, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	c.CryptBlocks(ciphertext, plaintext)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file to write.", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	if _, err := file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(w, "Error writing to file.", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": fmt.Sprintf("Sensitive value: '%s' encrypted and stored", param)}
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00050", &BenchmarkTest{})
	http.ListenAndServe(":8080", nil)
}
