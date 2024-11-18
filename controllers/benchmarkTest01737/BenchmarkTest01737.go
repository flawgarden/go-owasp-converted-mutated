package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01737 struct{}

func (b *BenchmarkTest01737) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		b.doPost(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01737) doPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := request.URL.Query().Get("BenchmarkTest01737")
	bar := b.doSomething(request, param)

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		http.Error(response, "Error generating key", http.StatusInternalServerError)
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(response, "Error initializing cipher", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(response, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		http.Error(response, "Error generating nonce", http.StatusInternalServerError)
		return
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(bar), nil)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	_, err = response.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", encoded)))
	if err != nil {
		http.Error(response, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func (b *BenchmarkTest01737) doSomething(request *http.Request, param string) string {
	return param
}
