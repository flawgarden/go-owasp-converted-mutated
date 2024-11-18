package controllers

import (
	"fmt"
	"html"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02401Controller struct {
	web.Controller
}

func (c *BenchmarkTest02401Controller) Get() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02401")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	c.Ctx.Output.Header("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	c.Ctx.Output.Body([]byte(output))
}

func doSomething(param string) string {
	return html.EscapeString(param)
}

func main() {
	web.Router("/xss-04/BenchmarkTest02401", &BenchmarkTest02401Controller{})
	web.Run()
}
