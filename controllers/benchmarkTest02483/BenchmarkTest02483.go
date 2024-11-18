package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02483Controller struct {
	web.Controller
}

func (c *BenchmarkTest02483Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02483Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	values := c.GetStrings("BenchmarkTest02483")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	return bar
}
