package controllers

import (
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01339Controller struct {
	web.Controller
}

func (c *BenchmarkTest01339Controller) Get() {
	c.handleRequest()
}

func (c *BenchmarkTest01339Controller) Post() {
	c.handleRequest()
}

func (c *BenchmarkTest01339Controller) handleRequest() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest01339")

	bar := testDoSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func testDoSomething(request *http.Request, param string) string {
	bar := htmlEscape(param)
	return bar
}

func htmlEscape(s string) string {
	return strconv.QuoteToASCII(s)
}
