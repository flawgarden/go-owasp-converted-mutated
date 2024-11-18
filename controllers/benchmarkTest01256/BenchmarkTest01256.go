package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01256Controller struct {
	web.Controller
}

func (c *BenchmarkTest01256Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01256Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01256")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.WriteString(bar)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
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
