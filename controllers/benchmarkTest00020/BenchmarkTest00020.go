package controllers

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00020Controller struct {
	web.Controller
}

func (c *BenchmarkTest00020Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00020Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00020")

	if param == "" {
		param = ""
	}

	// 8-byte initialization vector
	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		c.Ctx.WriteString("Error generating IV")
		return
	}

	block, err := des.NewCipher([]byte("12345678")) // DES requires 8-byte key
	if err != nil {
		c.Ctx.WriteString("Problem executing crypto - " + err.Error())
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	padded := pad([]byte(param))
	ciphertext := make([]byte, len(padded))
	mode.CryptBlocks(ciphertext, padded)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.Ctx.WriteString("Error opening file")
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		c.Ctx.WriteString("Error writing to file")
		return
	}

	c.Ctx.WriteString(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(param)))
}

func pad(src []byte) []byte {
	padding := des.BlockSize - len(src)%des.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
