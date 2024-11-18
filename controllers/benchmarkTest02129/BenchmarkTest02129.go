package controllers

import (
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02129Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02129Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02129Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02129")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	_, _ = c.Ctx.ResponseWriter.Write([]byte("Formatted like: " + obj[0].(string) + " and " + obj[1].(string) + "."))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}
