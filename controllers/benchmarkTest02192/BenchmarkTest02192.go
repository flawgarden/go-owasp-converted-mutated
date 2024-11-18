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
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02192Controller struct {
	web.Controller
}

func (c *BenchmarkTest02192Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02192Controller) Post() {
	defer c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02192")

	bar := doSomething(param)

	block, err := des.NewCipher(generateRandomBytes(8))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Cipher creation error", http.StatusInternalServerError)
		return
	}

	ciphertext := encryptAES([]byte(bar), block)

	filePath := "passwordFile.txt"
	appendToFile(filePath, "secret_value="+base64.StdEncoding.EncodeToString(ciphertext))

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(bar))))
}

func doSomething(param string) string {
	valuesList := []string{"safe", param, "moresafe"}
	valuesList = valuesList[1:] // remove the 1st safe value
	return valuesList[0]        // get the param value
}

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	return b
}

func encryptAES(plainText []byte, block cipher.Block) []byte {
	ciphertext := make([]byte, len(plainText))
	blockMode := cipher.NewCBCEncrypter(block, generateRandomBytes(block.BlockSize()))
	blockMode.CryptBlocks(ciphertext, plainText)
	return ciphertext
}

func appendToFile(filePath, text string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	if _, err := f.WriteString(text + "\n"); err != nil {
		return
	}
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
