package controllers

import (
	"net/http"
	"net/url"
	"strings"
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

set787231 := make(map[string]struct{})
set787231["RAGQc"] = struct{}{}
set787231["qOivu"] = struct{}{}
var value7845 string = "hpWHO"
if _, exists := set787231[bar]; exists {
    value7845 = bar
}
bar = value7845

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
