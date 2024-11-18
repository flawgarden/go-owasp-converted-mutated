package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00260Controller struct {
	web.Controller
}

func (c *BenchmarkTest00260Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00260Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00260Controller) DoPost() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest00260"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := htmlEscape(param)

	block, err := des.NewCipher([]byte("example_key")) // Example key
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Decrypt error", http.StatusInternalServerError)
		return
	}

	var input []byte
	if strings.TrimSpace(bar) == "" {
		input = []byte{'?'}
	} else {
		input = []byte(bar)
	}

	ciphertext := make([]byte, len(input))
	cipher.NewCBCEncrypter(block, []byte("example_iv")).CryptBlocks(ciphertext, input)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "File open error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("secret_value=%s\n", string(ciphertext))); err != nil {
		http.Error(c.Ctx.ResponseWriter, "File write error", http.StatusInternalServerError)
		return
	}

	c.Ctx.Output.Body([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", string(input))))
}

func htmlEscape(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "&", "&amp;"), "<", "&lt;")
}
