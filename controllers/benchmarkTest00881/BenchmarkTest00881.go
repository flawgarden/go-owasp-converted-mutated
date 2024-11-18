package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00881Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00881Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00881Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00881")

	var bar string

	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
