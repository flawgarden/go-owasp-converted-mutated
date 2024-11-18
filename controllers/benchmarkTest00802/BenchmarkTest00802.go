package controllers

import (
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00802Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00802Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00802Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest00802="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest00802' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := ""
	if param != "" {
		bar = string([]byte(param))
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
