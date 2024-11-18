package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"html"
	"io"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type CryptoController struct {
	web.Controller
}

func (c *CryptoController) Get() {
	c.DoPost()
}

func (c *CryptoController) Post() {
	c.DoPost()
}

func (c *CryptoController) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02191")

	bar := doSomething(param)

	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating key", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating nonce", http.StatusInternalServerError)
		return
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(bar), nil)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	fileTarget := "passwordFile.txt"
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + encodeForHTML(param) + "' encrypted and stored<br/>"))
}

func doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}

func encodeForHTML(input string) string {
	return html.EscapeString(input)
}
