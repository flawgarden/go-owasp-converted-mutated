package controllers

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02688Controller struct {
	web.Controller
}

func (c *BenchmarkTest02688Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02688Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02688Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Query("BenchmarkTest02688")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf(bar, "a", "b")
	c.Ctx.WriteString(output)
}

func doSomething(param string) string {
	return strings.ReplaceAll(param, "<", "&lt;")
}

func main() {
	web.Router("/xss-05/BenchmarkTest02688", &BenchmarkTest02688Controller{})
	web.Run()
}
