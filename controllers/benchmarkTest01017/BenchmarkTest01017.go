package controllers

import (
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BenchmarkTest01017 struct{}

func (b *BenchmarkTest01017) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01017")
	if param != "" {
		var err error
		param, err = url.QueryUnescape(param)
		if err != nil {
			http.Error(w, "Failed to decode parameter", http.StatusBadRequest)
			return
		}
	}

	bar := new(Test).doSomething(r, param)

	cipher, err := des.NewCipher(generateKey())
	if err != nil {
		http.Error(w, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Failed to create IV", http.StatusInternalServerError)
		return
	}

	output := make([]byte, len(bar))
	cipher.Encrypt(output, []byte(bar))

	encoded := base64.StdEncoding.EncodeToString(output)
	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", encoded)
}

func generateKey() []byte {
	key := make([]byte, 8) // DES key size is 8 bytes
	_, _ = io.ReadFull(rand.Reader, key)
	return key
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	a := param
	b := a + " SafeStuff"
	b = b[:len(b)-1] // take all but the last character
	return b
}
