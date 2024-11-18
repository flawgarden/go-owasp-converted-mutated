package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00552Controller struct {
	web.Controller
}

func (c *BenchmarkTest00552Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00552Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Input.Params()
	for name := range names {
		values := c.GetStrings(name)
		if values != nil && flag {
			for _, value := range values {
				if value == "BenchmarkTest00552" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := "safe!"
	map63597 := make(map[string]interface{})
	map63597["keyA-63597"] = "a-Value"
	map63597["keyB-63597"] = param
	map63597["keyC"] = "another-Value"
	bar = map63597["keyB-63597"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}
