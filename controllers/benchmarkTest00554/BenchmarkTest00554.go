package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00554 struct {
	web.Controller
}

func (b *BenchmarkTest00554) Get() {
	b.Post()
}

func (b *BenchmarkTest00554) Post() {
	b.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := b.Ctx.Input.Params()

	for name := range names {
		values := b.Ctx.Request.URL.Query()[name]
		if values != nil {
			for _, value := range values {
				if value == "BenchmarkTest00554" {
					param = name
					flag = false
					break
				}
			}
		}
		if !flag {
			break
		}
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	b.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	json.NewEncoder(b.Ctx.ResponseWriter).Encode(bar)
}
