package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01525Controller struct {
	web.Controller
}

func (c *BenchmarkTest01525Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01525Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01525")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte("Parameter value: " + bar))
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz67457 := []rune(param)
		bar = string(append(sbxyz67457[:len(sbxyz67457)-1], 'Z'))
	}
	return bar
}
