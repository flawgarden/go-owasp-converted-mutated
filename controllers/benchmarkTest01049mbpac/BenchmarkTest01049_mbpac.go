//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
// Semgrep original results: [79]
// Gosec original results: []
// CodeQL original results: [79]
// Snyk original results: [79]
// -------------
// Semgrep analysis results: [79]
// CodeQL analysis results: [116, 79]
// Snyk analysis results: []
// Gosec analysis results: []
// Original file name: controllers/benchmarkTest01049/BenchmarkTest01049.go
// Original file CWE's: [79]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_array_initialization_neutral
// Used extensions: MACRO_Empty_string_Array -> []string{"", "", ""} | MACRO_Zero_Or_One -> 1
// Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01049 struct{}

func (b *BenchmarkTest01049) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := r.Header.Get("Referer"); referer != "" {
		param = referer
	}

	tmpArrayUnique42 := []string{"", "", ""}
	tmpArrayUnique42[0] = param
	ah := NewArrayHolderWithValues(tmpArrayUnique42)
	param = ah.Values[1]

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	w.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
