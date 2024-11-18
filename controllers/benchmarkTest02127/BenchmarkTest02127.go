package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02127Controller struct {
	web.Controller
}

func (c *BenchmarkTest02127Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02127Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02127")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func main() {
	web.Router("/xss-04/BenchmarkTest02127", &BenchmarkTest02127Controller{})
	web.Run()
}
