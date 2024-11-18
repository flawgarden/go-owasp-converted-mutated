package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01745Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01745Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01745Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01745Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01745")
	bar := new(Test).doSomething(param)

	fileTarget := filepath.Join(bar, "Test.txt")
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))

	if _, err := os.Stat(fileTarget); err == nil {
		c.Ctx.ResponseWriter.Write([]byte(" And file already exists."))
	} else {
		c.Ctx.ResponseWriter.Write([]byte(" But file doesn't exist yet."))
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return param
}
