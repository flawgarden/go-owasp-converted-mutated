package controllers

import (
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00152Controller struct {
	web.Controller
}

func (c *BenchmarkTest00152Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00152Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00152Controller) doPost() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	referer := c.Ctx.Request.Header.Get("Referer")
	if referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

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

	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.Output.Body([]byte(strings.Join(strings.Split(bar, ""), "")))
}
