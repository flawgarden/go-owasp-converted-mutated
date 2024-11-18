//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00715/BenchmarkTest00715.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_2_array_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest00715 struct{}

func (b *BenchmarkTest00715) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("X-XSS-Protection", "0")

	values := r.URL.Query()["BenchmarkTest00715"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}

arr4124 := []string{bar}
nested7231 := NewNestedFields2FromArray(arr4124)
bar = nested7231.nested1.nested1.values[0]

	fmt.Fprint(w, bar)
}

func main() {
	http.Handle("/", &BenchmarkTest00715{})
	http.ListenAndServe(":8080", nil)
}
