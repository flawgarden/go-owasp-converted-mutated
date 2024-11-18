package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01658 struct{}

func (b *BenchmarkTest01658) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest01658="
	paramLoc := -1

	if queryString != "" {
		paramLoc = indexOf(queryString, paramVal)
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01658"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	w.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	return param
}

func indexOf(s, substr string) int {
	return strings.Index(s, substr)
}
