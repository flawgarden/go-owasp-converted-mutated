package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01481Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01481Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01481Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01481")
	if param == "" {
		param = ""
	}

	bar := testDoSomething(param)

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating key", http.StatusInternalServerError)
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, aes.BlockSize+len([]byte(bar)))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating IV", http.StatusInternalServerError)
		return
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(bar))

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + encoded + "' encrypted and stored<br/>"))
}

func testDoSomething(param string) string {
	bar := "safe!"
	map69587 := make(map[string]interface{})
	map69587["keyA-69587"] = "a_Value"
	map69587["keyB-69587"] = param
	map69587["keyC"] = "another_Value"

	bar = map69587["keyB-69587"].(string)
	bar = map69587["keyA-69587"].(string)

	return bar
}
