package controllers

import (
	"html"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00474 struct {
	web.Controller
}

func (c *BenchmarkTest00474) Get() {
	c.Do(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest00474) Post() {
	c.Do(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest00474) Do(req *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := req.FormValue("BenchmarkTest00474")
	bar := htmlEscape(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func htmlEscape(input string) string {
	return html.EscapeString(input)
}
