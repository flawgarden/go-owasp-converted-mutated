package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest00688 struct{}

func (b *BenchmarkTest00688) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00688")
	bar := ""

	if param != "" {
		decoded, _ := base64.StdEncoding.DecodeString(param)
		bar = string(decoded)
	}

	block, err := des.NewCipher([]byte("12345678"))
	if err != nil {
		http.Error(w, "Cipher creation error", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)

	input := []byte{'?'}
	if bar != "" {
		input = []byte(bar)
	}

	paddedInput := pad(input, block.BlockSize())
	output := make([]byte, len(paddedInput))
	mode.CryptBlocks(output, paddedInput)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "File open error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encodedResult := base64.StdEncoding.EncodeToString(output)
	if _, err := file.WriteString("secret_value=" + encodedResult + "\n"); err != nil {
		http.Error(w, "File write error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(string(input)))
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
