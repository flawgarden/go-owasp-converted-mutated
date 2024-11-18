package controllers

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02398Controller struct {
	web.Controller
}

func (c *BenchmarkTest02398Controller) Get() {
	param := c.GetString("BenchmarkTest02398")
	if param == "" {
		param = ""
	}
	bar := doSomething(param)
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(input string) string {
	// Custom HTML escape logic here
	return input // replace with actual escaping
}

func main() {
	web.Router("/xss-04/BenchmarkTest02398", &BenchmarkTest02398Controller{})
	web.Run()
}
