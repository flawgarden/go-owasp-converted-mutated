package controllers

import (
	"html"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01790Controller struct {
	web.Controller
}

func (c *BenchmarkTest01790Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01790Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01790")
	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.WriteString("Parameter value: " + bar)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return escapeHtml(param)
}

func escapeHtml(input string) string {
	// Здесь необходимо использовать библиотеку для экранирования HTML
	// Например, html.EscapeString в стандартной библиотеке
	return html.EscapeString(input)
}
