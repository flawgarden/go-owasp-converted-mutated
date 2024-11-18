package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01511Controller struct {
	web.Controller
}

func (c *BenchmarkTest01511Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01511Controller) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01511")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	response.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{bar, "b"}
	response.Write([]byte(fmt.Sprintf("Formatted like: %1$s and %2$s.", obj...)))
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
