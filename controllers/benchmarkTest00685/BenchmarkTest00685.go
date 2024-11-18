package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/crypto-00/BenchmarkTest00685", BenchmarkTest00685)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest00685(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This input source requires a POST.", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	values := r.Form["BenchmarkTest00685"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := param
	if param != "" && len(param) > 1 {
		bar = fmt.Sprintf("%sZ", param[:len(param)-1])
	}

	key := make([]byte, 8)
	if _, err := rand.Read(key); err != nil {
		http.Error(w, "Error generating key.", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher(key)
	if err != nil {
		http.Error(w, "Error creating cipher.", http.StatusInternalServerError)
		return
	}

	method := cipher.NewCBCEncrypter(block, key)
	padCount := des.BlockSize - len([]byte(bar))%des.BlockSize
	pad := bytes.Repeat([]byte{byte(padCount)}, padCount)
	input := append([]byte(bar), pad...)

	cipherText := make([]byte, len(input))
	method.CryptBlocks(cipherText, input)

	encoded := base64.StdEncoding.EncodeToString(cipherText)

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(bar))
	fmt.Fprintf(w, "Encrypted value: '%s'<br/>", encoded)
}
