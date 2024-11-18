package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01343Controller struct {
	web.Controller
}

func (c *BenchmarkTest01343Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01343Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01343")

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, _ = c.Ctx.ResponseWriter.Write([]byte(strings.TrimSpace(bar)))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the last 'safe' value
	}

	return bar
}
