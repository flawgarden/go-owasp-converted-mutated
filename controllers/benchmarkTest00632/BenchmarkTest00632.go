package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00632Controller struct {
	web.Controller
}

func (c *BenchmarkTest00632Controller) Get() {
	c.handleRequest()
}

func (c *BenchmarkTest00632Controller) Post() {
	c.handleRequest()
}

func (c *BenchmarkTest00632Controller) handleRequest() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Query("BenchmarkTest00632")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}

	hash := sha512.New384()
	input := []byte(bar)
	hash.Write(input)

	result := hash.Sum(nil)
	hashValue := base64.StdEncoding.EncodeToString(result)

	c.Ctx.Output.Body([]byte("Sensitive value '" + htmlEncode(string(input)) + "' hashed and stored<br/>"))
	c.Ctx.Output.Body([]byte("Hash Test executed with value: " + hashValue))
}

func htmlEncode(input string) string {
	return strings.ReplaceAll(input, "&", "&amp;")
}
