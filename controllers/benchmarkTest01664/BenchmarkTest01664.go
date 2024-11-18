package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01664Controller struct {
	web.Controller
}

func (c *BenchmarkTest01664Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01664Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest01664="
	paramLoc := -1
	if queryString != "" {
		paramLoc = c.findParamLocation(queryString, paramval)
	}

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01664' in query string."))
		return
	}

	param := c.extractParam(queryString, paramLoc, paramval)

	bar := c.doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func (c *BenchmarkTest01664Controller) findParamLocation(queryString, paramval string) int {
	return -1 // logic to find the parameter location
}

func (c *BenchmarkTest01664Controller) extractParam(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := c.findAmpersandLocation(queryString, paramLoc)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	return param
}

func (c *BenchmarkTest01664Controller) findAmpersandLocation(queryString string, paramLoc int) int {
	return -1 // logic to find the next ampersand
}

func (c *BenchmarkTest01664Controller) doSomething(param string) string {
	return escapeHTML(param) // use proper HTML escaping here
}

func escapeHTML(param string) string {
	return param // implement HTML escaping logic
}
