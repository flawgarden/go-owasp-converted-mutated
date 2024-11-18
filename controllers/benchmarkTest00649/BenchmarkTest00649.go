package controllers

import (
	"html"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00649 struct {
	web.Controller
}

func (b *BenchmarkTest00649) Get() {
	b.Post()
}

func (b *BenchmarkTest00649) Post() {
	b.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	b.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := b.GetString("BenchmarkTest00649")
	if param == "" {
		param = ""
	}

	bar := html.EscapeString(param)
	b.Ctx.ResponseWriter.Write([]byte(bar))
}

func main() {
	web.Router("/xss-01/BenchmarkTest00649", &BenchmarkTest00649{})
	web.Run()
}
