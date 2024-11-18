package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00471Controller struct {
	web.Controller
}

func (c *BenchmarkTest00471Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00471Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00471Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00471")
	bar := htmlEscape(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, _ = c.Ctx.ResponseWriter.Write([]byte("Formatted like: a and " + bar + "."))
}

func htmlEscape(input string) string {
	return jsonEscape(input)
}

func jsonEscape(input string) string {
	result, _ := json.Marshal(input)
	return string(result[1 : len(result)-1])
}
