package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00730Controller struct {
	web.Controller
}

func (c *BenchmarkTest00730Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00730Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00730")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := "safe!"
	map29173 := make(map[string]interface{})
	map29173["keyA-29173"] = "a_Value"
	map29173["keyB-29173"] = param
	map29173["keyC"] = "another_Value"
	bar = map29173["keyB-29173"].(string)
	bar = map29173["keyA-29173"].(string)

	c.Ctx.Output.Body([]byte(bar))
}
