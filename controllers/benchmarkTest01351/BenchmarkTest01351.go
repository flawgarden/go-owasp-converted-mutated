package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01351Controller struct {
	web.Controller
}

func (c *BenchmarkTest01351Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01351Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	params := c.Ctx.Request.URL.Query()
	param := params.Get("BenchmarkTest01351")

	bar := new(Test).doSomething(param)

	if bar != "" {
		c.Ctx.ResponseWriter.Write([]byte(bar))
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(param string) string {
	// Implement HTML escaping here
	// For example, using "html/template"
	escaped := param // replace with actual escaping logic
	return escaped
}
