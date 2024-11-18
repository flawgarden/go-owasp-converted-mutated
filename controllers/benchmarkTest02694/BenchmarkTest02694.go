package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02694Controller struct {
	web.Controller
}

func (c *BenchmarkTest02694Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02694Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	param := c.Ctx.Input.Query("BenchmarkTest02694")
	bar := doSomething(param)
	c.Ctx.Output.Body([]byte(bar))
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func main() {
	web.Router("/xss-05/BenchmarkTest02694", &BenchmarkTest02694Controller{})
	web.Run()
}
