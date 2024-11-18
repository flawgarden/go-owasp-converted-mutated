package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00126Controller struct {
	web.Controller
}

func (c *BenchmarkTest00126Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00126Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest00126")
	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}

	key, err := aes.NewCipher([]byte("examplekey123456")) // 16 bytes key for AES-128
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, 12) // for GCM
	ciphertext := make([]byte, len([]byte(bar)))
	aesGCM, err := cipher.NewGCM(key)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to create GCM", http.StatusInternalServerError)
		return
	}
	copy(ciphertext, aesGCM.Seal(nonce, nonce, []byte(bar), nil))

	fileTarget, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer fileTarget.Close()

	if _, err := fileTarget.WriteString(fmt.Sprintf("secret_value=%s\n", ciphertext)); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", string(ciphertext))))
}
