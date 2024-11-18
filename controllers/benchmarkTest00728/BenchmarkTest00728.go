package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00728 struct {
	web.Controller
}

func (c *BenchmarkTest00728) Get() {
	c.Post()
}

func (c *BenchmarkTest00728) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00728")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output, err := json.Marshal(bar)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
