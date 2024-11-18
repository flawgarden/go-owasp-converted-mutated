package controllers

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00156 struct {
	web.Controller
}

func (c *BenchmarkTest00156) Get() {
	c.Post()
}

func (c *BenchmarkTest00156) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("Referer")
	if param == "" {
		param = ""
	}

	// URL Decode the header value
	param, _ = url.QueryUnescape(param)

	var bar string
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func main() {
	web.Router("/xss-00/BenchmarkTest00156", &BenchmarkTest00156{})
	web.Run()
}
