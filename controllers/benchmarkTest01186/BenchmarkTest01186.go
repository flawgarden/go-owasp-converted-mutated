package controllers

import (
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01186Controller struct {
	web.Controller
}

func (c *BenchmarkTest01186Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01186Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest01186"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
	}

	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	c.Ctx.ResponseWriter.Write([]byte("Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: false"))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	return url.QueryEscape(param) // HTML escaping can be added if needed
}
