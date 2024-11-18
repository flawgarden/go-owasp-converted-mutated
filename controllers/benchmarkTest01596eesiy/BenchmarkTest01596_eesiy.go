// Semgrep original results: [79]
// Gosec original results: []
// CodeQL original results: [79]
// Snyk original results: []
// -------------
// Semgrep analysis results: [79, 319]
// CodeQL analysis results: []
// Gosec analysis results: []
// Snyk analysis results: []
// Original file name: controllers/benchmarkTest01596/BenchmarkTest01596.go
// Original file CWE's: [79]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/conditionswitch.tmt with name switch_multiple_conditions_positive
// Used extensions:
// Program:
package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest01596 struct{}

func (bt *BenchmarkTest01596) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest01596"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := bt.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")

	value := 7
	switch {
	case value < 0:
		bar = "fixed_string"
	case value >= 0 && value <= 10:
		bar = bar + "_suffix"
	case value > 10:
		bar = "fixed_string"
	default:
		bar = "fixed_string"
	}

	if bar != "" {
		w.Write([]byte(bar))
	}
}

func (bt *BenchmarkTest01596) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01596", &BenchmarkTest01596{})
	http.ListenAndServe(":8080", nil)
}
