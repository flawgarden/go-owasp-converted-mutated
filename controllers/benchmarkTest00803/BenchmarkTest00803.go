package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00803Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00803Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doGet()
}

func (c *BenchmarkTest00803Controller) doGet() {
	c.doPost()
}

func (c *BenchmarkTest00803Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Request.URL.RawQuery
	paramval := "BenchmarkTest00803="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(c.ResponseWriter, "getQueryString() couldn't find expected parameter 'BenchmarkTest00803' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := param

	c.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.ResponseWriter.Write([]byte(bar))
}
