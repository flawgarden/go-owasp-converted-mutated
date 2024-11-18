package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02509Controller struct {
	web.Controller
}

func (c *BenchmarkTest02509Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02509Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02509")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.WriteString("Parameter value: " + bar)
}

func doSomething(param string) string {
	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
