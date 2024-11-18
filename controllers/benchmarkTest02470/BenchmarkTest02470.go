package controllers

import (
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02470Controller struct {
	web.Controller
}

func (c *BenchmarkTest02470Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02470Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02470")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = filepath.Join("path/to/testfiles", bar)

	var err error
	fos, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open FileOutputStream on file: '" + fileName + "'"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Now ready to write to file: " + fileName))
}

func doSomething(param string) string {
	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
