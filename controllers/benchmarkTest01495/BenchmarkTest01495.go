package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01495Controller struct {
	web.Controller
}

func (c *BenchmarkTest01495Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01495Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01495")
	if param == "" {
		param = ""
	}

	bar := new(Test).DoSomething(c.Ctx.Request, param)

	startURIslashes := ""
	if os.PathSeparator == '\\' {
		startURIslashes = "/"
	} else {
		startURIslashes = "//"
	}

	fileURI := filepath.Join(startURIslashes, "path/to/testfiles", bar)
	fileTarget := &os.File{}
	if _, err := os.Stat(fileURI); err == nil {
		fileTarget = &os.File{}
	}

	c.Data["json"] = map[string]string{
		"message": "Access to file: '" + fileTarget.Name() + "' created.",
		"exists":  "File already exists.",
	}
	c.ServeJSON()
}

type Test struct{}

func (t *Test) DoSomething(request *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
