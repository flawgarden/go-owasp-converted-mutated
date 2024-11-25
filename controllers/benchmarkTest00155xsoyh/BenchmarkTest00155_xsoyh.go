//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00155/BenchmarkTest00155.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name closure_returning_closure_positive 
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

generator := makeMessageGenerator(param)()
param = generator("_suffix")

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

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


