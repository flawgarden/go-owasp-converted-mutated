package controllers

import (
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01047 struct {
	web.Controller
}

func (c *BenchmarkTest01047) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01047) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01047) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if c.Ctx.Request.Header.Get("Referer") != "" {
		param = c.Ctx.Request.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<body>\n<p>Formatted like: %s and %s.\n</p>\n</body>\n</html>", obj[0], obj[1])

	c.Ctx.ResponseWriter.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz50709 := []rune(param)
		bar = string(append(sbxyz50709[:len(sbxyz50709)-1], 'Z'))
	}
	return bar
}

func main() {
	web.Router("/xss-01/BenchmarkTest01047", &BenchmarkTest01047{})
	web.Run()
}
