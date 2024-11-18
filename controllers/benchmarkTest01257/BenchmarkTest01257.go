package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01257Controller struct {
	web.Controller
}

func (c *BenchmarkTest01257Controller) Get() {
	c.post()
}

func (c *BenchmarkTest01257Controller) Post() {
	c.post()
}

func (c *BenchmarkTest01257Controller) post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest01257")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	c.Ctx.Output.JSON(bar, false, false)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	thing := createThing()
	bar := thing.doSomething(param)
	return bar
}

func createThing() ThingInterface {
	return &Thing{}
}

type ThingInterface interface {
	doSomething(param string) string
}

type Thing struct{}

func (t *Thing) doSomething(param string) string {
	return param // пример простого возвращения параметра
}
