package controllers

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00153Controller struct {
	web.Controller
}

func (c *BenchmarkTest00153Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00153Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if c.Ctx.Input.Header("Referer") != "" {
		param = c.Ctx.Input.Header("Referer")
	}

	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	map96050 := make(map[string]interface{})
	map96050["keyA-96050"] = "a-Value"
	map96050["keyB-96050"] = param
	map96050["keyC"] = "another-Value"
	bar = map96050["keyB-96050"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
