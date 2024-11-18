package controllers

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00279 struct {
	web.Controller
}

func (c *BenchmarkTest00279) Get() {
	c.Post()
}

func (c *BenchmarkTest00279) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Input.Header("Referer")
	if headers != "" {
		param = headers
	}

	param, _ = url.QueryUnescape(param)

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func main() {
	web.Router("/xss-00/BenchmarkTest00279", &BenchmarkTest00279{})
	web.Run()
}
