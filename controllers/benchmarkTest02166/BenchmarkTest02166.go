package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02166Controller struct {
	web.Controller
}

func (c *BenchmarkTest02166Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02166Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02166Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02166")
	if param == "" {
		param = ""
	}

	bar := doSomething(c.Ctx.Request, param)

	c.Ctx.Request.Header.Set("userid", bar)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", htmlEscape(bar))))
}

func doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "<", "&lt;"), ">", "&gt;")
}

func main() {
	web.Router("/trustbound-01/BenchmarkTest02166", &BenchmarkTest02166Controller{})
	web.Run()
}
