package controllers

import (
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02622Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02622Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02622Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest02622="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02622' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	c.Ctx.Request.Context().Value("session").(map[string]interface{})[bar] = "10340"

	c.Ctx.ResponseWriter.Write([]byte("Item: '" + bar + "' with value: '10340' saved in session."))
}

func doSomething(param string) string {
	return param + "_SafeStuff"
}
