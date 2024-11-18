package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01292Controller struct {
	web.Controller
}

func (c *BenchmarkTest01292Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest01292Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01292")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	// SecureRandom simulation and cookie logic goes here
	// ...

	c.Ctx.Output.Body([]byte("User: " + bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]                                    // get the param value
	}
	return bar
}
