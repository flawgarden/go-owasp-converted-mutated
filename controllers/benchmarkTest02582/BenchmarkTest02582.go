package controllers

import (
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02582Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02582Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02582Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest02582="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02582' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param = param // simulate URL decoding with proper handling

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	var sbxyz18070 strings.Builder
	sbxyz18070.WriteString(param)
	bar := sbxyz18070.String() + "_SafeStuff"

	return bar
}
