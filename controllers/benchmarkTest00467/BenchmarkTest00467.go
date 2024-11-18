package controllers

import (
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00467Controller struct {
	web.Controller
}

func (c *BenchmarkTest00467Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00467Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00467")
	bar := "safe!"
	map88136 := make(map[string]interface{})
	map88136["keyA-88136"] = "a-Value"
	map88136["keyB-88136"] = param
	map88136["keyC"] = "another-Value"
	bar = map88136["keyB-88136"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(strconv.Quote(bar)))
}
