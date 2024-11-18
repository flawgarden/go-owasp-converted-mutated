package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00611Controller struct {
	web.Controller
}

func (c *BenchmarkTest00611Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00611Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00611")
	if param == "" {
		param = ""
	}

	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	// DES requires 8 byte keys
	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating IV", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher(iv)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, []byte(bar))

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + encoded + "' encrypted and stored<br/>"))
}
