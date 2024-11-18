package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
)

type BenchmarkTest01320Controller struct {
	http.Handler
}

func (c *BenchmarkTest01320Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else if r.Method == http.MethodPost {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest01320Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01320")
	bar := doSomething(param)

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Failed to generate IV", http.StatusInternalServerError)
		return
	}

	key := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		http.Error(w, "Failed to generate key", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher(key)
	if err != nil {
		http.Error(w, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar)+des.BlockSize)
	copy(ciphertext[des.BlockSize:], []byte(bar))
	mode := NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], []byte(bar))

	fileTarget := "passwordFile.txt"
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(bar))
}

func doSomething(param string) string {
	if param == "" {
		return ""
	}
	decoded, _ := base64.StdEncoding.DecodeString(param)
	return string(decoded)
}

type CBCEncrypter struct {
	b         cipher.Block
	blockSize int
	iv        []byte
}

func NewCBCEncrypter(b cipher.Block, iv []byte) *CBCEncrypter {
	return &CBCEncrypter{b, b.BlockSize(), iv}
}

func (x *CBCEncrypter) BlockSize() int { return x.blockSize }

func (x *CBCEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("dst short")
	}

	mode := x.iv
	for len(src) > 0 {
		for i := 0; i < x.blockSize; i++ {
			dst[i] = src[i] ^ mode[i]
		}
		x.b.Encrypt(dst[:x.blockSize], dst[:x.blockSize])
		mode = dst[:x.blockSize]
		dst = dst[x.blockSize:]
		src = src[x.blockSize:]
	}
}
