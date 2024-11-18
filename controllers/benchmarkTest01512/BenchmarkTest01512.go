package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01512Controller struct {
	web.Controller
}

func (c *BenchmarkTest01512Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01512Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01512")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	_, _ = c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(bar, obj...)))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sb := []rune(param)
		sb[len(param)-1] = 'Z'
		bar = string(sb)
	}
	return bar
}
