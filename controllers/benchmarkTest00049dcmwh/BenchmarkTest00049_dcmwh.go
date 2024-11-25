//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79, 319]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest00049/BenchmarkTest00049.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/dowhile.tmt with name do_while_operator_negative 
//Used extensions: EXPR_bool -> strings.EqualFold(~[EXPR_str]~, ~[EXPR_str]~)
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00049 struct{}

func (bt *BenchmarkTest00049) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00049="

for {
    queryString = ""
    if !strings.EqualFold(queryString, queryString) {
        break
    }
}

	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00049"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(param, "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte("Parameter value: " + param))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00049", &BenchmarkTest00049{})
	http.ListenAndServe(":8080", nil)
}
