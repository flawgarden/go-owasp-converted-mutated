package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01499Controller struct {
	web.Controller
}

func (c *BenchmarkTest01499Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01499Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01499Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01499")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = filepath.Join("testfiles", bar)

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", filepath.Base(fileName))))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
