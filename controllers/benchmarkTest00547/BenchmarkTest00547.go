package controllers

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00547Controller struct {
	web.Controller
}

func (c *BenchmarkTest00547Controller) Get() {
	param := ""
	flag := true
	names := c.Ctx.Input.Param(":param")
	if names != "" {
		for _, name := range strings.Split(names, ",") {
			if !flag {
				break
			}
			values := c.Ctx.Request.URL.Query()[name]
			for _, value := range values {
				if value == "BenchmarkTest00547" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("%s", bar)
	c.Ctx.WriteString(output)
}
