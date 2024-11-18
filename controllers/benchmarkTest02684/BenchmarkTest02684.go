package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02684Controller struct {
	web.Controller
}

func (c *BenchmarkTest02684Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02684Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02684")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
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
