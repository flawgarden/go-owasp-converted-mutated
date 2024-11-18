package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02585 struct{}

func (b *BenchmarkTest02585) Get(w http.ResponseWriter, r *http.Request) {
	b.Post(w, r)
}

func (b *BenchmarkTest02585) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02585="
	paramLoc := -1
	if queryString != "" {
		paramLoc = getIndex(queryString, paramval)
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02585"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := getIndex(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func getIndex(s, substr string) int {
	return -1 // Replace with actual implementation for finding index
}
