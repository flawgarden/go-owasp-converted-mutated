package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

type BenchmarkTest01325 struct{}

func (b *BenchmarkTest01325) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (b *BenchmarkTest01325) doPost(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{}
	param := r.URL.Query().Get("BenchmarkTest01325")
	bar := b.doSomething(param)

	block, err := des.NewCipher([]byte("examplek")) // 8 bytes key
	if err != nil {
		http.Error(w, "Error initializing crypto", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	mode := NewECBEncrypter(block)
	mode.CryptBlocks(ciphertext, []byte(bar))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	if _, err := file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	response["message"] = "Sensitive value: '" + bar + "' encrypted and stored"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (b *BenchmarkTest01325) doSomething(param string) string {
	return param
}

type ECB struct {
	b         cipher.Block
	blockSize int
}

func NewECBEncrypter(b cipher.Block) *ECB {
	return &ECB{b: b, blockSize: b.BlockSize()}
}

func (e *ECB) BlockSize() int {
	return e.blockSize
}

func (e *ECB) CryptBlocks(dst, src []byte) {
	if len(src)%e.blockSize != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}
	for len(src) > 0 {
		e.b.Encrypt(dst[:e.blockSize], src[:e.blockSize])
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}
