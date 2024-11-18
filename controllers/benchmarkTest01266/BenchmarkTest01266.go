package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01266Controller struct {
	web.Controller
}

func (c *BenchmarkTest01266Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01266Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01266Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01266")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		c.Ctx.ResponseWriter.Write([]byte(bar[:length]))
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}
