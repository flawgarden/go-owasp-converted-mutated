package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01594Controller struct {
	web.Controller
}

func (c *BenchmarkTest01594Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01594Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest01594")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	var sbxyz51154 = param
	bar := sbxyz51154 + "_SafeStuff"
	return bar
}
