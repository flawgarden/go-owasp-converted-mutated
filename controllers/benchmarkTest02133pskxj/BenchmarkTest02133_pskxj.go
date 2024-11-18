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
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/switch.tmt with name switch_operator_any_negative 
//Used extensions: MACRO_Any_string -> ~[VAR_string]~
//Program:
package controllers

import (
	"net/http"
)

type BenchmarkTest02133Controller struct{}

func (c *BenchmarkTest02133Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02133")

tmpUnique42 := ""
switch param {
case "abQgi":
    param = ""
default:
    param = tmpUnique42
}

	if param == "" {
		param = ""
	}

	bar := doSomething(param)

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
