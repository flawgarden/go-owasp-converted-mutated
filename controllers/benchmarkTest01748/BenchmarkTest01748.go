package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01748Controller struct {
	web.Controller
}

func (c *BenchmarkTest01748Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01748Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01748")
	bar := doSomething(param)

	fileName := filepath.Join("testfiles", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(c.Ctx.ResponseWriter, "Problem getting FileInputStream: %s", err.Error())
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, err := fis.Read(b)
	if err != nil {
		fmt.Fprintf(c.Ctx.ResponseWriter, "Couldn't read file: %s", err.Error())
		return
	}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("The beginning of file: '%s' is:\n\n%s", fileName, string(b[:size]))))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}
