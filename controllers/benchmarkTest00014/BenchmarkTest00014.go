package controllers

import (
	"fmt"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00014Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00014Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00014Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	referer := c.Ctx.Request.Header.Get("Referer")
	if referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	_, _ = fmt.Fprintf(c.Ctx.ResponseWriter, param, obj...)
}
