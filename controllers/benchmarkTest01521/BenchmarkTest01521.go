package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest01521 struct {
	web.Controller
}

func (c *BenchmarkTest01521) Get() {
	c.Post()
}

func (c *BenchmarkTest01521) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01521")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     c.Ctx.Request.RequestURI,
	}
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)

	output := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: %t", bar, false)
	c.Ctx.WriteString(output)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}
