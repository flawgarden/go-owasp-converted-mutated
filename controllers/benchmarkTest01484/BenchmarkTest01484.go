package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"html"
	"io"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01484Controller struct {
	web.Controller
}

func (c *BenchmarkTest01484Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01484Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01484Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01484")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	block, _ := des.NewCipher(make([]byte, 8)) // DES requires 8 byte keys
	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating IV", http.StatusInternalServerError)
		return
	}

	ciphertext := encrypt([]byte(bar), block, iv)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + html.EscapeString(bar) + "' encrypted and stored<br/>"))
}

func encrypt(plainText []byte, block cipher.Block, iv []byte) []byte {
	mode := cipher.NewCBCEncrypter(block, iv)
	paddedText := pad(plainText, block.BlockSize())
	ciphertext := make([]byte, len(paddedText))
	mode.CryptBlocks(ciphertext, paddedText)
	return ciphertext
}

func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}
