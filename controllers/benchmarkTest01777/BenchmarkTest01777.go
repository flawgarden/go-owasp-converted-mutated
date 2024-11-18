package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01777Controller struct {
	web.Controller
}

func (c *BenchmarkTest01777Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01777Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest01777")
	bar := new(Test).doSomething(param)

	length := 1
	if bar != "" {
		length = len(bar)
		c.Ctx.ResponseWriter.Write([]byte(bar[:length]))
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
