//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [563]
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02054/BenchmarkTest02054.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_string_initialization_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02054 struct{}

func (bt *BenchmarkTest02054) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := r.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0]
	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := doSomething(r, decodedParam)

	w.Header().Set("X-XSS-Protection", "0")

sh := NewStringHolder()
bar = sh.value

	obj := []interface{}{"a", "b"}
	fmt.Fprintf(w, bar, obj...)
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = param
	}

	return bar
}
