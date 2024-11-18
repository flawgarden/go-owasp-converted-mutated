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
)

type BenchmarkTest00447Controller struct {
	web.Controller
}

func (c *BenchmarkTest00447Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00447Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00447Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00447")
	bar := ""
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		http.Error(c.Ctx.ResponseWriter, "error generating key", http.StatusInternalServerError)
		return
	}

	block, _ := aes.NewCipher(key)
	aesGCM, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		http.Error(c.Ctx.ResponseWriter, "error generating nonce", http.StatusInternalServerError)
		return
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(bar), nil)
	encodedResult := base64.StdEncoding.EncodeToString(ciphertext)

	fileTarget := "passwordFile.txt"
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "error opening file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := f.WriteString("secret_value=" + encodedResult + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "error writing to file", http.StatusInternalServerError)
		return
	}

	responseMessage := "Sensitive value: '" + html.EscapeString(bar) + "' encrypted and stored<br/>"
	c.Ctx.ResponseWriter.Write([]byte(responseMessage))
}
