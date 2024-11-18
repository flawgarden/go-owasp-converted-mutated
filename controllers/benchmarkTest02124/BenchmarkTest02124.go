package controllers

import (
	"regexp"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02124Controller struct {
	web.Controller
}

func (c *BenchmarkTest02124Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02124Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest02124")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		re := regexp.MustCompile("Z$")
		bar = re.ReplaceAllString(param, "Z")
	}
	return bar
}
