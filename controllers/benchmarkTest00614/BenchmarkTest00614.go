package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00614 struct{}

func (b *BenchmarkTest00614) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00614")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	// 8-byte initialization vector
	iv := make([]byte, des.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		http.Error(w, "Failed to generate IV", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher([]byte("abcdefgh")) // Example key
	if err != nil {
		http.Error(w, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, []byte(bar))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encoded := base64.StdEncoding.EncodeToString(append(iv, ciphertext...))
	if _, err := file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", param)
}
