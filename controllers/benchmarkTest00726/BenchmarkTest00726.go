package controllers

import (
	"html"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00726 struct {
	web.Controller
}

func (c *BenchmarkTest00726) Get() {
	c.Post()
}

func (c *BenchmarkTest00726) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00726")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := html.EscapeString(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func main() {
	web.Router("/xss-01/BenchmarkTest00726", &BenchmarkTest00726{})
	web.Run()
}
