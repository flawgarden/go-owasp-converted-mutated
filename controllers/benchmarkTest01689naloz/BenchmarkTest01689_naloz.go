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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_with_generics_negative 
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

queryString = varargsWithGenerics("NAhAo", "uPzvg")

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

func getFirstString(lines ...string) string {
    return getStringWithIndex(0, lines...)
}

func getStringWithIndex(ind int, lines ...string) string {
    return lines[ind]
}

func getFirstStringFromArray(lines ...string) string {
    return lines[0]
}

func varargsWithGenerics[T any](elements ...T) T {
    return elements[0]
}

func combineStrings(strs ...string) string {
    return strings.Join(strs, ", ")
}


