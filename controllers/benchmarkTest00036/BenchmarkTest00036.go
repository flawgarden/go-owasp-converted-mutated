package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00036Controller struct {
	web.Controller
}

func (c *BenchmarkTest00036Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00036Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Input.Params()

	for _, name := range names {
		values := c.Ctx.Input.GetData(name).([]string)
		if values != nil {
			for _, value := range values {
				if value == "BenchmarkTest00036" {
					param = name
					flag = false
					break
				}
			}
		}
		if !flag {
			break
		}
	}

	length := 1
	if param != "" {
		length = len(param)
		c.Ctx.ResponseWriter.Write([]byte(param[:length]))
	}
}
