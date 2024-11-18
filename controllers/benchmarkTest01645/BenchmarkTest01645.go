package controllers

import (
	"net/http"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01645Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01645Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01645Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01645Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramVal := "BenchmarkTest01645="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramVal)
		if queryString[len(queryString)-len(paramVal):] != paramVal {
			c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01645' in query string."))
			return
		}
	}

	param := queryString[paramLoc+len(paramVal):]

	bar := new(Test).doSomething(c.Ctx.Request, param)

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
		c.Ctx.ResponseWriter.Write([]byte("Error reading file."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("The beginning of file: '" + sanitize(fileName) + "' is:\n\n"))
	c.Ctx.ResponseWriter.Write(b[:size])
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	var bar string
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

func sanitize(input string) string {
	// Implement proper sanitization here (e.g., using html/template)
	return input
}
