package controllers

import (
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02588 struct {
	web.Controller
}

func (c *BenchmarkTest02588) Get() {
	c.Post()
}

func (c *BenchmarkTest02588) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Input.URI()
	paramval := "BenchmarkTest02588="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.WriteString("getQueryString() couldn't find expected parameter 'BenchmarkTest02588' in query string.")
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString, "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	a1227 := param
	b1227 := strings.Builder{}
	b1227.WriteString(a1227)
	b1227.WriteString(" SafeStuff")
	b1227String := b1227.String()
	b1227String = strings.Replace(b1227String, b1227String[len(b1227String)-len("Chars"):], "Chars", 1)

	map1227 := make(map[string]interface{})
	map1227["key1227"] = b1227String
	c1227 := map1227["key1227"].(string)
	d1227 := c1227[:len(c1227)-1]

	e1227 := string([]byte(d1227))
	f1227 := strings.Split(e1227, " ")[0]

	// Assuming here you have a compatible thing interface in Go
	bar := doSomeReflection(f1227)

	return bar
}

// Dummy reflection function
func doSomeReflection(input string) string {
	return strings.ToUpper(input)
}

func main() {
	web.Router("/xss-05/BenchmarkTest02588", &BenchmarkTest02588{})
	web.Run()
}
