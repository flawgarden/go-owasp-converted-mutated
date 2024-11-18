package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02399Controller struct {
	web.Controller
}

func (c *BenchmarkTest02399Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02399Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02399")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
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
