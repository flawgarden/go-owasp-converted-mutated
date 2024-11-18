package controllers

import (
	"html"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02570Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02570Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02570Controller) Post() {
	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest02570="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.Output.Body([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02570' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	param = decode(param)

	bar := doSomething(param)
	fileName := filepath.Join("testfiles", bar)

	is, err := os.Open(fileName)
	if err != nil {
		c.Ctx.Output.Body([]byte("Couldn't open InputStream on file: '" + fileName + "'"))
		return
	}
	defer is.Close()

	b := make([]byte, 1000)
	size, _ := is.Read(b)

	c.Ctx.Output.Body([]byte("The beginning of file: '" + escapeHTML(fileName) + "' is:\n\n" + escapeHTML(string(b[:size]))))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}

func decode(param string) string {
	decoded, err := url.QueryUnescape(param)
	if err != nil {
		return param
	}
	return decoded
}

func escapeHTML(s string) string {
	return html.EscapeString(s)
}
