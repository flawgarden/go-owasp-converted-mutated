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
	"net/url"
	"os"
)

type BenchmarkTest00120 struct{}

func (bt *BenchmarkTest00120) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else if r.Method == http.MethodPost {
		bt.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (bt *BenchmarkTest00120) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00120")
	param, _ = url.QueryUnescape(param)

	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1] + "Z"
	}

	block, err := des.NewCipher(generateKey())
	if err != nil {
		http.Error(w, "Encryption error", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "IV generation error", http.StatusInternalServerError)
		return
	}

	ciphertext := encrypt(block, iv, []byte(bar))
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	if err := saveToFile(encoded); err != nil {
		http.Error(w, "File save error", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(bar))
	w.Write([]byte(response))
}

func generateKey() []byte {
	key := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	return key
}

func encrypt(block cipher.Block, iv, plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func saveToFile(data string) error {
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + data + "\n"); err != nil {
		return err
	}
	return nil
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00120", &BenchmarkTest00120{})
	http.ListenAndServe(":8080", nil)
}
