package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01336Controller struct {
	web.Controller
}

func (c *BenchmarkTest01336Controller) Get() {
	c.post(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest01336Controller) Post() {
	c.post(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest01336Controller) post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01336")
	bar := c.doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func (c *BenchmarkTest01336Controller) doSomething(param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
