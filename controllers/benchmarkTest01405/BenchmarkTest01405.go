package controllers

import (
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01405Controller struct {
	web.Controller
}

func (c *BenchmarkTest01405Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01405Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01405" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := new(Test).doSomething(param)

	var fileName string
	var err error
	fileName = filepath.Join(os.Getenv("TESTFILES_DIR"), bar)

	fis, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open FileInputStream on file: '" + fileName + "'"))
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, err := fis.Read(b)
	if err == nil {
		c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n" + string(b[:size])))
	} else {
		c.Ctx.ResponseWriter.Write([]byte("Problem getting FileInputStream: " + err.Error()))
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return param
}
