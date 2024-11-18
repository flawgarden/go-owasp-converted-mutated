package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00386Controller struct {
	web.Controller
}

func (c *BenchmarkTest00386Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00386Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00386Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00386")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map8057 := make(map[string]interface{})
	map8057["keyA-8057"] = "a_Value"
	map8057["keyB-8057"] = param
	map8057["keyC"] = "another_Value"
	bar = map8057["keyB-8057"].(string)
	bar = map8057["keyA-8057"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(c.Ctx.ResponseWriter, "Formatted like: %s and %s.", bar, "b")
}

func main() {
	web.Router("/xss-00/BenchmarkTest00386", &BenchmarkTest00386Controller{})
	web.Run()
}
