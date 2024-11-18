package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00894 struct {
	web.Controller
}

func (c *BenchmarkTest00894) Get() {
	c.Post()
}

func (c *BenchmarkTest00894) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00894")

	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		c.Ctx.ResponseWriter.Write([]byte(bar[:length]))
	}
}
