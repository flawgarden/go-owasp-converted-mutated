package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02658Controller struct {
	web.Controller
}

func (c *BenchmarkTest02658Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02658Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02658")

	bar := doSomething(param)

	block, err := des.NewCipher(generateKey())
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	cipherText, err := encrypt(block, []byte(bar))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	encodedText := base64.StdEncoding.EncodeToString(cipherText)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", encodedText)))
}

func doSomething(param string) string {
	num := 196
	if (500/42)+num > 200 {
		return param
	}
	return "This should never happen"
}

func generateKey() []byte {
	return []byte("abcdefgh") // DES requires a 8 byte key
}

func encrypt(block cipher.Block, plaintext []byte) ([]byte, error) {
	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)

	return ciphertext, nil
}
