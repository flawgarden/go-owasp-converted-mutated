package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02052Controller struct {
	web.Controller
}

func (c *BenchmarkTest02052Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02052Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Input.Header("Referer")
	if headers != "" {
		param = headers
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", bar, "b")
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(req *http.Request, param string) string {
	bar := "safe!"
	map20081 := make(map[string]interface{})
	map20081["keyA-20081"] = "a_Value"
	map20081["keyB-20081"] = param
	map20081["keyC"] = "another_Value"
	bar = map20081["keyB-20081"].(string)
	bar = map20081["keyA-20081"].(string)

	return bar
}
