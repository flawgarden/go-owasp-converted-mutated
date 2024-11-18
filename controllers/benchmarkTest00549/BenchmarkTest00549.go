package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00549Controller struct {
	web.Controller
}

func (c *BenchmarkTest00549Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00549Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()

	for name, values := range names {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest00549" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	output, err := json.Marshal(bar)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func main() {
	web.Router("/xss-01/BenchmarkTest00549", &BenchmarkTest00549Controller{})
	web.Run()
}
