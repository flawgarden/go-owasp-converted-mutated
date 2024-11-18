package controllers

import (
	"encoding/json"
	"log"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00887 struct {
	web.Controller
}

func (c *BenchmarkTest00887) Get() {
	c.Post()
}

func (c *BenchmarkTest00887) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00887")

	bar := "safe!"
	map39726 := map[string]interface{}{
		"keyA-39726": "a-Value",
		"keyB-39726": param,
		"keyC":       "another-Value",
	}
	bar = map39726["keyB-39726"].(string)

	output, err := json.Marshal(bar)
	if err != nil {
		log.Fatal(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func main() {
	web.Router("/xss-01/BenchmarkTest00887", &BenchmarkTest00887{})
	web.Run()
}
