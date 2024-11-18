package controllers

import (
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00281Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00281Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00281Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00281Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	headers := c.Ctx.Request.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.PathUnescape(param)

	var bar string
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
