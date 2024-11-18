//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02134/BenchmarkTest02134.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_nested_string_initialization_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
)

type BenchmarkTest02134 struct{}

func (b *BenchmarkTest02134) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	param := r.URL.Query().Get("BenchmarkTest02134")

nsh := NewNestedStringHolder("")
param = nsh.GetValue()

	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02134", &BenchmarkTest02134{})
	http.ListenAndServe(":8080", nil)
}
