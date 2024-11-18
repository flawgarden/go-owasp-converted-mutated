package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

type BenchmarkTest00019 struct {
	http.Handler
}

func (b *BenchmarkTest00019) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "This input source requires a POST, not a GET.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var input []byte
	if r.Method == http.MethodPost {
		input = make([]byte, 1000)
		n, err := r.Body.Read(input)
		if err != nil && err.Error() != "EOF" {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		input = input[:n]
	} else {
		input = []byte("?")
	}

	block, err := des.NewCipher([]byte("12345678")) // пример ключа
	if err != nil {
		http.Error(w, "Problem initializing cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(input))
	mode := NewECBEncrypter(block)
	mode.CryptBlocks(ciphertext, input)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n")

	output, _ := json.Marshal("Sensitive value: '" + string(input) + "' encrypted and stored")
	w.Write(output)
}

type ecb struct {
	b       cipher.Block
	block   []byte
	blocked int
}

func NewECBEncrypter(b cipher.Block) *ecb {
	return &ecb{
		b: b,
	}
}

func (e *ecb) BlockSize() int {
	return e.b.BlockSize()
}

func (e *ecb) CryptBlocks(dst, src []byte) {
	for len(src) > 0 {
		e.b.Encrypt(dst[:e.b.BlockSize()], src[:e.b.BlockSize()])
		src = src[e.b.BlockSize():]
		dst = dst[e.b.BlockSize():]
	}
}
