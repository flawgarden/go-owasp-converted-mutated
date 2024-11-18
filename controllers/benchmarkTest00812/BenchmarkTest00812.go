package controllers

import (
	"fmt"
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00812Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00812Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00812Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00812Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Input.URI()
	paramVal := "BenchmarkTest00812="
	paramLoc := strings.Index(queryString, paramVal)

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00812")))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	var bar string
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	c.Ctx.ResponseWriter.Write([]byte(bar))
}
