package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02492Controller struct {
	web.Controller
}

func (c *BenchmarkTest02492Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02492Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02492")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

	c.Ctx.Output.Body([]byte(bar))
}

func doSomething(param string) string {
	bar := encodeForHTML(param)
	return bar
}

func encodeForHTML(param string) string {
	// Замените следующую строку своей реализацией экранирования
	return jsonEscape(param)
}

func jsonEscape(s string) string {
	bytes, _ := json.Marshal(s)
	return string(bytes[1 : len(bytes)-1]) // убираем кавычки
}
