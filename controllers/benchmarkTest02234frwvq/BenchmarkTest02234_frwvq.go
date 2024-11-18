//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [563]
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02234/BenchmarkTest02234.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/conditionswitch.tmt with name switch_multiple_case_conditions_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02234 struct{}

func (b *BenchmarkTest02234) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	if values, ok := r.Form["BenchmarkTest02234"]; ok && len(values) > 0 {
		param = values[0]
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")

value := 6
switch value {
case 1, 2, 3:
    bar = "fixed_string"
case 4, 5, 6:
    bar = bar + "_suffix"
    fallthrough
default:
    bar = "fixed_string"
}

	_, _ = fmt.Fprintf(w, bar)
}

func doSomething(param string) string {
	sbxyz86132 := param + "_SafeStuff"
	return sbxyz86132
}
