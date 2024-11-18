package controllers

import (
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01767Controller struct {
	web.Controller
}

func (c *BenchmarkTest01767Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01767Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01767")

	bar := testDoSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(strconv.Quote(bar)))
}

func testDoSomething(param string) string {
	var bar string
	num := 86

	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
