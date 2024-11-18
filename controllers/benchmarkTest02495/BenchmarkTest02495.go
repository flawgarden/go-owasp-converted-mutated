package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02495Controller struct {
	web.Controller
}

func (c *BenchmarkTest02495Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02495Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02495")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		c.Ctx.ResponseWriter.Write([]byte(bar))
	}
}

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

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
