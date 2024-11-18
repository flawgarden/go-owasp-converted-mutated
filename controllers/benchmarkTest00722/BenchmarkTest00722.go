package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00722Controller struct {
	web.Controller
}

func (c *BenchmarkTest00722Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00722Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00722")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := "safe!"
	map49442 := make(map[string]interface{})
	map49442["keyA-49442"] = "a_Value"
	map49442["keyB-49442"] = param
	map49442["keyC"] = "another_Value"
	bar = map49442["keyB-49442"].(string)
	bar = map49442["keyA-49442"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(c.Ctx.ResponseWriter, bar, obj)
}
