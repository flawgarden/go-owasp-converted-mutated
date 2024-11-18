package controllers

import (
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01177Controller struct {
	web.Controller
}

func (c *BenchmarkTest01177Controller) Get() {
	c.Ctx.WriteString("Method Not Allowed")
}

func (c *BenchmarkTest01177Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")

	param := ""
	headers := c.Ctx.Request.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0]
	}
	param, _ = url.PathUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	obj := []interface{}{"a", bar}
	c.Ctx.Output.Context.ResponseWriter.Write([]byte("Formatted like: " + obj[0].(string) + " and " + obj[1].(string) + "."))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
