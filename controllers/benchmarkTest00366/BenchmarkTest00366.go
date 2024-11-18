package controllers

import (
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00366Controller struct {
	web.Controller
}

func (c *BenchmarkTest00366Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00366Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00366")
	if param == "" {
		param = ""
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	fileName := filepath.Join("testfiles", bar)
	var is *os.File

	defer func() {
		if is != nil {
			is.Close()
		}
	}()

	is, err := os.Open(fileName)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Couldn't open InputStream on file: '" + fileName + "'"))
		return
	}
	defer is.Close()

	b := make([]byte, 1000)
	size, err := is.Read(b)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Problem getting InputStream: " + err.Error()))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + bar + "' is:\n\n"))
	c.Ctx.ResponseWriter.Write(b[:size])
}
