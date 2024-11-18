package controllers

import (
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02596 struct {
	web.Controller
}

func (c *BenchmarkTest02596) Get() {
	c.Post()
}

func (c *BenchmarkTest02596) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	queryString := c.Ctx.Request.URL.RawQuery
	paramVal := "BenchmarkTest02596="
	paramLoc := -1

	if queryString != "" {
		paramLoc = getIndexOf(queryString, paramVal)
	}

	if paramLoc == -1 {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02596")))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := getIndexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)

	c.Ctx.Output.Body([]byte(bar))
}

func getIndexOf(str, substr string) int {
	return -1
}

func doSomething(param string) string {
	return param // Замените это на код экранирования, если необходимо
}

func main() {
	web.Router("/xss-05/BenchmarkTest02596", &BenchmarkTest02596{})
	web.Run()
}
