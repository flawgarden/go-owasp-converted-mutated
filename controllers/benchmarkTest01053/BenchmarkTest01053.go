package controllers

import (
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01053Controller struct {
	web.Controller
}

func (c *BenchmarkTest01053Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01053Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := c.Ctx.Request.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
