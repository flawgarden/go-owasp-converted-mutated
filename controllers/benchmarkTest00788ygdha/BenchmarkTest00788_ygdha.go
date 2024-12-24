package controllers

import (
	"net/url"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00788Controller struct {
	web.Controller
}

func (c *BenchmarkTest00788Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00788Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Input.URI()
	paramval := "BenchmarkTest00788="
	paramLoc := -1

d := NewFakeDuckWithAttribute(queryString)
queryString = MakeItQuackFieldAttr(d, "tmp_str")

	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if queryString[paramLoc:] != paramval {
			c.Ctx.Output.Body([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest00788' in query string."))
			return
		}
	}

	param := queryString[paramLoc+len(paramval):]
	decodedParam, err := url.QueryUnescape(param)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error decoding parameter."))
		return
	}

	fileName := decodedParam
	fos, err := os.Create(fileName)
	if err != nil {
		c.Ctx.Output.Body([]byte("Couldn't open FileOutputStream on file: '" + fileName + "'"))
		return
	}
	defer fos.Close()

	c.Ctx.Output.Body([]byte("Now ready to write to file: " + fileName))
}

func MakeItQuack(duck interface{ Quack(string) string }, arg string) string {
    return duck.Quack(arg)
}

func MakeItQuackAttr(duck interface{}, arg string) string {
    if d, ok := duck.(interface{ Quack(string) string }); ok {
        return d.Quack(arg)
    }
    return "fixed_string"
}

func MakeItQuackFieldAttr(duck interface{}, arg string) string {
	if d, ok := duck.(DuckWithAttribute); ok && d.constant == 42 {
		return d.Quack(arg)
	}
	return "fixed_string"
}


