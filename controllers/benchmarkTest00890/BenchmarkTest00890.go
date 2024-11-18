package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00890Controller struct {
	web.Controller
}

func (c *BenchmarkTest00890Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00890Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00890")
	bar := "safe!"
	map61765 := make(map[string]interface{})
	map61765["keyA-61765"] = "a-Value"
	map61765["keyB-61765"] = param
	map61765["keyC"] = "another-Value"
	bar = map61765["keyB-61765"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(bar))
}

func main() {
	web.Router("/xss-01/BenchmarkTest00890", &BenchmarkTest00890Controller{})
	web.Run()
}
