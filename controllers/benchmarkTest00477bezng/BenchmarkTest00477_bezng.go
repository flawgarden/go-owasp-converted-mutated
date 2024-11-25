//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00477/BenchmarkTest00477.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/staticField.tmt with name class_with_static_string_field_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest00477 struct{}

func (b *BenchmarkTest00477) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00477")
	bar := fmt.Sprintf("%s_SafeStuff", param)

	w.Header().Set("X-XSS-Protection", "0")

DEFAULT_VALUE = bar
sfh := NewStaticFieldHolder()
bar = sfh.value

	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00477", &BenchmarkTest00477{})
	http.ListenAndServe(":8080", nil)
}
