package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02240Controller struct {
	web.Controller
}

func (c *BenchmarkTest02240Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02240Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	m := c.Ctx.Request.URL.Query()
	if len(m) > 0 {
		if values, ok := m["BenchmarkTest02240"]; ok {
			param = values[0]
		}
	}

	bar := doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(request *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
