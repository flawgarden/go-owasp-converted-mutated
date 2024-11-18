package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02227 struct {
	web.Controller
}

func (c *BenchmarkTest02227) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02227) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02227) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	params := c.Ctx.Input.Params()
	param := params["BenchmarkTest02227"]

	bar := doSomething(c.Ctx.Request, param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func doSomething(req *http.Request, param string) string {
	sb := fmt.Sprintf("%s_SafeStuff", param)
	return sb
}

func main() {
	web.Router("/xss-04/BenchmarkTest02227", &BenchmarkTest02227{})
	web.Run()
}
