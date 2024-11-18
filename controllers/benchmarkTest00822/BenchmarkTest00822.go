package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest00822 struct{}

func (b *BenchmarkTest00822) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00822="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00822' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	bar := "This should never happen"
	num := 196

	if (500/42)+num > 200 {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte("Parameter value: " + bar))
}
