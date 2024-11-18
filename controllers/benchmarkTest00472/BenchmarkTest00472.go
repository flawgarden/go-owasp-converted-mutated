package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00472Controller struct {
	web.Controller
}

func (c *BenchmarkTest00472Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00472Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00472")
	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, err := fmt.Fprintf(c.Ctx.ResponseWriter, "Formatted like: %s and %s.", bar, "b")
	if err != nil {
		fmt.Println(err)
	}
}
