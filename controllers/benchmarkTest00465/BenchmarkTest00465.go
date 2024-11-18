package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00465Controller struct {
	web.Controller
}

func (c *BenchmarkTest00465Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00465Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00465")
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	hashValue := md5.Sum([]byte(bar))
	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		c.Ctx.WriteString("Error opening file")
		return
	}
	defer fw.Close()

	encodedHash := base64.StdEncoding.EncodeToString(hashValue[:])
	if _, err := fw.WriteString("hash_value=" + encodedHash + "\n"); err != nil {
		c.Ctx.WriteString("Error writing to file")
		return
	}

	c.Ctx.WriteString(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(bar))))
	c.Ctx.WriteString("Hash Test executed")
}
