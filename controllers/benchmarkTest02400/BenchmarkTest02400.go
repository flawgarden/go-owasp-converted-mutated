package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02400Controller struct {
	web.Controller
}

func (c *BenchmarkTest02400Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02400Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02400")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := struct {
		Formatted string `json:"formatted"`
	}{
		Formatted: "Formatted like: a and " + bar,
	}
	c.Data["json"] = output
	c.ServeJSON()
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}
