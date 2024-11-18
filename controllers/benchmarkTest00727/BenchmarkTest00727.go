package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00727Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00727Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00727Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00727")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
