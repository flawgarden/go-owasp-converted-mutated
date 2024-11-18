package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00030Controller struct {
	web.Controller
}

func (c *BenchmarkTest00030Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00030Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest00030")
	obj := []interface{}{"a", "b"}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(param, obj...)))
}
