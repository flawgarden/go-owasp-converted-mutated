package controllers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01160Controller struct {
	web.Controller
}

func (c *BenchmarkTest01160Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01160Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if headers := c.Ctx.Input.Header("BenchmarkTest01160"); headers != "" {
		param = headers
	}

	param = strings.ReplaceAll(param, "+", " ")

	bar := new(Test).doSomething(param)

	fileName := ""

	fileName = filepath.Join("path/to/test/files", bar)

	fos, err := os.Create(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open FileOutputStream on file: '" + fileName + "'"))
		return
	}
	defer fos.Close()

	c.Ctx.ResponseWriter.Write([]byte("Now ready to write to file: " + htmlEscape(fileName)))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
