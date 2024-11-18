package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02130Controller struct {
	web.Controller
}

func (c *BenchmarkTest02130Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02130Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02130")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	fmt.Fprintf(c.Ctx.ResponseWriter, "Formatted like: %s and %s.", obj...)
}

func doSomething(param string) string {
	sbxyz60224 := param + "_SafeStuff"
	return sbxyz60224
}

func main() {
	web.Router("/xss-04/BenchmarkTest02130", &BenchmarkTest02130Controller{})
	web.Run()
}
