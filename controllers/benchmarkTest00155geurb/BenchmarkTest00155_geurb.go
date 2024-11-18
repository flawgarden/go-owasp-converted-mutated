//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00155/BenchmarkTest00155.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_3_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00155 struct{}

func (b *BenchmarkTest00155) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referrer := r.Header.Get("Referer"); referrer != "" {
		param = referrer
	}

nested7231 := NewNestedFields3(param)
param = nested7231.nested1.nested1.nested1.value

	param, _ = url.QueryUnescape(param)

	thing := CreateThing()
	bar := thing.DoSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

type ThingInterface interface {
	DoSomething(input string) string
}

func CreateThing() ThingInterface {
	return &Thing{}
}

type Thing struct{}

func (t *Thing) DoSomething(input string) string {
	return fmt.Sprintf("Processed: %s", input)
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00155", &BenchmarkTest00155{})
	http.ListenAndServe(":8080", nil)
}
