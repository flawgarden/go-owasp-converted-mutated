package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type CryptoBenchmarkController struct {
	web.Controller
}

func (c *CryptoBenchmarkController) Get() {
	c.DoPost()
}

func (c *CryptoBenchmarkController) Post() {
	c.DoPost()
}

func (c *CryptoBenchmarkController) DoPost() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("X-Custom-Header")

	bar := c.performOperation(param)

	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating IV", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher([]byte("abcdefgh")) // 8 byte key for DES
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(ciphertext, []byte(bar))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.Output.Body([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", bar)))
}

func (c *CryptoBenchmarkController) performOperation(param string) string {
	bar := "safe!"
	mapData := map[string]interface{}{
		"keyA": "a-Value",
		"keyB": param,
		"keyC": "another-Value",
	}
	bar = fmt.Sprintf("%v", mapData["keyB"])

	return bar
}
