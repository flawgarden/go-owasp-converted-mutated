package controllers

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01171Controller struct {
	web.Controller
}

func (c *BenchmarkTest01171Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01171Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")

	param := ""
	headers := c.Ctx.Request.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.PathUnescape(param)

	bar := new(Test).doSomething(param)

	c.Ctx.Output.Body([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		sbxyz83647 := []rune(param)
		bar = string(sbxyz83647[:len(param)-1]) + "Z"
	}
	return bar
}
