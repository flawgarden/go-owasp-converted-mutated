package controllers

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BenchmarkTest02558 struct{}

func (b *BenchmarkTest02558) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest02558="
	paramLoc := strings.Index(queryString, paramVal)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02558' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)
	fileTarget := "/" + bar + "/Test.txt"

	w.Write([]byte("Access to file: '" + fileTarget + "' created.\n"))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists.\n"))
	} else {
		w.Write([]byte(" But file doesn't exist yet.\n"))
	}
}

func doSomething(param string) string {
	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
