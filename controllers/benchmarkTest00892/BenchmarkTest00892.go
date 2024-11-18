package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00892Controller struct {
	web.Controller
}

func (c *BenchmarkTest00892Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00892Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00892Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00892")

	bar := param

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		c.Ctx.ResponseWriter.Write([]byte(bar[:length]))
	}
}
