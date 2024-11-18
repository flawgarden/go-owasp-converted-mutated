package controllers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00365Controller struct {
	web.Controller
}

func (c *BenchmarkTest00365Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00365Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00365")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map10106 := make(map[string]interface{})
	map10106["keyA-10106"] = "a_Value"
	map10106["keyB-10106"] = param
	map10106["keyC"] = "another_Value"
	bar = map10106["keyB-10106"].(string)
	bar = map10106["keyA-10106"].(string)

	fileName := filepath.Join("path/to/testfiles", bar)
	var is *os.File
	var err error

	is, err = os.Open(fileName)
	if err != nil {
		c.Ctx.WriteString("Couldn't open InputStream on file: '" + fileName + "'")
		return
	}
	defer is.Close()

	b := make([]byte, 1000)
	size, err := is.Read(b)
	if err != nil {
		c.Ctx.WriteString("Problem getting InputStream: " + err.Error())
		return
	}

	c.Ctx.WriteString("The beginning of file: '" + htmlEncode(fileName) + "' is:\n\n")
	c.Ctx.WriteString(htmlEncode(string(b[:size])))
}

func htmlEncode(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}
