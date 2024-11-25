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
//Original file name: controllers/benchmarkTest01922/BenchmarkTest01922.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_array_initialization_negative 
//Used extensions: MACRO_Empty_string_Array -> []string{"", "", ""}
//Program:
package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01922 struct{}

func (b *BenchmarkTest01922) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if referrer := r.Header.Get("Referer"); referrer != "" {
		param = referrer
	}

	decodedParam, _ := url.QueryUnescape(param)

tmpArrayUnique42 := []string{"", "", ""}
tmpArrayUnique42[0] = decodedParam
ah := NewArrayHolderWithValues(tmpArrayUnique42)
decodedParam = ah.Values[1]

	bar := doSomething(decodedParam)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
