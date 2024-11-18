package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00651Controller struct {
	web.Controller
}

func (c *BenchmarkTest00651Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00651Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00651Controller) doPost() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")

	param := c.GetString("BenchmarkTest00651")
	if param == "" {
		param = ""
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	length := 1
	if bar != "" {
		length = len(bar)
		c.Ctx.Output.Body([]byte(bar[:length]))
	}
}
