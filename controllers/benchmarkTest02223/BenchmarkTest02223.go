package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02223Controller struct {
	web.Controller
}

func (c *BenchmarkTest02223Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02223Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := c.GetString("BenchmarkTest02223")
	bar := doSomething(id)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	sbxyz60073 := param + "_SafeStuff"
	return sbxyz60073
}
