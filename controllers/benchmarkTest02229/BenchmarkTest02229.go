package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02229 struct {
	web.Controller
}

func (c *BenchmarkTest02229) Get() {
	c.Post()
}

func (c *BenchmarkTest02229) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02229", "")

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	bar := "safe!"
	map26903 := map[string]interface{}{
		"keyA-26903": "a_Value",
		"keyB-26903": param,
		"keyC":       "another_Value",
	}
	bar = map26903["keyB-26903"].(string)
	bar = map26903["keyA-26903"].(string)

	return bar
}
