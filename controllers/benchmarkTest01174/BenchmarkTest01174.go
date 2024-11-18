package controllers

import (
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	web.Router("/xss-02/BenchmarkTest01174", &BenchmarkTest01174{})
}

type BenchmarkTest01174 struct {
	web.Controller
}

func (c *BenchmarkTest01174) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01174) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01174) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := (&Test{}).doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
