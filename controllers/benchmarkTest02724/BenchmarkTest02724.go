package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02724Controller struct {
	web.Controller
}

func (c *BenchmarkTest02724Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02724Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02724")
	bar := doSomething(param)

	c.SetSession(bar, "10340")
	c.Ctx.Output.Body([]byte("Item: '" + bar + "' with value: 10340 saved in session."))
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
