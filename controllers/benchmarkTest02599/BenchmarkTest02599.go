package controllers

import (
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02599Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02599Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02599Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02599Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramVal := "BenchmarkTest02599="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramVal)
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02599' in query string."))
		return
	}

	param := extractParam(queryString, paramLoc, paramVal)
	param = decodeParam(param)

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func findParamLocation(queryString, paramVal string) int {
	return strings.Index(queryString, paramVal)
}

func extractParam(queryString string, paramLoc int, paramVal string) string {
	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	return param
}

func decodeParam(param string) string {
	decodedParam, _ := url.QueryUnescape(param)
	return decodedParam
}

func doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
