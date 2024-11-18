package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00889Controller struct {
	web.Controller
}

func (c *BenchmarkTest00889Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00889Controller) Post() {
	c.Data["content_type"] = "text/html;charset=UTF-8"
	param := c.GetString("BenchmarkTest00889")
	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
