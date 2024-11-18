package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00720Controller struct {
	web.Controller
}

func (c *BenchmarkTest00720Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest00720Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00720")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

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

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, _ = c.Ctx.ResponseWriter.Write([]byte(bar))
}
