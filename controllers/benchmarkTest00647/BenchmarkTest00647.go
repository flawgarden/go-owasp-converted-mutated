package controllers

import (
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00647Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00647Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00647Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00647Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00647")
	if param == "" {
		param = ""
	}

	bar := sanitizeForHTML(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func sanitizeForHTML(input string) string {
	return strings.ReplaceAll(input, "<", "&lt;")
}
