package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01774Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01774Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01774Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01774Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest01774")

	bar := new(Test).doSomething(param)

	c.Ctx.ResponseWriter.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map58318 := make(map[string]interface{})
	map58318["keyA-58318"] = "a_Value"
	map58318["keyB-58318"] = param
	map58318["keyC"] = "another_Value"
	bar = map58318["keyB-58318"].(string)
	bar = map58318["keyA-58318"].(string)

	return bar
}
