//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01645/BenchmarkTest01645.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_simple_negative 
//Used extensions: 
//Program:
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

nested7231 := NewNestedFields1("lgNQN")
queryString = nested7231.nested1.value

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
