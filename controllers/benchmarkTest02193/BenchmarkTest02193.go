package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest02193 struct{}

func (b *BenchmarkTest02193) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest02193) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02193")
	bar := b.doSomething(param)

	block, err := des.NewCipher(make([]byte, 8))
	if err != nil {
		http.Error(w, "Could not create cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Could not generate IV", http.StatusInternalServerError)
		return
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	input := []byte(bar)
	ciphertext := make([]byte, len(input))
	cfb.XORKeyStream(ciphertext, input)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(string(input)))
}

func (b *BenchmarkTest02193) doSomething(param string) string {
	bar := ""
	switch "B" {
	case "A":
		bar = param
	case "B":
		bar = "bob"
	case "C", "D":
		bar = param
	default:
		bar = "bob's your uncle"
	}
	return bar
}

func htmlEscape(str string) string {
	return html.EscapeString(str)
}
