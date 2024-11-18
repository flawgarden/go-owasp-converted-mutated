package controllers

import (
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02601Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02601Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02601Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02601Controller) doPost() {
	queryString := c.Ctx.Input.URI()
	paramval := "BenchmarkTest02601="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02601' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		c.Ctx.ResponseWriter.Write([]byte(bar))
	}
}

func doSomething(param string) string {
	return param // Здесь можно добавить логику экранирования, если необходимо
}
