package controllers

import (
	"fmt"
	"html"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest02049Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02049Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02049Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["Referer"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func doSomething(param string) string {
	bar := html.EscapeString(param)
	return bar
}
