package controllers

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01926Controller struct {
	web.Controller
}

func (c *BenchmarkTest01926Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01926Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01926Controller) DoPost() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := c.Ctx.Request.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Data["json"] = bar
	c.ServeJSON()
}

func doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
