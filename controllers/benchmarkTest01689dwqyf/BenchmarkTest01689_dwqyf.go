//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01689/BenchmarkTest01689.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/pointers.tmt with name pointer_init2_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

type BenchmarkTest01689Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest01689Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01689Controller) Post() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Request.URL.RawQuery

var ptr123 *string = new(string)
var ptr234 *string = new(string)
*ptr234 = queryString
ptr123 = ptr234
queryString = *ptr123

	paramval := "BenchmarkTest01689="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		c.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01689")))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Request, param)

	cmd := ""
	osName := "Linux" // Replace with actual OS check if needed

	if strings.Contains(osName, "Windows") {
		cmd = "cmd /C echo "
	} else {
		cmd = "echo "
	}

	argsEnv := []string{"Foo=bar"}
	r := exec.Command(cmd+bar, argsEnv...)
	out, err := r.CombinedOutput()
	if err != nil {
		c.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase"))
		return
	}

	c.ResponseWriter.Write(out)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}

func addSuffix(s *string, suf string) {
	*s = *s + suf
}

func addSuffixDoublePointer(s **string, suf *string) {
	**s = **s + *suf
}

func addSuffixDoublePointerBroken(s **string, suf *string) {
	*s = new(string)
	**s = **s + *suf
}

func swapStrings(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func removeSpaces(s *string) {
    *s = strings.ReplaceAll(*s, " ", "")
}


