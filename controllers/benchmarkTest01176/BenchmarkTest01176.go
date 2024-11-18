package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01176Controller struct {
	web.Controller
}

func (c *BenchmarkTest01176Controller) Get() {
	c.post()
}

func (c *BenchmarkTest01176Controller) Post() {
	c.post()
}

func (c *BenchmarkTest01176Controller) post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := "Formatted like: %s and %s."
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(output, "a", bar)))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C':
	case 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
