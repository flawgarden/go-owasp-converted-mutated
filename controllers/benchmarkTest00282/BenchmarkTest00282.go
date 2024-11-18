package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00282Controller struct {
	web.Controller
}

func (c *BenchmarkTest00282Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00282Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0] // just grab first element
	}

	param, _ = url.QueryUnescape(param)

	bar := escapeForHTML(param)

	c.Ctx.Output.Header("X-XSS-Protection", "0")
	output := "Formatted like: " + bar + " and " + "b."
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Ctx.Output.Body([]byte(output))
}

func escapeForHTML(input string) string {
	return strings.ReplaceAll(input, "&", "&amp;") +
		strings.ReplaceAll(input, "<", "&lt;") +
		strings.ReplaceAll(input, ">", "&gt;")
}
