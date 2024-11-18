package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest01597Controller struct {
	web.Controller
}

func (c *BenchmarkTest01597Controller) Get() {
	c.HandleRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest01597Controller) Post() {
	c.HandleRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest01597Controller) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.Form["BenchmarkTest01597"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := new(Test).doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	length := len(bar)
	if length > 0 {
		w.Write([]byte(bar))
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	sbxyz81751 := param + "_SafeStuff"
	return sbxyz81751
}
