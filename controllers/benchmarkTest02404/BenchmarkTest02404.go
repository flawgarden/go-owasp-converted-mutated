package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02404Controller struct {
	web.Controller
}

func (c *BenchmarkTest02404Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02404Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	param := c.GetString("BenchmarkTest02404")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	responseData := map[string]string{"output": bar}
	output, err := json.Marshal(responseData)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
