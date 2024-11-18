package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01255Controller struct {
	web.Controller
}

func (c *BenchmarkTest01255Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01255Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01255Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01255")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := escapeHtml(param)
	return bar
}

func escapeHtml(s string) string {
	// Implementation of HTML escape (you can use a library or implement your own)
	return s // Placeholder for actual escaping logic
}
