package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01258Controller struct {
	web.Controller
}

func (c *BenchmarkTest01258Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01258Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01258Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01258")
	if param == "" {
		param = ""
	}

	bar := new(Test).DoSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.Output.Body([]byte("Formatted like: a and " + bar))
}

type Test struct{}

func (t *Test) DoSomething(param string) string {
	bar := ""
	if param != "" {
		bar = string(param)
	}
	return bar
}
