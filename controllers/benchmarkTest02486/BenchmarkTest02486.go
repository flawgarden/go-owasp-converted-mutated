package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02486Controller struct {
	web.Controller
}

func (c *BenchmarkTest02486Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02486Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest02486")
	bar := doSomething(param)

	obj := []interface{}{"a", bar}
	c.Data["json"] = map[string]interface{}{"formatted": "Formatted like: " + obj[0].(string) + " and " + obj[1].(string)}
	c.ServeJSON()
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		sbxyz61588 := []rune(param)
		bar = string(append(sbxyz61588[:len(sbxyz61588)-1], 'Z'))
	}
	return bar
}
