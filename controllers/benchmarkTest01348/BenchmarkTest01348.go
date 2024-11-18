package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01348Controller struct {
	web.Controller
}

func (c *BenchmarkTest01348Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01348Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01348")

	bar := c.doSomething(param)

	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func (c *BenchmarkTest01348Controller) doSomething(param string) string {
	bar := htmlEscape(param)
	return bar
}

func htmlEscape(input string) string {
	return jsonEscape(input)
}

func jsonEscape(input string) string {
	escaped, _ := json.Marshal(input)
	return string(escaped)
}
