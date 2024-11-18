package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02023Controller struct {
	web.Controller
}

func (c *BenchmarkTest02023Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02023Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02023"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	block, err := des.NewCipher([]byte("12345678")) // Example key
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	mode := cipher.NewCBCEncrypter(block, []byte("12345678")) // Example IV
	mode.CryptBlocks(ciphertext, []byte(bar))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("secret_value=%x\n", ciphertext)); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", bar)))
}

func doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
