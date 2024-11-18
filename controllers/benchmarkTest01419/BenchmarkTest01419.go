package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01419Controller struct {
	web.Controller
}

func (c *BenchmarkTest01419Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01419Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()

	for name, values := range names {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest01419" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := new(Test).doSomething(param)

	output := []interface{}{"a", bar}
	c.Ctx.ResponseWriter.Write([]byte("Formatted like: " + output[0].(string) + " and " + output[1].(string) + "."))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
