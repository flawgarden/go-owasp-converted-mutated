package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02201Controller struct {
	web.Controller
}

func (c *BenchmarkTest02201Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02201Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02201")
	bar := doSomething(param)

	fileName := filepath.Join("TESTFILES_DIR", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Couldn't open FileInputStream on file: "+fileName, http.StatusInternalServerError)
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n" + string(b[:size])))
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1]

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

	return strings.TrimSpace(bar)
}
