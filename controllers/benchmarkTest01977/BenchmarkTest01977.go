package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type CryptoTestController struct {
	web.Controller
}

func (c *CryptoTestController) Get() {
	c.DoPost()
}

func (c *CryptoTestController) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	names := c.Ctx.Request.Header
	for name := range names {
		if !isCommonHeader(name) {
			param = name
			break
		}
	}

	bar := doSomething(param)

	key := generateKey()
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not generate IV", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not create cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := encrypt(block, iv, []byte(bar))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("secret_value=%s\n", ciphertext)); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not write to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", bar)))
}

func generateKey() []byte {
	key := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil
	}
	return key
}

func encrypt(block cipher.Block, iv []byte, plaintext []byte) string {
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return fmt.Sprintf("%x", ciphertext)
}

func isCommonHeader(header string) bool {
	commonHeaders := []string{"Content-Type", "User-Agent", "Accept"}
	for _, h := range commonHeaders {
		if h == header {
			return true
		}
	}
	return false
}

func doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe
	var bar string
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
	return bar
}
