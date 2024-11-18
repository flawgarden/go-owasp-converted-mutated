package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01634Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest01634Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	if r.Method == http.MethodGet {
		c.doPost()
	} else if r.Method == http.MethodPost {
		c.doPost()
	}
}

func (c *BenchmarkTest01634Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Request.URL.Query()
	param := queryString.Get("BenchmarkTest01634")
	if param == "" {
		c.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01634' in query string."))
		return
	}

	param, _ = url.QueryUnescape(param)
	bar := c.processParam(param)

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(c.ResponseWriter, "Failed to generate IV", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher([]byte("abcdefgh")) // DES key should be 8 bytes
	if err != nil {
		http.Error(c.ResponseWriter, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	padded := pad([]byte(bar), des.BlockSize)
	result := make([]byte, len(padded))
	encrypter.CryptBlocks(result, padded)

	fileData := fmt.Sprintf("secret_value=%s\n", base64.StdEncoding.EncodeToString(result))
	err = appendToFile("passwordFile.txt", fileData)
	if err != nil {
		http.Error(c.ResponseWriter, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(c.ResponseWriter, "Sensitive value: '%s' encrypted and stored<br/>", url.QueryEscape(bar))
}

func (c *BenchmarkTest01634Controller) processParam(param string) string {
	return param + " SafeStuff" // simplified processing
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func appendToFile(filename, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	return err
}
