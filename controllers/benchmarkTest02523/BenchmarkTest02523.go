package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02523Controller struct {
	web.Controller
}

func (c *BenchmarkTest02523Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02523Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02523")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

	c.Ctx.Output.Context.ResponseWriter.Write([]byte("Item: '" + encodeForHTML(bar) + "' with value: 10340 saved in session."))
}

func doSomething(param string) string {
	var bar string
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

func encodeForHTML(input string) string {
	// Implementation of HTML encoding (can be using a library)
	return input // Placeholder, should be replaced with actual HTML encoding
}
