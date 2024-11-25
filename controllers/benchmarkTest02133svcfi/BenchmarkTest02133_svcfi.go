//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02133/BenchmarkTest02133.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/assertions.tmt with name type_assertion_with_struct_pointer_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
)

type BenchmarkTest02133Controller struct{}

func (c *BenchmarkTest02133Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02133")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

var i123 interface{} = &Anon{Value1: bar}
if ptr, ok := i123.(*EmbeddedStruct); ok {
     bar = ptr.Field1
} else {
    bar = "IDEOR"
}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}
