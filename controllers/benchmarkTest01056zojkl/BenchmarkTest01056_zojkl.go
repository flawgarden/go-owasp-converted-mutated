//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01056/BenchmarkTest01056.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_array_initialization_positive 
//Used extensions: MACRO_Empty_string_Array -> []string{"", "", ""}
//Program:
package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01056 struct{}

func (b *BenchmarkTest01056) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")

tmpArrayUnique42 := []string{"", "", ""}
tmpArrayUnique42[0] = param
ah := NewArrayHolderWithValues(tmpArrayUnique42)
param = ah.Values[0]

	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	length := 1
	if bar != "" {
		length = len(bar)
		w.Write([]byte(bar)[:length])
	}
}

func (b *BenchmarkTest01056) doSomething(r *http.Request, param string) string {
	bar := ""

	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	return bar
}
