package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02318Controller struct {
	web.Controller
}

func (c *BenchmarkTest02318Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02318Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02318Controller) doPost() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	param := ""

	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		for _, value := range values {
			if value == "BenchmarkTest02318" {
				param = name
				break
			}
		}
	}

	bar := doSomething(param)
	c.Ctx.Output.Body([]byte(bar))
}

func doSomething(param string) string {
	return strings.ReplaceAll(param, "<", "&lt;")
}
