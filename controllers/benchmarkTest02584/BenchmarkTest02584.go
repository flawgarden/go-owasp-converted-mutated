package controllers

import (
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02584Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02584Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02584Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.Query().Encode()
	paramval := "BenchmarkTest02584="

	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02584' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(param, "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}

	param = decodeURIComponent(param)

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	var sbxyz36210 strings.Builder
	sbxyz36210.WriteString(param)
	bar := sbxyz36210.String() + "_SafeStuff"
	return bar
}

func decodeURIComponent(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return decoded
}
