package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02494Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02494Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02494Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02494")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz57919 := []rune(param)
		bar = string(append(sbxyz57919[:len(sbxyz57919)-1], 'Z'))
	}

	return bar
}
