package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00493 struct {
	web.Controller
}

func (c *BenchmarkTest00493) Get() {
	c.Post()
}

func (c *BenchmarkTest00493) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00493")
	bar := "safe!"
	map8943 := make(map[string]interface{})
	map8943["keyA-8943"] = "a_Value"
	map8943["keyB-8943"] = param
	map8943["keyC"] = "another_Value"

	bar = fmt.Sprintf("%v", map8943["keyB-8943"])
	bar = fmt.Sprintf("%v", map8943["keyA-8943"])

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte("Parameter value: " + bar))
}

func main() {
	web.Router("/xss-00/BenchmarkTest00493", &BenchmarkTest00493{})
	web.Run()
}
