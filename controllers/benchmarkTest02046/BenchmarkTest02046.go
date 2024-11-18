package controllers

import (
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02046Controller struct {
	web.Controller
}

func (c *BenchmarkTest02046Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02046Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(request *http.Request, param string) string {
	bar := "safe!"
	map81510 := make(map[string]interface{})
	map81510["keyA-81510"] = "a-Value"
	map81510["keyB-81510"] = param
	map81510["keyC"] = "another-Value"
	bar = map81510["keyB-81510"].(string)

	return bar
}
