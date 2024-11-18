package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01340Controller struct {
	web.Controller
}

func (c *BenchmarkTest01340Controller) Get() {
	cdoGet(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest01340Controller) Post() {
	cdoPost(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func cdoGet(request *http.Request, response http.ResponseWriter) {
	cdoPost(request, response)
}

func cdoPost(request *http.Request, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := request.URL.Query().Get("BenchmarkTest01340")
	bar := new(Test).doSomething(request, param)
	response.Header().Set("X-XSS-Protection", "0")
	response.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := "safe!"
	map36950 := make(map[string]interface{})
	map36950["keyA-36950"] = "a_Value"
	map36950["keyB-36950"] = param
	map36950["keyC"] = "another_Value"
	bar = map36950["keyB-36950"].(string)
	bar = map36950["keyA-36950"].(string)
	return bar
}
