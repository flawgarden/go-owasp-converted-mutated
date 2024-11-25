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
//Original file name: controllers/benchmarkTest02597/BenchmarkTest02597.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/defer.tmt with name simple_non_defer_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02597 struct{}

func (b *BenchmarkTest02597) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02597="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02597"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

func() {
    bar = "iFsgQ"
}()

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz19350 := strings.Builder{}
		sbxyz19350.WriteString(param)
		bar = sbxyz19350.String()[:len(param)-1] + "Z"
	}
	return bar
}

func main() {
	http.Handle("/xss-05/BenchmarkTest02597", &BenchmarkTest02597{})
	http.ListenAndServe(":8080", nil)
}

func foo(f string) (s string) {
	defer func() {
		s = "constant_string"
	}()
	s = f + " suffix"
	return s
}

func foo2(f string) (s string) {
	defer func() {
		s = s + f
	}()
	s = f + " suffix"
	return s
}

