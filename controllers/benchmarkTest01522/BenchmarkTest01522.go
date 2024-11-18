package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01522Controller struct {
	web.Controller
}

func (c *BenchmarkTest01522Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01522Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01522")
	if param == "" {
		param = "No cookie value supplied"
	}

	bar := new(Test).doSomething(param)

	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     c.Ctx.Request.URL.Path,
	})

	response := struct {
		Message string `json:"message"`
		Cookie  string `json:"cookie"`
	}{
		Message: "Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: true",
		Cookie:  bar,
	}

	output, err := json.Marshal(response)
	if err == nil {
		c.Ctx.ResponseWriter.Write(output)
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return param // здесь можно добавить код для экранирования
}
