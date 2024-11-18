package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01595 struct {
	web.Controller
}

func (c *BenchmarkTest01595) Get() {
	c.Post()
}

func (c *BenchmarkTest01595) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest01595")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.Output.Body([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := escapeHtml(param)
	return bar
}

func escapeHtml(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(s)
}
