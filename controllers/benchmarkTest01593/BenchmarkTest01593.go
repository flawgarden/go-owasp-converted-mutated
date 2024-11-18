package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01593Controller struct {
	web.Controller
}

func (c *BenchmarkTest01593Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01593Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01593Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.Ctx.Input.Query("BenchmarkTest01593")
	var param string
	if values != "" {
		param = values
	} else {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.Output.Body([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
