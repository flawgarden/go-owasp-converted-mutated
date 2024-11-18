package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02111Controller struct {
	web.Controller
}

func (c *BenchmarkTest02111Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02111Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02111", "")
	bar := doSomething(param)

	fileName := "/path/to/testfiles/" + bar

	fos, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Couldn't open FileOutputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}
	defer fos.Close()

	c.Ctx.ResponseWriter.Write([]byte("Now ready to write to file: " + htmlEncode(fileName)))
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

func htmlEncode(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}
