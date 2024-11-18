package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02125Controller struct {
	web.Controller
}

func (c *BenchmarkTest02125Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02125Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02125Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02125")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	return param // здесь вы можете использовать экранирование, если необходимо
}
