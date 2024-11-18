package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01335Controller struct {
	web.Controller
}

func (c *BenchmarkTest01335Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01335Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01335")

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := "<!DOCTYPE html>\n<html>\n<body>\n<p>"
	output += "Formatted like: a and " + bar + "."
	output += "\n</p>\n</body>\n</html>"
	c.Ctx.ResponseWriter.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	m := make(map[string]interface{})
	m["keyA"] = "a-Value"
	m["keyB"] = param
	m["keyC"] = "another-Value"
	bar = m["keyB"].(string)
	return bar
}
