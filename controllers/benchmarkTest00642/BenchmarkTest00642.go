package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00642 struct {
	web.Controller
}

func (c *BenchmarkTest00642) Get() {
	c.Post()
}

func (c *BenchmarkTest00642) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00642")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
