package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest02662Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02662Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02662Controller) Post() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Request.URL.Query().Get("BenchmarkTest02662")
	bar := doSomething(param)

	key := make([]byte, 32) // AES-256
	if _, err := os.ReadFile("path/to/keyfile"); err == nil {
		copy(key, []byte("your32characterlongkey!"))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(c.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(c.ResponseWriter, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	input := []byte(bar)
	ciphertext := gcm.Seal(nonce, nonce, input, nil)

	file, err := os.OpenFile("path/to/passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		http.Error(c.ResponseWriter, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("secret_value=%s\n", base64.StdEncoding.EncodeToString(ciphertext))); err != nil {
		http.Error(c.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(string(input)))))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}
