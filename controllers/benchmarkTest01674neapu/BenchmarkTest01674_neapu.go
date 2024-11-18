//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [676, 78, 703]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: [79, 78]
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest01674/BenchmarkTest01674.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_3_array_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
)

type BenchmarkTest01674 struct{}

func (b *BenchmarkTest01674) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery

arr4124 := []string{queryString}
nested7231 := NewNestedFields3FromArray(arr4124)
queryString = nested7231.nested1.nested1.nested1.values[0]

	paramval := "BenchmarkTest01674="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		for i := 0; i < len(queryString)-len(paramval); i++ {
			if queryString[i:i+len(paramval)] == paramval {
				paramLoc = i
				break
			}
		}
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01674"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	for i := paramLoc + len(paramval); i < len(queryString); i++ {
		if queryString[i] == '&' {
			ampersandLoc = i
			break
		}
	}
	if ampersandLoc != len(queryString) {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	a1 := "sh"
	a2 := "-c"
	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Problem executing cmdi", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	thing := createThing()
	bar := thing.doSomething(param)
	return bar
}

type ThingInterface interface {
	doSomething(string) string
}

func createThing() ThingInterface {
	return &thing{}
}

type thing struct{}

func (t *thing) doSomething(param string) string {
	return param
}

func main() {
	http.Handle("/cmdi-01/BenchmarkTest01674", &BenchmarkTest01674{})
	http.ListenAndServe(":8080", nil)
}
