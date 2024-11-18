package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type BenchmarkTest01015 struct{}

func (b *BenchmarkTest01015) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This input source requires a POST.", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest01015")
	param = decodeParam(param)

	bar := b.processInput(param)

	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher([]byte("12345678"))
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	padded := pad([]byte(bar))
	ciphertext := make([]byte, len(padded))
	mode.CryptBlocks(ciphertext, padded)

	encrypted := base64.StdEncoding.EncodeToString(ciphertext)

	file, _ := json.Marshal(map[string]string{"secret_value": encrypted})
	w.Write([]byte("Sensitive value encrypted and stored: " + string(file)))
}

func (b *BenchmarkTest01015) processInput(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func decodeParam(param string) string {
	decoded, _ := base64.StdEncoding.DecodeString(param)
	return string(decoded)
}

func pad(src []byte) []byte {
	padding := des.BlockSize - len(src)%des.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func main() {
	http.Handle("/crypto-01/BenchmarkTest01015", &BenchmarkTest01015{})
	http.ListenAndServe(":8080", nil)
}
