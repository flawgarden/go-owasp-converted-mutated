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
//Original file name: controllers/benchmarkTest01173/BenchmarkTest01173.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/embedding.tmt with name embedding_simple_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01173 struct{}

func (bt *BenchmarkTest01173) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("X-XSS-Protection")
	param := r.Header.Get("Referer")

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")

d123 := Derived{
    Base: Base{Value: bar},
}
d123.Value = "const_string"
bar = d123.Value

	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the param value
	}
	return bar
}

func init() {
	http.Handle("/xss-02/BenchmarkTest01173", &BenchmarkTest01173{})
}
