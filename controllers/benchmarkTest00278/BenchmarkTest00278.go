package controllers

import (
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00278Controller struct {
	web.Controller
}

func (c *BenchmarkTest00278Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00278Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	referer := c.Ctx.Request.Header.Get("Referer")
	if referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)
	bar := htmlEscape(param)

	c.Ctx.Output.Body([]byte(bar))
}

func htmlEscape(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}

func main() {
	web.Router("/xss-00/BenchmarkTest00278", &BenchmarkTest00278Controller{})
	web.Run()
}
