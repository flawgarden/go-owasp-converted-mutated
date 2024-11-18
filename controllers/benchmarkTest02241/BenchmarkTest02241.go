package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02241Controller struct {
	web.Controller
}

func (c *BenchmarkTest02241Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02241Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02241")
	bar := doSomething(c.Ctx.Request, param)

	c.Ctx.Output.Body([]byte(bar))
}

func doSomething(request *http.Request, param string) string {
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
