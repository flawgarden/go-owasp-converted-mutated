package controllers

import (
	"fmt"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00013Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00013Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00013Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00013Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	referer := c.Ctx.Request.Header.Get("Referer")
	if referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	fmt.Fprintf(c.Ctx.ResponseWriter, param, obj...)
}
