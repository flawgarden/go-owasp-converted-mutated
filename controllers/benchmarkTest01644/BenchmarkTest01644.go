package controllers

import (
	"net/url"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01644Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01644Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01644Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest01644="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01644' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	fileName := "testfiles/" + bar
	fis, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open FileInputStream on file: '" + fileName + "'"))
		c.Ctx.ResponseWriter.Write([]byte("Problem getting FileInputStream: " + err.Error()))
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n" + string(b[:size])))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}
