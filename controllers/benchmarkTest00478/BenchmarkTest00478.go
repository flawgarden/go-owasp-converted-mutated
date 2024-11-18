package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00478Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00478Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00478Controller) Post() {
	c.Ctx.Output.Header("X-XSS-Protection", "0")
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00478")
	bar := ""

	if param != "" {
		decodedBytes, _ := decodeBase64(param)
		bar = string(decodedBytes)
	}

	c.Ctx.Output.Body([]byte(bar))
}

func decodeBase64(data string) ([]byte, error) {
	return json.Marshal(string(data))
}
