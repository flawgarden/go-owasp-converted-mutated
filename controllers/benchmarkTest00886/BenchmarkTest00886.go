package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00886 struct {
	web.Controller
}

func (c *BenchmarkTest00886) Get() {
	c.doPost()
}

func (c *BenchmarkTest00886) Post() {
	c.doPost()
}

func (c *BenchmarkTest00886) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00886")

	bar := "safe!"
	map8361 := make(map[string]interface{})
	map8361["keyA-8361"] = "a_Value"
	map8361["keyB-8361"] = param
	map8361["keyC"] = "another_Value"
	bar = map8361["keyB-8361"].(string)
	bar = map8361["keyA-8361"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{bar, "b"}
	_, _ = fmt.Fprintf(c.Ctx.ResponseWriter, "Formatted like: %s and %s.", obj[0], obj[1])
}

func main() {
	web.Router("/xss-01/BenchmarkTest00886", &BenchmarkTest00886{})
	web.Run()
}
