package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Initialization logic here if needed
}

type BenchmarkTest01147 struct {
}

func (b *BenchmarkTest01147) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest01147) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01147")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(param)

	block, err := aes.NewCipher(generateKey())
	if err != nil {
		http.Error(w, "Error initializing cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	ciphertext := b.encrypt(bar, block, iv)

	err = b.storeEncryptedValue(ciphertext)
	if err != nil {
		http.Error(w, "Error storing encrypted value", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", url.QueryEscape(string(ciphertext)))
}

func (b *BenchmarkTest01147) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1] + "Z"
	}
	return bar
}

func (b *BenchmarkTest01147) encrypt(data string, block cipher.Block, iv []byte) []byte {
	ciphertext := make([]byte, aes.BlockSize+len(data))
	copy(ciphertext[:aes.BlockSize], iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

	return ciphertext
}

func (b *BenchmarkTest01147) storeEncryptedValue(data []byte) error {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	encodedValue := base64.StdEncoding.EncodeToString(data)
	_, err = db.Exec("INSERT INTO storage_table (encrypted_data) VALUES (?)", encodedValue)
	return err
}

func generateKey() []byte {
	return []byte("thisisaverysecret") // Use a secure key generator in production
}
