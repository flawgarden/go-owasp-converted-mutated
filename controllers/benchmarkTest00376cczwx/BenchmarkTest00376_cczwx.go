//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00376/BenchmarkTest00376.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_2_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00376 struct{}

func (b *BenchmarkTest00376) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00376")
	if param == "" {
		param = ""
	}

	bar := ""
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

nested7231 := NewNestedFields2("IzQpL")
bar = nested7231.nested1.nested1.value

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := fmt.Sprintf("Formatted like: %s and %s.", obj[0], obj[1])
	w.Write([]byte(output))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00376", &BenchmarkTest00376{})
	http.ListenAndServe(":8080", nil)
}
