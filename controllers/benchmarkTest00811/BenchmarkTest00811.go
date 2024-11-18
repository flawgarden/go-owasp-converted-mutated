package controllers

import (
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00811Controller struct {
	web.Controller
}

func (c *BenchmarkTest00811Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest00811Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest00811="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParam(queryString, paramval)
	}

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest00811' in query string."))
		return
	}

	param := extractParam(queryString, paramLoc, paramval)
	param = decodeURL(param)

	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func findParam(queryString, paramval string) int {
	return strings.Index(queryString, paramval)
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc+paramLoc]
	}
	return param
}

func decodeURL(param string) string {
	decoded, err := url.QueryUnescape(param)
	if err != nil {
		return param
	}
	return decoded
}
