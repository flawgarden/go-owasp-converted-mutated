package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01051Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01051Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01051Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01051Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := c.Ctx.Request.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)
	bar := new(Test).DoSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	responseMsg := "Formatted like: %s and %s."
	output := []interface{}{bar, "b"}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(responseMsg, output...)))
}

type Test struct{}

func (t *Test) DoSomething(r *http.Request, param string) string {
	bar := ""

	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
