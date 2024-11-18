package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02480Controller struct {
	web.Controller
}

func (c *BenchmarkTest02480Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest02480Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02480")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	num := 106
	bar := ""
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}
