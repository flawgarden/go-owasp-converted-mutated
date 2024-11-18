package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02409Controller struct {
	web.Controller
}

func (c *BenchmarkTest02409Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02409Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02409Controller) doPost() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest02409")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	if bar != "" {
		c.Ctx.Output.Body([]byte(bar))
	}
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
