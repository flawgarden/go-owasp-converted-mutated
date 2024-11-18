package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02236Controller struct {
	web.Controller
}

func (c *BenchmarkTest02236Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02236Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02236Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02236")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.WriteString(bar)
}

func doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
