package controllers

import (
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01232Controller struct {
	web.Controller
}

func (c *BenchmarkTest01232Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01232Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01232")
	if param == "" {
		param = ""
	}

	bar := c.doSomething(param)

	fileTarget := "/path/to/testfiles/" + bar
	c.Ctx.ResponseWriter.Write([]byte("Access to file: '" + fileTarget + "' created."))
	if fileExists(fileTarget) {
		c.Ctx.ResponseWriter.Write([]byte(" And file already exists."))
	} else {
		c.Ctx.ResponseWriter.Write([]byte(" But file doesn't exist yet."))
	}
}

func (c *BenchmarkTest01232Controller) doSomething(param string) string {
	// Simulate some processing with the parameter
	return param // Replace with actual logic if needed
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
