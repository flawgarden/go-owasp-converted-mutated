package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01772Controller struct {
	web.Controller
}

func (c *BenchmarkTest01772Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01772Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01772")
	bar := testDoSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func testDoSomething(param string) string {
	sbxyz76680 := param + "_SafeStuff"
	return sbxyz76680
}
