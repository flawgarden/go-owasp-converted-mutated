package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02206Controller struct {
	web.Controller
}

func (c *BenchmarkTest02206Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02206Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := c.GetString("BenchmarkTest02206")
	bar := doSomething(param)

	fileName := filepath.Join("testfiles", bar)
	is, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Couldn't open InputStream on file: '" + fileName + "'")
		c.Ctx.ResponseWriter.Write([]byte("Problem getting InputStream: " + err.Error()))
		return
	}
	defer is.Close()

	b := make([]byte, 1000)
	size, _ := is.Read(b)

	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + htmlEscape(fileName) + "' is:\n\n"))
	c.Ctx.ResponseWriter.Write([]byte(htmlEscape(string(b[:size]))))
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

func htmlEscape(s string) string {
	return fmt.Sprintf("%q", s) // Simple escape representation
}
