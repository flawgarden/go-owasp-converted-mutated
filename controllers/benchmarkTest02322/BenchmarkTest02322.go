package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02322 struct {
	web.Controller
}

func (c *BenchmarkTest02322) Get() {
	c.Post()
}

func (c *BenchmarkTest02322) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02322" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output := struct {
		Bar string
		B   string
	}{
		Bar: bar,
		B:   "b",
	}
	json.NewEncoder(c.Ctx.ResponseWriter).Encode(output)
}

func doSomething(param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}
