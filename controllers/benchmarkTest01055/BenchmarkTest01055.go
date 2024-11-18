package controllers

import (
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01055Controller struct {
	web.Controller
}

func (c *BenchmarkTest01055Controller) Get() {
	c.post()
}

func (c *BenchmarkTest01055Controller) Post() {
	c.post()
}

func (c *BenchmarkTest01055Controller) post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := ""

	if referer := c.Ctx.Request.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)
	bar := new(Test).doSomething(c.Ctx.Request, param)
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
