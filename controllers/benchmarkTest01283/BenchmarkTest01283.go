package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01283Controller struct {
	web.Controller
}

func (c *BenchmarkTest01283Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01283Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01283")
	if param == "" {
		param = ""
	}


	var str string
	if param != "" {
		str = param
	} else {
		str = "No cookie value supplied"
	}

	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:     "SomeCookie",
		Value:    str,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
		Secure:   false,
	})

	c.Ctx.ResponseWriter.Write([]byte("Created cookie: 'SomeCookie': with value: '" + str + "' and secure flag set to: false"))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
