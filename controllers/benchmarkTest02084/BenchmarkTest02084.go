package controllers

import (
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02084Controller struct {
	web.Controller
}

func (c *BenchmarkTest02084Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02084Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02084Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02084"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param = strings.ReplaceAll(param, "%20", " ")

	bar := doSomething(c.Ctx.Request, param)

	c.Ctx.Request.RequestURI = c.Ctx.Request.RequestURI // simulate setting session attribute
	c.Ctx.ResponseWriter.Write([]byte("Item: 'userid' with value: '" + bar + "' saved in session."))
}

func doSomething(req *http.Request, param string) string {
	// Placeholder for actual implementation
	return param
}

func main() {
	web.Router("/trustbound-01/BenchmarkTest02084", &BenchmarkTest02084Controller{})
	web.Run()
}
