package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type BenchmarkTest01148 struct{}

func (b *BenchmarkTest01148) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
		return
	}
	r.ParseForm()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01148")
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	block, err := des.NewCipher(make([]byte, 8)) // using a dummy key for this example
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Error creating IV", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, []byte(bar))

	encrypted := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", encrypted)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := param
	if (7*42)-86 > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
