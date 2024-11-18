package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02427 struct {
	web.Controller
}

func (c *BenchmarkTest02427) Get() {
	c.Post()
}

func (c *BenchmarkTest02427) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02427")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.Output.Header("Set-Cookie", fmt.Sprintf("SomeCookie=%s; HttpOnly; Path=%s", bar, c.Ctx.Request.RequestURI))
	c.Ctx.Output.Body([]byte(fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)))
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	web.Router("/securecookie-00/BenchmarkTest02427", &BenchmarkTest02427{})
	web.Run()
}
