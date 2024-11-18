//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00149/BenchmarkTest00149.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name string_factory_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00149 struct{}

func (b *BenchmarkTest00149) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")

sf := CreateStringFactory()
param = sf.val

	}

	param, _ = url.QueryUnescape(param)

	sbxyz19132 := param + "_SafeStuff"
	bar := sbxyz19132

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	w.Write([]byte(output))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00149", &BenchmarkTest00149{})
	http.ListenAndServe(":8080", nil)
}
