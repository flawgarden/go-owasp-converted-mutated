package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02311Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02311Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02311Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02311Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02311" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(param)

	hash := md5.Sum([]byte(bar))
	hashValue := base64.StdEncoding.EncodeToString(hash[:])

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	if _, err := fw.WriteString(fmt.Sprintf("hash_value=%s\n", hashValue)); err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEscape(bar))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func htmlEscape(s string) string {
	return strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	).Replace(s)
}
