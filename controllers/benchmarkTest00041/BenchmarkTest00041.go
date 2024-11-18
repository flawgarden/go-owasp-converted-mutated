package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00041Controller struct {
	web.Controller
}

func (c *BenchmarkTest00041Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00041Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00041Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest00041")
	if param == "" {
		param = ""
	}

	length := len(param)
	if length > 0 {
		c.Ctx.ResponseWriter.Write([]byte(param))
	}
}
