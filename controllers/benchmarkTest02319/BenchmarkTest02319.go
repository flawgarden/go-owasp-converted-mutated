package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02319Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02319Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02319Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02319Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02319" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.Output.Body([]byte(bar))
}

func doSomething(param string) string {
	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
