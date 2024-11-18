package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01589Controller struct {
	web.Controller
}

func (c *BenchmarkTest01589Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01589Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")

	values := c.GetStrings("BenchmarkTest01589")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := new(Test).doSomething(param)

	obj := []interface{}{bar, "b"}
	_, _ = c.Ctx.ResponseWriter.Write([]byte("Formatted like: " + obj[0].(string) + " and " + obj[1].(string) + "."))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
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
