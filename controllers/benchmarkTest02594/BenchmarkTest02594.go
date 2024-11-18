package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02594Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02594Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02594Controller) Post() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Request.URL.Query().Encode()
	paramval := "BenchmarkTest02594="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		c.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02594' in query string."))
		return
	}

	param := ""
	param = queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(param, "&")

	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	c.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := ""

	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
