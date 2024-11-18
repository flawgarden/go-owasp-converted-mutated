package controllers

import (
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00382Controller struct {
	web.Controller
}

func (c *BenchmarkTest00382Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00382Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00382")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	c.Ctx.Output.Body([]byte(strings.ReplaceAll(bar, "%s", "a, b")))
}

func main() {
	web.Router("/xss-00/BenchmarkTest00382", &BenchmarkTest00382Controller{})
	web.Run()
}
