package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02608Controller struct {
	web.Controller
}

func (c *BenchmarkTest02608Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02608Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest02608="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02608' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte("Parameter value: " + bar))
}

func doSomething(param string) string {
	bar := param
	return bar
}
