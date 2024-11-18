package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02579Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02579Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02579Controller) Post() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Request.URL.RawQuery
	paramVal := "BenchmarkTest02579="
	paramLoc := strings.Index(queryString, paramVal)
	if paramLoc == -1 {
		c.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02579")))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(c.Request, param)

	c.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	_, _ = fmt.Fprintf(c.ResponseWriter, bar, obj...)
}

func doSomething(req *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}
