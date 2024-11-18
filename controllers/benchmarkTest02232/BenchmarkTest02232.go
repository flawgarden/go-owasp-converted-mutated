package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02232Controller struct {
	web.Controller
}

func (c *BenchmarkTest02232Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02232Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02232")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(bar, obj...)))
}

func doSomething(param string) string {
	return param
}
