package controllers

import (
	"os"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02667Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02667Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02667Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02667")
	bar := doSomething(param)

	fileName := filepath.Join("testfiles", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		c.Ctx.WriteString("Couldn't open FileInputStream on file: '" + fileName + "'")
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	c.Ctx.WriteString("The beginning of file: '" + fileName + "' is:\n\n")
	c.Ctx.WriteString(string(b[:size]))
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
