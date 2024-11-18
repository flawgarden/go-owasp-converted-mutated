package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
	"text/template"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01228Controller struct {
	web.Controller
}

func (c *BenchmarkTest01228Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01228Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01228Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01228")
	if param == "" {
		param = ""
	}

	bar := testFunction(param)

	block, err := des.NewCipher(make([]byte, 8))
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error creating cipher"))
		return
	}

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error generating IV"))
		return
	}

	ciphertext := encrypt(block, iv, []byte(bar))

	fileTarget, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error opening file"))
		return
	}
	defer fileTarget.Close()

	_, err = fileTarget.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n")
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error writing to file"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + htmlEscape(string(bar)) + "' encrypted and stored<br/>"))
}

func testFunction(param string) string {
	return param
}

func encrypt(block cipher.Block, iv []byte, plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}
