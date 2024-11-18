package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01284Controller struct {
	web.Controller
}

func (c *BenchmarkTest01284Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01284Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01284Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest01284")
	if param == "" {
		param = "0"
	}

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Write([]byte("Parameter value: " + bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
