//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: []
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: []
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02228/BenchmarkTest02228.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/duck.tmt with name duck_typing_multiple_attributes_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest02228 struct{}

func (b *BenchmarkTest02228) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02228")
	bar := doSomething(param)

d := NewFakeDuckWithAttribute(bar)
bar = MakeItQuackFieldAttr(d, "tmp_str")

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(param string) string {
	var sb strings.Builder
	sb.WriteString(param)
	bar := sb.String() + "_SafeStuff"
	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02228", &BenchmarkTest02228{})
	http.ListenAndServe(":8080", nil)
}

func MakeItQuack(duck interface{ Quack(string) string }, arg string) string {
    return duck.Quack(arg)
}

func MakeItQuackAttr(duck interface{}, arg string) string {
    if d, ok := duck.(interface{ Quack(string) string }); ok {
        return d.Quack(arg)
    }
    return "fixed_string"
}

func MakeItQuackFieldAttr(duck interface{}, arg string) string {
	if d, ok := duck.(DuckWithAttribute); ok && d.constant == 42 {
		return d.Quack(arg)
	}
	return "fixed_string"
}


