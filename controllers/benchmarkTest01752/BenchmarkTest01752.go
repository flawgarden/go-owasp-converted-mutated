package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01752Controller struct {
	web.Controller
}

func (c *BenchmarkTest01752Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01752Controller) Post() {
	c.Ctx.Output.Context.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01752")

	bar := new(Test).doSomething(param)

	fileName := "/path/to/testfiles/" + bar // Укажите правильный путь

	output, err := json.Marshal(map[string]string{"message": "Now ready to write to file: " + fileName})
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.Output.Context.ResponseWriter.Write(output)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map11351 := make(map[string]interface{})
	map11351["keyA-11351"] = "a-Value"
	map11351["keyB-11351"] = param
	map11351["keyC"] = "another-Value"
	bar = map11351["keyB-11351"].(string)

	return bar
}
