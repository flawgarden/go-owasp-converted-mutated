package controllers

import (
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01750Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01750Controller) Get() {
	c.post()
}

func (c *BenchmarkTest01750Controller) Post() {
	c.post()
}

func (c *BenchmarkTest01750Controller) post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01750")
	bar := new(Test).doSomething(param)

	fileName := "/path/to/test/files/" + bar
	fis, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open FileInputStream on file: '" + fileName + "'"))
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, err := fis.Read(b)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error reading file"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n"))
	c.Ctx.ResponseWriter.Write(b[:size])
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	// Here we simulate some processing on the param
	return strings.TrimSpace(param) // Potentially unsafe processing
}
