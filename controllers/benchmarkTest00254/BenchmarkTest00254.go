package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

type BenchmarkTest00254 struct{}

func (b *BenchmarkTest00254) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (b *BenchmarkTest00254) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00254")
	decodedParam, _ := base64.StdEncoding.DecodeString(param)

	block, _ := des.NewCipher(generateKey())
	iv := make([]byte, block.BlockSize())
	rand.Read(iv)

	ciphertext := encrypt(block, iv, decodedParam)
	encodedResult := base64.StdEncoding.EncodeToString(ciphertext)

	w.Write([]byte("Sensitive value: '" + encodedResult + "' encrypted and stored<br/>"))
}

func generateKey() []byte {
	key := make([]byte, 8)
	rand.Read(key)
	return key
}

func encrypt(block cipher.Block, iv, plaintext []byte) []byte {
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00254", &BenchmarkTest00254{})
	http.ListenAndServe(":8080", nil)
}
