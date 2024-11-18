package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00380Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00380Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00380Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00380Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00380")
	if param == "" {
		param = ""
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
