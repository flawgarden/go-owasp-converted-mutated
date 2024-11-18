// FAIL
// Semgrep original results: [79]
// Gosec original results: []
// CodeQL original results: [79]
// Snyk original results: [79]
// -------------
// Semgrep analysis results: [79]
// CodeQL analysis results: []
// Snyk analysis results: [79]
// Gosec analysis results: []
// Original file name: controllers/benchmarkTest02600/BenchmarkTest02600.go
// Original file CWE's: [79]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/time.tmt with name simple_time_neutral
// Used extensions:
// Program:
package controllers

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BenchmarkTest02600 struct{}

func (b *BenchmarkTest02600) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02600="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02600' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	decodedParam, _ := url.QueryUnescape(param)

	bar := doSomething(decodedParam)

	time := time.Now()
	if time.Hour() > 12 {
		bar = "fixed_string"
	}

	w.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		w.Write([]byte(bar))
	}
}

func doSomething(param string) string {
	num := 106
	bar := ""

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	return bar
}
