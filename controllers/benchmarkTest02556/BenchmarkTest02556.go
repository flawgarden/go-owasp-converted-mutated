package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest02556 struct{}

func (b *BenchmarkTest02556) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02556="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if queryString[paramLoc:paramLoc+len(paramval)] != paramval {
			http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02556' in query string.", http.StatusBadRequest)
			return
		}
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	if idx := paramLoc + len(paramval); idx < len(queryString) {
		ampersandLoc = idx + len(paramval)
	}
	param = queryString[paramLoc+len(paramval) : ampersandLoc]

	decodedParam, _ := url.QueryUnescape(param)
	bar := doSomething(decodedParam)

	fileTarget := fmt.Sprintf("/path/to/file/%s", bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Access to file: '%s' created.", fileTarget)
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		fmt.Fprintln(w, " But file doesn't exist yet.")
	} else {
		fmt.Fprintln(w, " And file already exists.")
	}
}

func doSomething(param string) string {
	bar := param
	num := 196
	if (500/42)+num > 200 {
		return bar
	}
	return "This should never happen"
}
