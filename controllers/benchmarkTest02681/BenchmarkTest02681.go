package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02681Controller struct {
	web.Controller
}

func (c *BenchmarkTest02681Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02681Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest02681")
	bar := doSomething(param)

	obj := []interface{}{"a", "b"}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(bar, obj...)))
}

func doSomething(param string) string {
	bar := param
	return bar
}
