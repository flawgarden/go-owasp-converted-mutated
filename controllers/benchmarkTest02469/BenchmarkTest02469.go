package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02469Controller struct {
	web.Controller
}

func (c *BenchmarkTest02469Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest02469Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.Ctx.Input.Query("BenchmarkTest02469")
	var param string
	if values != "" {
		param = values
	} else {
		param = ""
	}

	bar := doSomething(param)

	var fileName string
	if bar != "" {
		fileName = fmt.Sprintf("some_directory/%s", bar) // Adjust path as needed
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
}

func doSomething(param string) string {
	var bar string
	if param != "" {
		bar = string(param) // This simulates some processing
	}
	return bar
}
