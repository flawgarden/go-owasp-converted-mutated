//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Snyk analysis results: [79]
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02578/BenchmarkTest02578.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/pointers.tmt with name pointer_init_positive 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02578 struct{}

func (b *BenchmarkTest02578) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery

var ptr123 *string = new(string)
*ptr123 = queryString
queryString = *ptr123

	paramval := "BenchmarkTest02578="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02578' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(param, "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}

func addSuffix(s *string, suf string) {
	*s = *s + suf
}

func addSuffixDoublePointer(s **string, suf *string) {
	**s = **s + *suf
}

func addSuffixDoublePointerBroken(s **string, suf *string) {
	*s = new(string)
	**s = **s + *suf
}

func swapStrings(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func removeSpaces(s *string) {
    *s = strings.ReplaceAll(*s, " ", "")
}


