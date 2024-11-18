package controllers

import (
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00284Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00284Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00284Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00284Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := c.Ctx.Request.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0] // just grab the first element
	}

	param, _ = url.QueryUnescape(param)

	bar := param

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
