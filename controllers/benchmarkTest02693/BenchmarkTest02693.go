package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02693Controller struct {
	web.Controller
}

func (c *BenchmarkTest02693Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02693Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02693")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.Output.Body([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz14220 := []rune(param)
		bar = string(append(sbxyz14220[:len(sbxyz14220)-1], 'Z'))
	}
	return bar
}
