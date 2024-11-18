package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01775Controller struct {
	web.Controller
}

func (c *BenchmarkTest01775Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01775Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01775Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01775")
	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
