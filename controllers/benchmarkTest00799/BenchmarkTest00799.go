package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00799Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00799Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest00799Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := c.Request.URL.RawQuery
	paramval := "BenchmarkTest00799="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		c.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00799")))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}

	c.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<body>\n<p>Formatted like: %s and %s.\n</p>\n</body>\n</html>", obj[0], obj[1])
	c.ResponseWriter.Write([]byte(output))
}
