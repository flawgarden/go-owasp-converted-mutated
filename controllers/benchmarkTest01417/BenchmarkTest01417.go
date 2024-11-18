package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01417Controller struct {
	web.Controller
}

func (c *BenchmarkTest01417Controller) Get() {
	c.doPost(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest01417Controller) Post() {
	c.doPost(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest01417Controller) doPost(req *http.Request, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := req.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01417" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := new(Test).doSomething(req, param)

	res.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(res, "<!DOCTYPE html>\n<html>\n<body>\n<p>")
	_, _ = fmt.Fprintf(res, "Formatted like: %s and %s.", "a", bar)
	_, _ = fmt.Fprintf(res, "\n</p>\n</body>\n</html>")
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	bar := param
	return bar
}
