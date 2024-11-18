package controllers

import (
	"html/template"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02489Controller struct {
	web.Controller
}

func (c *BenchmarkTest02489Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02489Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	params := c.Ctx.Input.Query("BenchmarkTest02489")
	bar := doSomething(params)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	tmpl := template.Must(template.New("response").Parse(bar))
	tmpl.Execute(c.Ctx.ResponseWriter, obj)
}

func doSomething(param string) string {
	bar := param
	return bar
}
