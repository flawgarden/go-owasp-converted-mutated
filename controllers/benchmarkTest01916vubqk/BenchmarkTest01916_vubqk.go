//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: [116, 79]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest01916/BenchmarkTest01916.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/assertions.tmt with name type_assertion_in_switch_positive 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest01916 struct{}

func (b *BenchmarkTest01916) Get(w http.ResponseWriter, r *http.Request) {
	b.Post(w, r)
}

func (b *BenchmarkTest01916) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

var i interface{} = param
switch i.(type) {
case int:
    param = "RmFsr"
case string:
    param = param + "RAIjG"
default:
    param = "jNGYC"
}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func (b *BenchmarkTest01916) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}
