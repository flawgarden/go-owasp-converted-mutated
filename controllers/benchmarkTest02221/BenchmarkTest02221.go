package controllers

import (
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02221Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02221Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02221Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02221")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := "<!DOCTYPE html>\n<html>\n<body>\n<p>" + formatOutput(obj) + "\n</p>\n</body>\n</html>"
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz71523 := strings.Builder{}
		sbxyz71523.WriteString(param)
		bar = sbxyz71523.String()
		bar = bar[:len(bar)-1] + "Z"
	}
	return bar
}

func formatOutput(obj []interface{}) string {
	return "Formatted like: " + obj[0].(string) + " and " + obj[1].(string) + "."
}
