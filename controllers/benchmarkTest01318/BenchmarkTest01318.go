package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

type BenchmarkTest01318 struct{}

func (bt *BenchmarkTest01318) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (bt *BenchmarkTest01318) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01318")
	bar := test(param)

	block, err := des.NewCipher(randomKey())
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := encrypt(block, []byte(bar))

	fileTarget := "passwordFile.txt"
	writeToFile(fileTarget, base64.StdEncoding.EncodeToString(ciphertext))

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(bar))
}

func randomKey() []byte {
	key := make([]byte, 8) // DES requires 8 byte keys
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		panic(err)
	}
	return key
}

func encrypt(block cipher.Block, plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, block.BlockSize())
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func writeToFile(filePath string, content string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString("secret_value=" + content + "\n")
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}

func test(param string) string {
	return fmt.Sprintf("%s SafeStuff", param)
}

func main() {
	http.Handle("/crypto-01/BenchmarkTest01318", &BenchmarkTest01318{})
	http.ListenAndServe(":8080", nil)
}
